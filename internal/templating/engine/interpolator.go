package engine

import (
	"bytes"
	"errors"
	"fmt"
	"maps"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:generate go tool enumer -type=interpolatorEntity -transform=snake -output=interpolator_entity_string.go -linecomment=true

// Interpolator is used for interpolating placeholder variables like
// ($ .project $) or ($ project $) with actual values
type Interpolator struct {
	Version Version

	Organization     string
	OrganizationName string

	ParentOrganization     string
	ParentOrganizationName string

	Project             string
	ProjectName         string
	ProjectOwnerCan     string
	ProjectOwnerName    string
	ProjectOwnerSurname string
	ProjectOwnerEmail   string

	Environment         string
	EnvironmentName     string
	EnvironmentType     string
	EnvironmentTypeName string

	SCSURL             string
	SCSBranch          string
	SCSCredentialType  string
	SCSCredentialPath  string
	SCS                string
	SCSName            string
	StackPath          string
	Stack              string
	StackName          string
	StackVersionType   string
	StackVersionName   string
	StackVersionCommit string

	CRURL            string
	CRBranch         string
	CRCredentialType string
	CRCredentialPath string
	CR               string
	CRName           string

	Component        string
	ComponentName    string
	ComponentUseCase string

	StackFormUpdatedByUser      string
	StackFormUpdatedByUserEmail string

	InventoryJWT string
	APIURL       string
	ConsoleURL   string

	ConfigRoot string

	CurrentUserUsername string

	// EnvVars holds the per-environment variables exposed to templates as
	// `.env_vars.<key>` (aliased `.environment_vars.<key>`). Keys map to
	// user-defined variable names; values are the decoded Go values (string,
	// bool, int64, list, map).
	EnvVars map[string]any

	// EnvProvider holds the cloud accounts attached to the environment,
	// exposed to templates under the `.env_providers` top-level key (aliased
	// `.environment_providers`) and keyed by the cloud provider canonical (aws,
	// google, azurerm) or, for custom providers, the cloud account canonical.
	// Each entry flattens the account name/canonical with the credential
	// metadata and decoded values, e.g. `($ .env_providers.aws.access_key $)`.
	EnvProvider map[string]any
}

const (
	InterpolatorDelimLeft  = "($"
	InterpolatorDelimRight = "$)"
	defaultTemplateName    = "cycloid_interpolator"
)

var reCyInterpolation = regexp.MustCompile(`\(\$([^)]+)\$\)`)

// HasInterpolation checks if 's' has ($ $)
func HasInterpolation(s string) bool {
	if res := reCyInterpolation.FindAllStringSubmatch(s, -1); len(res) != 0 {
		return true
	}
	return false
}

// Interpolate will interpolate the given string with the actual values
// defined by the interpolator. Parameter templateName is used to build a meaningful error message
// if the interpolation fails.
func (i Interpolator) Interpolate(s, templateName string) (string, error) {
	result, err := i.InterpolateWithExtraData(s, templateName, nil)
	if err != nil {
		return "", fmt.Errorf("failed to interpolate: %w", err)
	}

	return result, nil
}

// InterpolateWithExtraData will interpolate the given string with the actual values
// defined by the interpolator and additional, passed by extraData parameter.
// If extraData defines fields already defined by the interpolator, they will be overridden.
// Extra data is omitted if using a deprecated version.
// Parameter templateName is optional.
func (i Interpolator) InterpolateWithExtraData(s, templateName string, extraData map[string]interface{}) (string, error) {
	if !i.Version.IsAVersion() {
		return "", fmt.Errorf("invalid interpolator version: %s", i.Version)
	}

	s = escapeLinesWithIncompleteDelimiter(s)

	// inventory_jwt must never be substituted with an empty value: callers
	// persist the rendered string and a silently empty JWT defeats inventory
	// authorization at the consumer. extraData can override interpolator
	// fields on the new-style path (the deprecated string-replace ignores
	// extraData), so a caller-supplied JWT counts as available there.
	effectiveJWT := i.InventoryJWT
	if effectiveJWT == "" && i.Version.IsNewInterpolation() {
		if v, ok := extraData[inventoryJWT.String()].(string); ok {
			effectiveJWT = v
		}
	}
	if effectiveJWT == "" && reInventoryJWTRef.MatchString(s) {
		msg := `inventory_jwt is referenced by the template but is not available — ensure an ExternalBackend providing a JWT is configured at the org, project, environment, or component scope`
		if templateName != "" {
			msg = fmt.Sprintf("template %q: %s", templateName, msg)
		}
		return "", errors.New(msg)
	}

	if !i.Version.IsNewInterpolation() {
		result, err := i.interpolateDeprecatedVersion(s)
		if err != nil {
			return "", fmt.Errorf("error interpolating: %w", err)
		}

		return unescapeIncompleteDelimiters(result), nil
	}

	result, err := i.interpolate(s, templateName, extraData)
	if err != nil {
		return "", fmt.Errorf("error interpolating: %w", err)
	}

	return unescapeIncompleteDelimiters(result), nil
}

// escapeLinesWithIncompleteDelimiter will escape the lines that have the left delimiter
// but not the right one
// It works for multiline incomplete delimiters as well
func escapeLinesWithIncompleteDelimiter(s string) string {
	lines := strings.Split(s, "\n")
	var result []string
	var buffer string
	inIncompleteDelimiter := false

	for _, line := range lines {
		if strings.Contains(line, InterpolatorDelimLeft) && !inIncompleteDelimiter {
			if strings.Contains(line, InterpolatorDelimRight) {
				result = append(result, line)
			} else {
				buffer = line
				inIncompleteDelimiter = true
			}
		} else if inIncompleteDelimiter {
			buffer += "\n" + line
			if strings.Contains(line, InterpolatorDelimRight) {
				result = append(result, buffer)
				buffer = ""
				inIncompleteDelimiter = false
			}
		} else {
			result = append(result, line)
		}
	}

	if inIncompleteDelimiter {
		result = append(result, escapeDelimiter(buffer))
	}

	return strings.Join(result, "\n")
}

// escapeDelimiter will escape the given delimiter in the string
func escapeDelimiter(s string) string {
	s = strings.ReplaceAll(s, InterpolatorDelimLeft, `\(\$`)
	return s
}

// unescapeIncompleteDelimiters will unescape the delimiters in the given string
func unescapeIncompleteDelimiters(s string) string {
	s = strings.ReplaceAll(s, `\(\$`, InterpolatorDelimLeft)
	s = strings.ReplaceAll(s, `\$\)`, InterpolatorDelimRight)
	return s
}

// interpolate will replace all placeholders with actual values from
// Interpolator. If value is not present, the placeholder will not be
// interpolated. This version of the interpolator (> V2) uses the Go template
// for rendering. Parameter templateName is optional and set to default if not given.
// Parameter extraData defines additional variables used to execute a template.
// If extraData defines fields already defined by the interpolator, they will be overridden.
func (i Interpolator) interpolate(s, templateName string, extraData map[string]interface{}) (string, error) {
	data := i.dataMap()
	// Add extra data if any
	maps.Copy(data, extraData)

	if templateName == "" {
		templateName = defaultTemplateName
	}

	getOption := func(strict bool) string {
		if strict {
			return "missingkey=error"
		}
		return "missingkey=zero"
	}

	strict := false
	tmpl := template.New(templateName).
		Delims(InterpolatorDelimLeft, InterpolatorDelimRight).
		Option(getOption(strict))

	// Init function map with extra functions
	funcMap := sprig.FuncMap()
	// Security reasons
	delete(funcMap, "env")
	delete(funcMap, "expandenv")
	// Add helm functions
	helmFuncs := FuncMap(tmpl, strict)
	maps.Copy(funcMap, helmFuncs)
	// Assign to the template
	tmpl.Funcs(funcMap)

	tmpl, err := tmpl.Parse(s)
	if err != nil {
		return "", wrapInterpolatorErr(templateName, err)
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, data)
	if err != nil {
		return "", fmt.Errorf("error rendering template: %w", err)
	}

	return buff.String(), nil
}

type interpolatorEntity int

// These are the key for the rendering data
const (
	// Organization related interpolation keys
	// org == organization == orgCanonical == organizationCanonical
	// orgName == organizationName
	org                   interpolatorEntity = iota // org
	organization                                    // organization
	orgCanonical                                    // org_canonical
	organizationCanonical                           // organization_canonical
	orgName                                         // org_name
	organizationName                                // organization_name

	parentOrg                   // parent_org
	parentOrganization          // parent_organization
	parentOrgCanonical          // parent_org_canonical
	parentOrganizationCanonical // parent_organization_canonical
	parentOrgName               // parent_org_name
	parentOrganizationName      // parent_organization_name

	// Project related interpolation keys
	// project == projectCanonical
	// projectOwner == projectOwnerCanonical
	project               // project
	projectCanonical      // project_canonical
	projectName           // project_name
	projectOwnerCanonical // project_owner_canonical
	projectOwner          // project_owner
	projectOwnerName      // project_owner_name
	projectOwnerSurname   // project_owner_surname
	projectOwnerEmail     // project_owner_email

	// Environment related interpolation keys
	// env == environment == envCanonical == environmentCanonical
	// envName == environmentName
	env                      // env
	environment              // environment
	envCanonical             // env_canonical
	environmentCanonical     // environment_canonical
	envName                  // env_name
	environmentName          // environment_name
	envType                  // env_type
	environmentType          // environment_type
	envTypeCanonical         // env_type_canonical
	environmentTypeCanonical // environment_type_canonical
	envTypeName              // env_type_name
	environmentTypeName      // environment_type_name

	// Component related interpolation keys
	// component == componentCanonical
	component                      // component
	componentCanonical             // component_canonical
	componentName                  // component_name
	stackFormUpdatedByUser         // stackform_updated_by_user
	stackFormUpdatedByUserUsername // stackform_updated_by_user_username
	stackFormUpdatedByUserEmail    // stackform_updated_by_user_email

	// Stack related interpolation keys
	// stack == stackCanonical
	// catalogRepository == catalogRepositoryCanonical
	// Deprecated, use catalogRepositoryURL
	scsURL // scs_url
	// Deprecated, use catalogRepositoryBranch
	scsBranch // scs_branch
	// Deprecated, use catalogRepositoryCredentialType
	scsCredType // scs_cred_type
	// Deprecated, use catalogRepositoryCredentialPath
	scsCredPath // scs_cred_path
	// Deprecated, use catalogRepositoryCanonical
	scsCanonical // scs_canonical
	// Deprecated, use catalogRepositoryName
	scsName                         // scs_name
	catalogRepository               // catalog_repository
	catalogRepositoryCanonical      // catalog_repository_canonical
	catalogRepositoryName           // catalog_repository_name
	catalogRepositoryURL            // catalog_repository_url
	catalogRepositoryBranch         // catalog_repository_branch
	catalogRepositoryCredentialType // catalog_repository_credential_type
	catalogRepositoryCredentialPath // catalog_repository_credential_path
	stack                           // stack
	stackCanonical                  // stack_canonical
	stackName                       // stack_name
	stackPath                       // stack_path
	stackVersionName                // stack_version_name
	stackVersionRef                 // stack_version_ref
	stackVersionCommit              // stack_version_commit
	stackVersionType                // stack_version_type

	// CR related interpolation keys
	// configRepository == configRepositoryCanonical
	// Deprecated, use configRepositoryURL
	crURL // cr_url
	// Deprecated, use configRepositoryBranch
	crBranch // cr_branch
	// Deprecated, use configRepositoryCredentialType
	crCredType // cr_cred_type
	// Deprecated, use configRepositoryCredentialPath
	crCredPath                     // cr_cred_path
	configRepository               // config_repository
	configRepositoryCanonical      // config_repository_canonical
	configRepositoryName           // config_repository_name
	configRepositoryURL            // config_repository_url
	configRepositoryBranch         // config_repository_branch
	configRepositoryCredentialType // config_repository_credential_type
	configRepositoryCredentialPath // config_repository_credential_path

	// Miscellaneous interpolation keys
	inventoryJWT // inventory_jwt
	apiURL       // api_url
	consoleURL   // console_url
	useCase      // use_case
	configRoot   // config_root

	currentUserUsername // current_user_username
)

func (i Interpolator) dataMap() map[string]any {
	// These are the mappings used to render the template
	mappings := map[interpolatorEntity]any{}
	if i.Organization != "" {
		mappings[organizationCanonical] = i.Organization
		mappings[organization] = i.Organization
		mappings[orgCanonical] = i.Organization
		mappings[org] = i.Organization
	}
	if i.OrganizationName != "" {
		mappings[organizationName] = i.OrganizationName
		mappings[orgName] = i.OrganizationName
	}
	if i.ParentOrganization != "" {
		mappings[parentOrg] = i.ParentOrganization
		mappings[parentOrganization] = i.ParentOrganization
		mappings[parentOrgCanonical] = i.ParentOrganization
		mappings[parentOrganizationCanonical] = i.ParentOrganization
	}
	if i.ParentOrganizationName != "" {
		mappings[parentOrgName] = i.ParentOrganizationName
		mappings[parentOrganizationName] = i.ParentOrganizationName
	}
	if i.Project != "" {
		mappings[project] = i.Project
		mappings[projectCanonical] = i.Project
	}
	if i.ProjectName != "" {
		mappings[projectName] = i.ProjectName
	}
	if i.ProjectOwnerCan != "" {
		mappings[projectOwnerCanonical] = i.ProjectOwnerCan
		mappings[projectOwner] = i.ProjectOwnerCan
	}
	if i.ProjectOwnerName != "" {
		mappings[projectOwnerName] = i.ProjectOwnerName
	}
	if i.ProjectOwnerSurname != "" {
		mappings[projectOwnerSurname] = i.ProjectOwnerSurname
	}
	if i.ProjectOwnerEmail != "" {
		mappings[projectOwnerEmail] = i.ProjectOwnerEmail
	}
	if i.Environment != "" {
		mappings[environment] = i.Environment
		mappings[environmentCanonical] = i.Environment
		mappings[env] = i.Environment
		mappings[envCanonical] = i.Environment
	}
	if i.EnvironmentName != "" {
		mappings[environmentName] = i.EnvironmentName
		mappings[envName] = i.EnvironmentName
	}
	if i.EnvironmentType != "" {
		mappings[envType] = i.EnvironmentType
		mappings[environmentType] = i.EnvironmentType
		mappings[envTypeCanonical] = i.EnvironmentType
		mappings[environmentTypeCanonical] = i.EnvironmentType
	}
	if i.EnvironmentTypeName != "" {
		mappings[envTypeName] = i.EnvironmentTypeName
		mappings[environmentTypeName] = i.EnvironmentTypeName
	}
	if i.SCS != "" {
		mappings[scsCanonical] = i.SCS
		mappings[catalogRepositoryCanonical] = i.SCS
		mappings[catalogRepository] = i.SCS
	}
	if i.SCSName != "" {
		mappings[scsName] = i.SCSName
		mappings[catalogRepositoryName] = i.SCSName
	}
	if i.SCSURL != "" {
		mappings[scsURL] = i.SCSURL
		mappings[catalogRepositoryURL] = i.SCSURL
	}
	if i.SCSBranch != "" {
		mappings[scsBranch] = i.SCSBranch
		mappings[catalogRepositoryBranch] = i.SCSBranch
	}
	if i.SCSCredentialType != "" {
		mappings[scsCredType] = i.SCSCredentialType
		mappings[catalogRepositoryCredentialType] = i.SCSCredentialType
	}
	if i.SCSCredentialPath != "" {
		mappings[scsCredPath] = i.SCSCredentialPath
		mappings[catalogRepositoryCredentialPath] = i.SCSCredentialPath
	}
	if i.Stack != "" {
		mappings[stack] = i.Stack
		mappings[stackCanonical] = i.Stack
	}
	if i.StackName != "" {
		mappings[stackName] = i.StackName
	}
	if i.StackVersionName != "" {
		mappings[stackVersionName] = i.StackVersionName
		mappings[stackVersionRef] = i.StackVersionName
	}
	if i.StackVersionCommit != "" {
		mappings[stackVersionCommit] = i.StackVersionCommit
	}
	if i.StackVersionType != "" {
		mappings[stackVersionType] = i.StackVersionType
	}
	if i.StackPath != "" {
		mappings[stackPath] = i.StackPath
	}
	if i.CR != "" {
		mappings[configRepository] = i.CR
		mappings[configRepositoryCanonical] = i.CR
	}
	if i.CRName != "" {
		mappings[configRepositoryName] = i.CRName
	}
	if i.CRURL != "" {
		mappings[crURL] = i.CRURL
		mappings[configRepositoryURL] = i.CRURL
	}
	if i.CRBranch != "" {
		mappings[crBranch] = i.CRBranch
		mappings[configRepositoryBranch] = i.CRBranch
	}
	if i.CRCredentialType != "" {
		mappings[crCredType] = i.CRCredentialType
		mappings[configRepositoryCredentialType] = i.CRCredentialType
	}
	if i.CRCredentialPath != "" {
		mappings[crCredPath] = i.CRCredentialPath
		mappings[configRepositoryCredentialPath] = i.CRCredentialPath
	}
	if i.InventoryJWT != "" {
		mappings[inventoryJWT] = i.InventoryJWT
	}
	if i.APIURL != "" {
		mappings[apiURL] = i.APIURL
	}
	if i.ConsoleURL != "" {
		mappings[consoleURL] = i.ConsoleURL
	}
	if i.ComponentUseCase != "" {
		mappings[useCase] = i.ComponentUseCase
	}
	if i.Component != "" {
		mappings[component] = i.Component
		mappings[componentCanonical] = i.Component
	}
	if i.ComponentName != "" {
		mappings[componentName] = i.ComponentName
	}
	if i.StackFormUpdatedByUser != "" {
		mappings[stackFormUpdatedByUser] = i.StackFormUpdatedByUser
		mappings[stackFormUpdatedByUserUsername] = i.StackFormUpdatedByUser
	}
	if i.StackFormUpdatedByUserEmail != "" {
		mappings[stackFormUpdatedByUserEmail] = i.StackFormUpdatedByUserEmail
	}
	if i.ConfigRoot != "" {
		mappings[configRoot] = i.ConfigRoot
	}
	if i.CurrentUserUsername != "" {
		mappings[currentUserUsername] = i.CurrentUserUsername
	}

	// Convert the map to a map[string]interface{}
	// so it can be used in the template

	data := make(map[string]any)
	for k, v := range mappings {
		data[k.String()] = v
	}

	// environment stays a flat string (set above via mappings) so it keeps
	// working in comparisons and pipes, e.g. ($ if eq .environment "prod" $)
	// or ($ .environment | upper $), exactly as before environment governance.
	// Per-environment variables and cloud accounts are dynamic collections, so
	// they are exposed as their own top-level map keys (indexable by user or
	// cloud data), each under both the env_ and environment_ prefix to alias
	// like the scalar env_*/environment_* vars above.
	if len(i.EnvVars) > 0 {
		data["env_vars"] = i.EnvVars
		data["environment_vars"] = i.EnvVars
	}
	if len(i.EnvProvider) > 0 {
		data["env_providers"] = i.EnvProvider
		data["environment_providers"] = i.EnvProvider
	}

	return data
}

// Values below represent known placeholders
// Deprecated, they're being used in the outdated version interpolation
var (
	// rePrj is the regexp to replace '($ project $)'
	rePrj = regexp.MustCompile(`\(\$(?:\s+)?project(?:\s+)?\$\)`)

	// rePrjOwnerCan is the regexp to replace '($ project_owner_canonical $)'
	rePrjOwnerCan = regexp.MustCompile(`\(\$(?:\s+)?project_owner_canonical(?:\s+)?\$\)`)

	// rePrjOwnerName is the regexp to replace '($ project_owner_name $)'
	rePrjOwnerName = regexp.MustCompile(`\(\$(?:\s+)?project_owner_name(?:\s+)?\$\)`)

	// rePrjOwnerSurname is the regexp to replace '($ project_owner_surname $)'
	rePrjOwnerSurname = regexp.MustCompile(`\(\$(?:\s+)?project_owner_surname(?:\s+)?\$\)`)

	// rePrjOwnerEmail is the regexp to replace '($ project_owner_email $)'
	rePrjOwnerEmail = regexp.MustCompile(`\(\$(?:\s+)?project_owner_email(?:\s+)?\$\)`)

	// reEnv is the regexp to replace '($ environment $)'
	reEnv = regexp.MustCompile(`\(\$(?:\s+)?environment(?:\s+)?\$\)`)

	// reOrg is the regexp to replace '($ organization_canonical $)'
	reOrg = regexp.MustCompile(`\(\$(?:\s+)?organization_canonical(?:\s+)?\$\)`)

	// reSCSURL is the regexp to replace '($ scs_url $)'
	reSCSURL = regexp.MustCompile(`\(\$(?:\s+)?scs_url(?:\s+)?\$\)`)

	// reSCSBranch is the regexp to replace '($ scs_branch $)'
	reSCSBranch = regexp.MustCompile(`\(\$(?:\s+)?scs_branch(?:\s+)?\$\)`)

	// reSCSrecentialType is the regexp to replace '($ scs_cred_type $)'
	reSCSrecentialType = regexp.MustCompile(`\(\$(?:\s+)?scs_cred_type(?:\s+)?\$\)`)

	// reSCSrecentialPath is the regexp to replace '($ scs_cred_path $)'
	reSCSrecentialPath = regexp.MustCompile(`\(\$(?:\s+)?scs_cred_path(?:\s+)?\$\)`)

	// reStackPath is the regexp to replace '($ stack_path $)'
	reStackPath = regexp.MustCompile(`\(\$(?:\s+)?stack_path(?:\s+)?\$\)`)

	// reCRURL is the regexp to replace '($ cr_url $)'
	reCRURL = regexp.MustCompile(`\(\$(?:\s+)?cr_url(?:\s+)?\$\)`)

	// reCRBranch is the regexp to replace '($ cr_branch $)'
	reCRBranch = regexp.MustCompile(`\(\$(?:\s+)?cr_branch(?:\s+)?\$\)`)

	// reCRCrecentialType is the regexp to replace '($ cr_cred_type $)'
	reCRCrecentialType = regexp.MustCompile(`\(\$(?:\s+)?cr_cred_type(?:\s+)?\$\)`)

	// reCRCrecentialPath is the regexp to replace '($ cr_cred_path $)'
	reCRCrecentialPath = regexp.MustCompile(`\(\$(?:\s+)?cr_cred_path(?:\s+)?\$\)`)

	// reInventoryJWT is the regexp to replace '($ inventory_jwt $)'
	reInventoryJWT = regexp.MustCompile(`\(\$(?:\s+)?inventory_jwt(?:\s+)?\$\)`)

	// reInventoryJWTRef matches a bare `($ inventory_jwt $)` or
	// `($ .inventory_jwt $)`. Wrapper expressions
	// (`($ if .inventory_jwt $)`, `($ default "" .inventory_jwt $)`,
	// `($ .inventory_jwt | upper $)`) don't match, but a bare reference
	// inside a guarded body does — and that's intentional: those guards
	// existed only to work around the EB-creation-ordering bug that this
	// change fixes at the source, so callers can drop them.
	reInventoryJWTRef = regexp.MustCompile(`\(\$\s*\.?\s*inventory_jwt\s*\$\)`)

	// reAPIURL is the regexp to replace '($ api_url $)'
	reAPIURL = regexp.MustCompile(`\(\$(?:\s+)?api_url(?:\s+)?\$\)`)

	// reConsoleURL is the regexp to replace '($ console_url $)'
	reConsoleURL = regexp.MustCompile(`\(\$(?:\s+)?console_url(?:\s+)?\$\)`)

	// reUseCase is the regexp to replace '($ use_case $)'
	reUseCase = regexp.MustCompile(`\(\$(?:\s+)?use_case(?:\s+)?\$\)`)

	// reComponent is the regexp to replace '($ component $)'
	reComponent = regexp.MustCompile(`\(\$(?:\s+)?component(?:\s+)?\$\)`)
)

func (i Interpolator) interpolateDeprecatedVersion(s string) (string, error) {
	if i.Organization != "" {
		s = reOrg.ReplaceAllString(s, i.Organization)
	}
	if i.Project != "" {
		s = rePrj.ReplaceAllString(s, i.Project)
	}
	if i.ProjectOwnerCan != "" {
		s = rePrjOwnerCan.ReplaceAllString(s, i.ProjectOwnerCan)
	}
	if i.ProjectOwnerName != "" {
		s = rePrjOwnerName.ReplaceAllString(s, i.ProjectOwnerName)
	}
	if i.ProjectOwnerSurname != "" {
		s = rePrjOwnerSurname.ReplaceAllString(s, i.ProjectOwnerSurname)
	}
	if i.ProjectOwnerEmail != "" {
		s = rePrjOwnerEmail.ReplaceAllString(s, i.ProjectOwnerEmail)
	}
	if i.Environment != "" {
		s = reEnv.ReplaceAllString(s, i.Environment)
	}
	if i.SCSURL != "" {
		s = reSCSURL.ReplaceAllString(s, i.SCSURL)
	}
	if i.SCSBranch != "" {
		s = reSCSBranch.ReplaceAllString(s, i.SCSBranch)
	}
	if i.SCSCredentialType != "" {
		s = reSCSrecentialType.ReplaceAllString(s, i.SCSCredentialType)
	}
	if i.SCSCredentialPath != "" {
		s = reSCSrecentialPath.ReplaceAllString(s, i.SCSCredentialPath)
	}
	if i.StackPath != "" {
		s = reStackPath.ReplaceAllString(s, i.StackPath)
	}
	if i.CRURL != "" {
		s = reCRURL.ReplaceAllString(s, i.CRURL)
	}
	if i.CRBranch != "" {
		s = reCRBranch.ReplaceAllString(s, i.CRBranch)
	}
	if i.CRCredentialType != "" {
		s = reCRCrecentialType.ReplaceAllString(s, i.CRCredentialType)
	}
	if i.CRCredentialPath != "" {
		s = reCRCrecentialPath.ReplaceAllString(s, i.CRCredentialPath)
	}
	if i.InventoryJWT != "" {
		s = reInventoryJWT.ReplaceAllString(s, i.InventoryJWT)
	}
	if i.APIURL != "" {
		s = reAPIURL.ReplaceAllString(s, i.APIURL)
	}
	if i.ConsoleURL != "" {
		s = reConsoleURL.ReplaceAllString(s, i.ConsoleURL)
	}
	if i.ComponentUseCase != "" {
		s = reUseCase.ReplaceAllString(s, i.ComponentUseCase)
	}
	if i.Component != "" {
		s = reComponent.ReplaceAllString(s, i.Component)
	}
	return s, nil
}

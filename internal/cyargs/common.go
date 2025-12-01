package cyargs

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

var (
	v            = viper.GetViper()
	DefaultIcon  = ""
	DefaultColor = ""
	ValidIcons   = []string{
		"folder_open",
		"mdi-cube-outline",
		"public",
		"extension",
		"science",
		"bug_report",
		"bolt",
		"call_merge",
		"commit",
		"mdi-source-branch",
		"traffic",
		"mdi-clipboard-check-outline",
		"mdi-progress-clock",
		"visibility",
		"vpn_key",
		"lightbulb",
		"favorite",
		"star",
		"auto_awesome",
		"mdi-controller-classic",
		"precision_manufacturing",
		"tour",
		"podcasts",
		"inventory",
		"save",
		"security",
		"mdi-lifebuoy",
		"mdi-ab-testing",
		"mdi-api",
		"mdi-console",
		"mdi-database",
		"mdi-vpn",
		"mdi-server",
		"mdi-server-security",
		"mdi-network-outline",
		"mdi-lan",
		"mdi-nas",
		"mdi-ansible",
		"mdi-aws",
		"mdi-microsoft-azure",
		"mdi-google-cloud",
		"mdi-kubernetes",
		"mdi-terraform",
	}
	ValidColors = []string{
		"dev",
		"prod",
		"demo",
		"success",
		"default",
		"error",
		"staging",
		"preprod",
	}
)

func AddCyContext(cmd *cobra.Command) {
	AddProjectFlag(cmd)
	AddEnvFlag(cmd)
	AddComponentFlag(cmd)
}

func GetCyContext(cmd *cobra.Command) (org, project, env, component string, err error) {
	org, err = GetOrg(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	project, err = GetProject(cmd)
	if err != nil {
		return org, "", "", "", err
	}

	env, err = GetEnv(cmd)
	if err != nil {
		return org, project, "", "", err
	}

	component, err = GetComponent(cmd)
	if err != nil {
		return org, project, env, "", err
	}

	return org, project, env, component, nil
}

func AddNameFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("name", "n", "", "set a human friendly name.")
}

func GetName(cmd *cobra.Command) (string, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return "", err
	}

	return name, nil
}

func AddProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("project", "p", "", "the project canonical, can also be set with the CY_PROJECT env var")
	v.BindPFlag("project", cmd.Flags().Lookup("project"))
	cmd.RegisterFlagCompletionFunc("project", func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
				cobra.ShellCompDirectiveNoFileComp
		}

		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		projects, err := m.ListProjects(org)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to fetch project list for completion in org '"+org+"': "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var canonicals = make([]cobra.Completion, len(projects))
		for index, project := range projects {
			if project.Canonical != nil && strings.HasPrefix(*project.Canonical, toComplete) {
				canonicals[index] = cobra.CompletionWithDesc(*project.Canonical, *project.Name)
			}
		}

		return canonicals, cobra.ShellCompDirectiveNoFileComp
	})
}

func GetProject(cmd *cobra.Command) (string, error) {
	project, _ := cmd.Flags().GetString("project")
	if project != "" {
		return project, nil
	}

	project = v.GetString("project")
	if project == "" {
		return "", errors.New("project is not set, use --project flag or CY_PROJECT env var")
	}

	return project, nil
}

func AddEnvFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("env", "e", "", "the env canonical, can also be set with the CY_ENV env var")
	v.BindPFlag("env", cmd.Flags().Lookup("env"))

	cmd.RegisterFlagCompletionFunc("env", func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		project, err := GetProject(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to fetch project list for completion in org '"+org+"': "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		envs, err := m.ListProjectsEnv(org, project)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to list env from org '"+org+"' in project '"+project+"' for completion: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var canonicals = make([]cobra.Completion, len(envs))
		for index, env := range envs {
			if env.Canonical != nil && strings.HasPrefix(*env.Canonical, toComplete) {
				canonicals[index] = cobra.CompletionWithDesc(*env.Canonical, env.Name)
			}
		}

		return canonicals, cobra.ShellCompDirectiveNoFileComp
	})
}

func GetEnv(cmd *cobra.Command) (string, error) {
	env, _ := cmd.Flags().GetString("env")
	if env != "" {
		return env, nil
	}

	env = v.GetString("env")
	if env == "" {
		return "", errors.New("env is not set, use --env flag or CY_ENV env var")
	}

	return env, nil
}

func AddComponentFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "the component canonical, can also be set with the CY_COMPONENT env var")
	v.BindPFlag("component", cmd.Flags().Lookup("component"))
	cmd.RegisterFlagCompletionFunc("component", func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		project, err := GetProject(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		env, err := GetEnv(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		components, err := m.ListComponents(org, project, env)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to list components for completion in org '"+org+"' with project '"+project+"' and env '"+env+"': "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var canonicals = make([]string, len(components))
		for index, component := range components {
			if component.Canonical != nil && strings.HasPrefix(*component.Canonical, toComplete) {
				canonicals[index] = cobra.CompletionWithDesc(*component.Canonical, *component.Name)
			}
		}

		return canonicals, cobra.ShellCompDirectiveNoFileComp
	})
}

func GetComponent(cmd *cobra.Command) (string, error) {
	component, _ := cmd.Flags().GetString("component")
	if component != "" {
		return component, nil
	}

	component = v.GetString("component")
	if component == "" {
		return "", errors.New("component is not set, use --component flag or CY_COMPONENT component var")
	}

	return component, nil
}

func AddColorFlag(cmd *cobra.Command) {
	cmd.Flags().String("color", DefaultColor, "set the color.")
	cmd.RegisterFlagCompletionFunc("color", func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		var completions []cobra.Completion
		for _, color := range ValidColors {
			if strings.HasPrefix(color, toComplete) {
				completions = append(completions, color)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	})
}

func GetColor(cmd *cobra.Command) (string, error) {
	color, err := cmd.Flags().GetString("color")
	if err != nil {
		return "", nil
	}

	return color, nil
}

// PickRandomColor will select a random color, if you
// fill env != nil it will try to infer the color based
// on the env name with our convention.
// This is define in the FE here:
// https://github.com/cycloidio/youdeploy-frontend-web/blob/develop/src/utils/config/icons.js
func PickRandomColor(env *string) string {
	if env != nil {
		switch *env {
		case "prod", "prd":
			return "prod"
		case "staging", "stg":
			return "staging"
		case "preprod", "pp", "pre-prd":
			return "preprod"
		case "test", "testing", "dev", "developpement":
			return "dev"
		case "demo", "integration", "intg":
			return "demo"
		default:
			return "success"
		}
	}

	randomIndex := rand.IntN(len(ValidColors))
	return ValidColors[randomIndex]
}

func AddIconFlag(cmd *cobra.Command) {
	cmd.Flags().String("icon", DefaultIcon, "set the icon.")
	err := cmd.RegisterFlagCompletionFunc("icon", func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		var completions []cobra.Completion
		for _, icon := range ValidIcons {
			if strings.HasPrefix(icon, toComplete) {
				completions = append(completions, icon)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		panic(err)
	}
}

func GetIcon(cmd *cobra.Command) (string, error) {
	icon, err := cmd.Flags().GetString("icon")
	if err != nil {
		return "", nil
	}

	return icon, nil
}

// PickRandomIcon will select a random icon, if you
// fill env != nil it will try to infer the icon based
// on the env name with our convention.
// This is define in the FE here:
// https://github.com/cycloidio/youdeploy-frontend-web/blob/develop/src/utils/config/icons.js
func PickRandomIcon(env *string) string {
	if env != nil {
		switch *env {
		case "prod", "prd":
			return "public"
		case "staging", "stg":
			return "science"
		case "preprod", "pp", "pre-prd":
			return "traffic"
		case "test", "testing", "dev", "developpement":
			return "commit"
		case "demo", "integration", "intg":
			return "auto_awesome"
		default:
			return "extension"
		}
	}

	randomIndex := rand.IntN(len(ValidIcons))
	return ValidIcons[randomIndex]
}

func AddOwnerFlag(cmd *cobra.Command) string {
	flagName := "owner"
	cmd.Flags().String(flagName, "", "canonical of a user to set as owner, will be the user attached to the current api key if empty.")
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		members, err := m.ListMembers(org)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "cannot list org members for owner completion: "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var owners = make([]cobra.Completion, len(members))
		for index, member := range members {
			if strings.HasPrefix(member.Username, toComplete) {
				owners[index] = cobra.CompletionWithDesc(
					member.Username,
					fmt.Sprintf("%s (%s)", member.FullName, member.Email.String()),
				)
			}
		}

		return owners, cobra.ShellCompDirectiveNoFileComp
	})

	return flagName
}

func GetOwner(cmd *cobra.Command) (string, error) {
	owner, err := cmd.Flags().GetString("owner")
	if err != nil {
		return "", nil
	}

	return owner, nil
}

func AddDescriptionFlag(cmd *cobra.Command) {
	cmd.Flags().String("description", "", "a human friendly description for your peers")
}

func GetDescription(cmd *cobra.Command) (string, error) {
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return "", nil
	}

	return description, nil
}

func GetOrg(cmd *cobra.Command) (string, error) {
	org := v.GetString("org")
	if org == "" {
		return "", fmt.Errorf("org is not set, use --org flag or CY_ORG env var, current value: %q", org)
	}

	return org, nil
}

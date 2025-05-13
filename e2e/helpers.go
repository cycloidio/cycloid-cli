package e2e

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	rootCmd "github.com/cycloidio/cycloid-cli/cmd"
)

var (
	now          = time.Now()
	NowTimestamp = now.UnixNano()

	TestGitSshKey = []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY
8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A
AAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V
25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==
-----END OPENSSH PRIVATE KEY-----`)

	TestPipelineSample = []byte(`jobs:
- name: job-hello-world
  build_logs_to_retain: 3
  plan:
  - task: hello-world
    config:
        platform: linux
        image_resource:
            type: docker-image
            source: {repository: busybox}
        run:
          path: /bin/sh
          args:
          - -ec
          - |
            echo ${MESSAGE}
        params:
          MESSAGE: ((message))`)

	TestPipelineVariables = []byte(`message: "hello world and especially to ($ organization_canonical $)"`)

	TestInfraPolicySample = []byte(`
	package test
	import input.tfplan as tfplan
	resource_types = { "aws_instance" }
	
	tags_ok(index) {
	  tfplan.resource_changes[index].change.after.tags["env"] == "test"
	}
	
	deny[reason] {
		resources_not_ok := [resource_not_ok | resource_types[i] ==  tfplan.resource_changes[j].type
											   not tags_ok(j); 
											   resource_not_ok := resource_types[i]]
		reason = sprintf("tag not in env: %s %s", [(resources_not_ok), test])
	}
`)
)

func LoginToRootOrg() {
	buf := new(bytes.Buffer)
	cmd := rootCmd.NewRootCommand()
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{
		"login",
		"--org", TestRootOrg,
		"--api-key", TestAPIKey,
	})

	err := cmd.Execute()
	if err != nil {
		panic(fmt.Sprintf("Test setup LoginToRootOrg, unable to login: %s", err.Error()))
	}
}

// executeCommandStdin will execute the command with args + stdin input
// and return stdin, stderr and err
func executeCommandStdin(stdin string, args []string) (string, string, error) {
	cmd := rootCmd.NewRootCommand()

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	cmd.SetOut(stdoutBuf)
	cmd.SetErr(stderrBuf)

	cmd.SetIn(strings.NewReader(stdin))
	cmd.SetArgs(args)
	cmdErr := cmd.Execute()
	stdout, err := io.ReadAll(stdoutBuf)
	if err != nil {
		return "", "", errors.Join(cmdErr, fmt.Errorf("failed to read stdout buffer from cli: %s", err))
	}

	stderr, err := io.ReadAll(stderrBuf)
	if err != nil {
		return string(stdout), "", errors.Join(cmdErr, fmt.Errorf("failed to read stderr buffer from cli: %s", err))
	}

	return string(stdout), string(stderr), cmdErr
}

func executeCommand(args []string) (string, error) {
	cmd := rootCmd.NewRootCommand()

	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(errBuf)

	cmd.SetArgs(args)
	cmdErr := cmd.Execute()
	cmdOut, err := io.ReadAll(buf)
	if err != nil {
		panic(fmt.Sprintf("Unable to read command output buffer"))
	}
	return string(cmdOut), cmdErr
}

// AddNowTimestamp add a timestamp suffix to a string
func AddNowTimestamp(txt string) string {
	return fmt.Sprintf(txt, NowTimestamp)
}

func WriteFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(fmt.Sprintf("Test setup, unable to write file %s : %s", path, err.Error()))
	}
}

// toString convert interface from default json unmarchal to string
func toString(value interface{}) string {
	out := ""
	// Handle the value conversion based on unmarchal default types
	switch v := value.(type) {
	case string:
		out = v
	case float64:
		out = fmt.Sprint(v)
	default:
		// Skip, unsuported
	}
	return out
}

// JsonListExtractFields Extract a field from a json entity
func JsonListExtractFields(js string, field, filterField, filterRegex string) ([]string, error) {
	var es []interface{}
	var out []string

	err := json.Unmarshal([]byte(js), &es)
	if err != nil {
		return nil, err
	}

	for _, e := range es {
		// Cast our map from default json unmarchal
		m := e.(map[string]interface{})

		value := toString(m[field])

		// Filter
		if filterField != "" {
			fv := toString(m[filterField])
			re := regexp.MustCompile(filterRegex)
			if re.MatchString(fv) {
				out = append(out, value)
			}
		} else {
			out = append(out, value)
		}
	}

	return out, nil
}

func JsonListFindObjectValue(list []map[string]interface{}, key, value string) bool {
	for _, item := range list {
		if item[key] == value {
			return true
		}
	}

	return false
}

// WriteTempFile will write the content to a temporary file and
// return the file path.
func WriteTempFile(content string) (string, error) {
	file, err := os.CreateTemp("", "cli-test-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return "", fmt.Errorf("failed to write to temp file '%s': %v", file.Name(), err)
	}

	return file.Name(), nil
}

// randomCanonical will add 4 random letter after the baseName
func randomCanonical(baseName string) string {
	var size = 4
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return baseName + "-" + string(b)
}

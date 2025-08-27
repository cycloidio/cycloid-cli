package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/stretchr/testify/assert"
	"go.yaml.in/yaml/v4"
)

func TestGetCmd(t *testing.T) {
	var baseURL = fmt.Sprintf("cy://org/%s", config.Org)
	testCases := []struct {
		name   string
		args   []string
		expect func(t *testing.T, actual string)
	}{
		{
			"getSSHKeyOk",
			[]string{"uri", "get", baseURL + "/cred/local-git?key=.raw.ssh_key"},
			func(t *testing.T, actual string) {
				assert.Equal(t, strings.Join([]string{
					"-----BEGIN OPENSSH PRIVATE KEY-----",
					"b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW",
					"QyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY",
					"8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A",
					"AAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V",
					"25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==",
					"-----END OPENSSH PRIVATE KEY-----",
				}, "\n"), actual, "ssh key must match")
			},
		},
		{
			"getSSHKeyLPadOK",
			[]string{"uri", "get", baseURL + "/cred/local-git?key=.raw.ssh_key&indent=4"},
			func(t *testing.T, actual string) {
				assert.Equal(t, strings.Join([]string{
					"    -----BEGIN OPENSSH PRIVATE KEY-----",
					"    b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW",
					"    QyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY",
					"    8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A",
					"    AAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V",
					"    25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==",
					"    -----END OPENSSH PRIVATE KEY-----",
				}, "\n"), actual, "ssh key must be indented")
			},
		},
		{
			"getSSHKeyNLPadOK",
			[]string{"uri", "get", baseURL + "/cred/local-git?key=.raw.ssh_key&nindent=4"},
			func(t *testing.T, actual string) {
				assert.Equal(t, strings.Join([]string{
					"",
					"    -----BEGIN OPENSSH PRIVATE KEY-----",
					"    b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW",
					"    QyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY",
					"    8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A",
					"    AAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V",
					"    25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==",
					"    -----END OPENSSH PRIVATE KEY-----",
				}, "\n"), actual, "ssh key must be indented after newline")
			},
		},
		{
			"getJSONListOK",
			[]string{"uri", "get", baseURL + "/projects"},
			func(t *testing.T, actual string) {
				var projects []*models.Project
				err := json.Unmarshal([]byte(actual), &projects)
				assert.NoError(t, err, "failed to parse JSON response:", actual)
			},
		},
		{
			"testYAMLOutput",
			[]string{"uri", "get", baseURL + "/projects?yaml"},
			func(t *testing.T, actual string) {
				var projects []*models.Project
				err := yaml.Unmarshal([]byte(actual), &projects)
				assert.NoError(t, err, "failed to parse YAML response:", actual)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cmdOut, cmdErr := executeCommand(testCase.args)
			if cmdErr != nil {
				t.Fatalf("failed to get from urls: %v: %s", testCase.args, cmdErr)
			}

			testCase.expect(t, cmdOut)
		})
	}
}

func TestE2e(t *testing.T) {
	var baseURL = fmt.Sprintf("cy://org/%s", config.Org)
	t.Run("TestInterpolateStdinOk", func(t *testing.T) {
		stdin := strings.Join([]string{
			"ssh: | " + baseURL + "/cred/local-git?key=.raw.ssh_key&nindent=2",
			"json: '" + baseURL + "/cred/local-git?key=.raw&json'",
		}, "\n")
		args := []string{"uri", "interpolate"}
		cmdOut, cmdErr, err := executeCommandStdin(stdin, args)
		assert.NoError(t, err, "cmd should not fail")
		assert.Empty(t, cmdErr, "stderr should be empty")
		assert.Equal(t, "ssh: | \n  -----BEGIN OPENSSH PRIVATE KEY-----\n  b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\n  QyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY\n  8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A\n  AAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V\n  25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==\n  -----END OPENSSH PRIVATE KEY-----\njson: '{\n  \"ssh_key\": \"-----BEGIN OPENSSH PRIVATE KEY-----\\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\\nQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY\\n8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A\\nAAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V\\n25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==\\n-----END OPENSSH PRIVATE KEY-----\"\n}'\n", cmdOut, "the output should match expectation")
	})
}

func TestInterpolateCmd(t *testing.T) {
	sampleContent := fmt.Sprintf("ssh_key: cy://organizations/%s/credentials/%s?key=.canonical", config.Org, config.ConfigRepo.CredentialCanonical)
	sampleExpect := fmt.Sprintf("ssh_key: %s", config.ConfigRepo.CredentialCanonical)

	t.Run("InPlaceInterpolationOk", func(t *testing.T) {
		tempDir := t.TempDir()
		sampleFilePath := tempDir + "/sample.yml"
		err := os.WriteFile(sampleFilePath, []byte(sampleContent), 0666)
		if err != nil {
			t.Logf("test setup failed, cannot write sample file at path %q: %v", tempDir, err)
			t.FailNow()
		}

		args := []string{
			"uri", "interpolate", "--in-place", sampleFilePath,
		}
		cmdOut, cmdErr := executeCommand(args)
		assert.NoError(t, cmdErr, "the cmd should not fail")
		assert.Empty(t, cmdOut, "valid in-place interpolation returns nothing")
		fileContent, err := os.ReadFile(sampleFilePath)
		assert.NoError(t, err, "we should be able to read the file after interpolation")
		assert.Equal(t, sampleExpect, string(fileContent))
	})

	t.Run("RecurseInterpolationOk", func(t *testing.T) {
		// setup
		tempDir := t.TempDir()
		nestedDir := tempDir + "/nested"
		ignoreDir := tempDir + "/ignore-me"
		sampleFilePath := tempDir + "/sample.yml"
		defaultDotGitDir := tempDir + "/.git"
		nestedSampleFilePath := nestedDir + "/nested.yml"
		ignoreSampleFilePath := ignoreDir + "/ignored.yml"
		dotGitFile := defaultDotGitDir + "/somegitfile"

		for _, dir := range []string{tempDir, nestedDir, ignoreDir, defaultDotGitDir} {
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				t.Logf("test setup failed, create nested dir %q: %v", dir, err)
				t.FailNow()
			}
		}

		for _, file := range []string{sampleFilePath, nestedSampleFilePath, ignoreSampleFilePath, dotGitFile} {
			err := os.WriteFile(file, []byte(sampleContent), 0666)
			if err != nil {
				t.Logf("test setup failed, cannot write sample file at path %q: %v", file, err)
				t.FailNow()
			}
		}
		// end setup

		args := []string{
			"uri", "interpolate", "--in-place", "--recurse", tempDir, "--ignore", "ignore-me*", "--ignore", ".git",
		}
		cmdOut, cmdErr := executeCommand(args)
		assert.NoError(t, cmdErr, "the cmd", args, "should not fail")
		assert.Empty(t, cmdOut, "valid in-place interpolation returns nothing")

		// check top level
		fileContent, err := os.ReadFile(sampleFilePath)
		assert.NoError(t, err, "we should be able to read the file after interpolation")
		assert.Equal(t, sampleExpect, string(fileContent), "this file must be interpolated")

		// check nested
		fileContent, err = os.ReadFile(nestedSampleFilePath)
		assert.NoError(t, err, "we should be able to read the nested file after interpolation")
		assert.Equal(t, sampleExpect, string(fileContent), "this file must be interpolated")

		// check ignore
		fileContent, err = os.ReadFile(ignoreSampleFilePath)
		assert.NoError(t, err, "we should be able to read the ignored file after interpolation")
		assert.Equal(t, sampleContent, string(fileContent), "this file must not be interpolated")

		// check default ignore
		t.Run("defaultDotGitIgnoreOk", func(t *testing.T) {
			fileContent, err = os.ReadFile(dotGitFile)
			assert.NoError(t, err, "we should be able to read the ignored file after interpolation")
			assert.Equal(t, sampleContent, string(fileContent), ".git file must not be interpolated")
		})

		t.Run("StdoutRecuseInterpolationOk", func(t *testing.T) {
			args := []string{
				"uri", "interpolate", "--recurse", tempDir,
			}
			cmdOut, cmdErr := executeCommand(args)
			assert.NoError(t, cmdErr, "the cmd should not fail")
			assert.Equal(t, "ssh_key: local-git\nssh_key: local-git\nssh_key: local-git\n", cmdOut, "we should get both file content")
		})
	})

	t.Run("StdoutInterpolationOk", func(t *testing.T) {
		args := []string{
			"uri", "interpolate",
		}
		cmdOut, cmdErr, err := executeCommandStdin(sampleContent, args)
		assert.NoError(t, err, "cmd should not err")
		assert.Empty(t, cmdErr, "cmd should not output on stderr if no err")
		assert.Equal(t, sampleExpect+"\n", cmdOut, "output should match interpolation")
	})

}

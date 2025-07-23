package cyargs_test

import (
	"os"
	"strings"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	org       = "myOrg"
	project   = "myProject"
	env       = "myEnv"
	component = "myComponent"
	v         = viper.GetViper()
)

func parseArgFunc(t *testing.T, cmd *cobra.Command, args []string) {
	t.Logf("Input args: %v", args)
	t.Logf("cmd persistent flags: %s", litter.Sdump(cmd.PersistentFlags()))
	if !cmd.HasPersistentFlags() {
		t.Fatal("cmd does not have the org persistent flag")
	}

	err := cmd.ParseFlags(args)
	if err != nil {
		t.Fatalf("Failed to parse flags '%v' with cmd: %v", args, err)
	}

	orgArg, projectArg, envArg, componentArg, err := cyargs.GetCyContext(cmd)
	if err != nil {
		t.Fatal("failed to parse args for cy context:", err)
	}

	assert.Equal(t, org, orgArg)
	assert.Equal(t, project, projectArg)
	assert.Equal(t, env, envArg)
	assert.Equal(t, component, componentArg)
}

func TestCyContext(t *testing.T) {
	// This test does not work as intended, remove or fix later
	t.Skip()

	// Add the same viper setup that our root cmd
	v.SetEnvPrefix("CY")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	// Add the org flag that is normally declared in the root org
	cmd := cobra.Command{}
	cmd.PersistentFlags().String("org", "", "Specify the org to use. override CY_ORG env var. Required for all Org scoped endpoint.")
	v.BindPFlag("org", cmd.PersistentFlags().Lookup("org"))
	cyargs.AddCyContext(&cmd)

	t.Run("WithFullFlag", func(t *testing.T) {
		parseArgFunc(t, &cmd, []string{
			"--org", org,
			"--project", project,
			"--env", env,
			"--component", component,
		},
		)
	})

	t.Run("WithShortFlag", func(t *testing.T) {
		parseArgFunc(t, &cmd, []string{
			"--org", org,
			"-p", project,
			"-e", env,
			"-c", component,
		},
		)
	})

	t.Run("WithEnvVar", func(t *testing.T) {
		os.Setenv("CY_ORG", org)
		os.Setenv("CY_PROJECT", project)
		os.Setenv("CY_ENV", env)
		os.Setenv("CY_COMPONENT", component)
		parseArgFunc(t, &cmd, []string{})
	})
}

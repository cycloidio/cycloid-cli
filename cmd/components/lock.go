package components

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var errComponentArgAndFlagDiffer = errors.New("the component given as argument and with --component differ")

func NewLockComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock [canonical]",
		Args:  cobra.MatchAll(cyargs.RequireArgsOrFlag("component"), cobra.MaximumNArgs(1)),
		Short: "Lock the StackForms configuration of a component.",
		Long: `Lock the StackForms configuration of a component.

While a component is locked, nobody can change its configuration or its stack version until it
is unlocked. Renaming, deleting and the other component actions keep working. Locking a locked
component with another reason replaces its reason.`,
		Example: `
	# lock a component
	cy --org my-org component lock -p my-proj -e my-env -c my-comp

	# lock a component with a reason
	cy --org my-org component lock -p my-proj -e my-env my-comp --reason "frozen for audit"
`,
		RunE: lockComponent,
	}
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	cmd.Flags().String("reason", "", "the reason of the lock, shown to whoever opens the configuration of the component")
	cyout.RegisterModel(cmd, models.Component{})
	return cmd
}

func NewUnlockComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock [canonical]",
		Args:  cobra.MatchAll(cyargs.RequireArgsOrFlag("component"), cobra.MaximumNArgs(1)),
		Short: "Unlock the StackForms configuration of a component.",
		Example: `
	# unlock a component
	cy --org my-org component unlock -p my-proj -e my-env -c my-comp
`,
		RunE: unlockComponent,
	}
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	cyout.RegisterModel(cmd, models.Component{})
	return cmd
}

func lockComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := getLockTarget(cmd, args)
	if err != nil {
		return err
	}

	reason, err := cmd.Flags().GetString("reason")
	if err != nil {
		return err
	}

	m := apiclient.NewAPIClient(common.NewAPI())
	c, _, err := m.LockComponent(org, project, env, component, reason)
	return cyout.PrintWithOptions(cmd, c, err, "failed to lock component '"+component+"'", componentTableOptions)
}

func unlockComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := getLockTarget(cmd, args)
	if err != nil {
		return err
	}

	m := apiclient.NewAPIClient(common.NewAPI())
	c, _, err := m.UnlockComponent(org, project, env, component)
	return cyout.PrintWithOptions(cmd, c, err, "failed to unlock component '"+component+"'", componentTableOptions)
}

// getLockTarget reads the component to lock or unlock, from the positional arg or the flag
func getLockTarget(cmd *cobra.Command, args []string) (org, project, env, component string, err error) {
	org, err = cyargs.GetOrg(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	project, err = cyargs.GetProject(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	env, err = cyargs.GetEnv(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	if len(args) == 0 {
		component, err = cyargs.GetComponent(cmd)
		if err != nil {
			return "", "", "", "", err
		}
		return org, project, env, component, nil
	}

	component, err = cyargs.GetComponentOrEmpty(cmd)
	if err != nil {
		return "", "", "", "", err
	}
	if component != "" && component != args[0] {
		return "", "", "", "", errComponentArgAndFlagDiffer
	}
	component = args[0]
	return org, project, env, component, nil
}

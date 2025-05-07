package cy_args

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddCyContext(cmd *cobra.Command) {
	AddProjectFlag(cmd)
	AddEnvFlag(cmd)
	AddComponentFlag(cmd)
}

var (
	v = viper.GetViper()
)

// GetOrg will return the current org using the env_var < flag precedence
//
//	func AddOrgFlag(cmd *cobra.Command) {
//		cmd.PersistentFlags().StringP("org", "o", "", "the org canonical.")
//		v.BindPFlag("org", cmd.PersistentFlags().Lookup("org"))
//		v.BindEnv("CY_ORG")
//	}
func AddProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("project", "p", "", "the project canonical, can also be set with the CY_PROJECT env var")
	v.BindPFlag("project", cmd.Flags().Lookup("project"))
}

func AddEnvFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("env", "e", "", "the env canonical, can also be set with the CY_ENV env var")
	v.BindPFlag("env", cmd.Flags().Lookup("env"))
}

func AddComponentFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "the component canonical, can also be set with the CY_COMPONENT env var")
	v.BindPFlag("component", cmd.Flags().Lookup("component"))
}

func AddColorFlag(cmd *cobra.Command) {
	cmd.Flags().String("color", "", "set the color.")
}

func AddIconFlag(cmd *cobra.Command) {
	cmd.Flags().String("icon", "", "set the icon.")
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

func GetOrg(cmd *cobra.Command) (string, error) {
	org := v.GetString("org")
	if org == "" {
		return "", fmt.Errorf("org is not set, use --org flag or CY_ORG env var, value: %s", org)
	}
	return org, nil
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

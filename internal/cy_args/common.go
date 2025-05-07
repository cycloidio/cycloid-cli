package cy_args

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddCyContext(cmd *cobra.Command) {
	AddProjectFlag(cmd)
	AddEnvFlag(cmd)
	AddComponentFlag(cmd)
}

// GetOrg will return the current org using the env_var < flag precedence
func GetOrg(cmd *cobra.Command) (org string, err error) {
	org = viper.GetString("org")
	if org == "" {
		return "", errors.New("org is not set, use --org flag or CY_ORG env var")
	}

	return org, nil
}

func AddProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("project", "p", "", "the project canonical, can also be set with the CY_PROJECT env var")
	viper.BindPFlag("project", cmd.Flags().Lookup("project"))
}

func AddEnvFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("env", "e", "", "the env canonical, can also be set with the CY_ENV env var")
	viper.BindPFlag("env", cmd.Flags().Lookup("env"))
}

func AddComponentFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "the component canonical, can also be set with the CY_COMPONENT env var")
	viper.BindPFlag("component", cmd.Flags().Lookup("component"))
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
		return "", "", "", "", err
	}

	env, err = GetEnv(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	component, err = GetComponent(cmd)
	if err != nil {
		return "", "", "", "", err
	}

	return org, project, env, component, nil
}

func GetProject(cmd *cobra.Command) (project string, err error) {
	project = viper.GetString("project")
	if project == "" {
		return "", errors.New("project is not set, use --project flag or CY_PROJECT env var")
	}

	return project, nil
}

func GetEnv(cmd *cobra.Command) (env string, err error) {
	env = viper.GetString("env")
	if env == "" {
		return "", errors.New("env is not set, use --env flag or CY_ENV env var")
	}

	return env, nil
}

func GetComponent(cmd *cobra.Command) (component string, err error) {
	component = viper.GetString("component")
	if component == "" {
		return "", errors.New("component is not set, use --component flag or CY_COMPONENT env var")
	}

	return component, nil
}

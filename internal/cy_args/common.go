package cy_args

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	v = viper.GetViper()
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
	cmd.Flags().String("color", "blue", "set the color.")
}

func GetColor(cmd *cobra.Command) (string, error) {
	color, err := cmd.Flags().GetString("color")
	if err != nil {
		return "", nil
	}

	return color, nil
}

func AddIconFlag(cmd *cobra.Command) {
	cmd.Flags().String("icon", "folder_open", "set the icon.")
}

func GetIcon(cmd *cobra.Command) (string, error) {
	icon, err := cmd.Flags().GetString("icon")
	if err != nil {
		return "", nil
	}

	return icon, nil
}

func AddOwnerFlag(cmd *cobra.Command) {
	cmd.Flags().String("owner", "", "canonical of a user to set as owner, will be the user attached to the current api key if empty.")
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
		return "", fmt.Errorf("org is not set, use --org flag or CY_ORG env var, value: %s", org)
	}
	return org, nil
}

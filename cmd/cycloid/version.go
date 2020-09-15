package root

import (
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  version,
	}
	return cmd

}

func version(cmd *cobra.Command, args []string) error {
	// api := NewAPI()
	// m := middleware.NewMiddleware(api)
	//
	// org, err := cmd.Flags().GetString("org")
	// if err != nil {
	// 	return err
	// }
	//
	// d, err := m.ListProjects(org)
	//
	// for _, pr := range d {
	// 	fmt.Printf("cannonical: %s    svcat: %s    name: %s  \n", *pr.Canonical, pr.ServiceCatalogName, *pr.Name)
	// }
	// fmt.Println(d)
	// fmt.Printf("%+v\n", err)
	return nil
}

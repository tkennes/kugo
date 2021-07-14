package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)

var AllResourcesCmd = &cobra.Command{
	Use:   "all",
	Short: "List all resources",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		fmt.Println(src.GetAllResources(namespace))
	},
}

func init() {
	// Top Level Containers
	var namespace string
	RootCmd.AddCommand(AllResourcesCmd)
	AllResourcesCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")
}

package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)

var TolerationsCmd = &cobra.Command{
	Use:   "tolerations",
	Short: "List Tolerations",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.Table(src.GetAndShowTolerations(namespace), src.TolerationsHeaders)
	},
}


func init() {
	var namespace string

	// Top Levels: containers, container
	RootCmd.AddCommand(TolerationsCmd)
	TolerationsCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")
}

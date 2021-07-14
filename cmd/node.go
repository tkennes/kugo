package kugo_cmd

import (
	"github.com/spf13/cobra"

	src "github.com/tkennes/kugo/src"
)

var NodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "List nodes",
	Run: func(cmd *cobra.Command, args []string) {
		src.Table(src.GetAndShowNodes(), src.NodeHeaders)
	},
}


func init() {
	// Top Levels: nodes, node
	RootCmd.AddCommand(NodesCmd)
}

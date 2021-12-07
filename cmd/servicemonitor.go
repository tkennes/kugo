package kugo_cmd

import (
	"github.com/spf13/cobra"

	src "github.com/tkennes/kugo/src"
)

var ServicedMonitorCmd = &cobra.Command{
	Use:   "servicemonitor",
	Short: "service monitor entry",
}

var ServicedMonitorsCmd = &cobra.Command{
	Use:   "list",
	Short: "List service monitors",
	Run: func(cmd *cobra.Command, args []string) {
		src.Table(src.GetAndShowServiceMonitors(), src.ServiceMonitorHeaders)
	},
}

var ServicedMonitorRelabelingCmd = &cobra.Command{
	Use:   "relabelings",
	Short: "List Relabelings",
	Run: func(cmd *cobra.Command, args []string) {
		src.Table(src.GetAndShowServiceMonitorsRelabelings(), src.ServiceMonitorRelabelingsHeaders)
	},
}



func init() {
	// Top Levels: nodes, node
	RootCmd.AddCommand(ServicedMonitorCmd)
	ServicedMonitorCmd.AddCommand(ServicedMonitorsCmd)
	ServicedMonitorCmd.AddCommand(ServicedMonitorRelabelingCmd)
}

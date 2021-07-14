package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)

var VolumeCmd = &cobra.Command{
	Use:   "volumes",
	Short: "List Volumes",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.Table(src.GetAndShowVolumes(namespace), src.VolumesHeaders)
	},
}


func init() {
	var namespace string

	// Top Levels: containers, container
	RootCmd.AddCommand(VolumeCmd)
	VolumeCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")

}

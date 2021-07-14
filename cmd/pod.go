package kugo_cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Containers Entrypoint",
}

var ContainersCmd = &cobra.Command{
	Use:   "containers",
	Short: "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.Table(src.GetAndShowContainers(namespace), src.ContainerHeaders)
	},
}

var ContainerEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "List Env",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.Table(src.GetAndShowContainerEnvs(namespace), src.ContainerEnvHeaders)
	},
}

var ContainerPortCmd = &cobra.Command{
	Use:   "port",
	Short: "List Ports",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.Table(src.GetAndShowContainerPorts(namespace), src.ContainerPortHeaders)
	},
}

func init() {
	var namespace string

	// Top Levels: containers, container
	RootCmd.AddCommand(ContainersCmd)
	RootCmd.AddCommand(ContainerCmd)
	ContainersCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")

	// Sub level container: env, port
	ContainerCmd.AddCommand(ContainerEnvCmd)
	ContainerEnvCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")

	ContainerCmd.AddCommand(ContainerPortCmd)
	ContainerPortCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")
}

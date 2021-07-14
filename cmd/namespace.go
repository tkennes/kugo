package kugo_cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)

var NamespaceCmd = &cobra.Command{
	Use:   "n",
	Short: "Namespace Entrypoint",
}

var NamespacesCmd = &cobra.Command{
	Use:   "ns",
	Short: "Show all Namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		src.Table(src.GetAndShowNamespaces(), src.NamespacesHeaders)
	},
}

var GetCurrentNamespaceCmd = &cobra.Command{
	Use:   "c",
	Short: "Show Current Namespace",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(src.GetCurrentNamespace())
	},
}

var SetCurrentNamespaceCmd = &cobra.Command{
	Use:   "sc",
	Short: "Set Current Namespace",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.SetCurrentnamespace(namespace)
	},
}

func init() {
	var namespace string

	// Top Levels: namespaces, namespace
	RootCmd.AddCommand(NamespaceCmd)
	RootCmd.AddCommand(NamespacesCmd)

	// Sub level container: env, port
	NamespaceCmd.AddCommand(GetCurrentNamespaceCmd)
	NamespaceCmd.AddCommand(SetCurrentNamespaceCmd)
	SetCurrentNamespaceCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")
}

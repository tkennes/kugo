package kugo_cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)

var IngressCmd = &cobra.Command{
	Use:   "ingress",
	Short: "Get Branch Info",
}

var ListIngressRulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "List ingress rules",
	Run: func(cmd *cobra.Command, args []string) {
		namespace := ParseArgument(cmd, "namespace")
		src.Table(src.GetAndShowIngress(namespace), src.IngressTableHeaders)
	},
}

func init() {
	// Top-level: Project
	RootCmd.AddCommand(IngressCmd)

	// Sub-level Ingress
	var namespace string
	IngressCmd.AddCommand(ListIngressRulesCmd)
	ListIngressRulesCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "The name of the namespace")
}

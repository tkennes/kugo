package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/kugo/src"
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func ParseArgument(cmd *cobra.Command, argument_name string) string {
	value, err := cmd.Flags().GetString(argument_name)
	if err != nil {
		src.ErrorLog(err)
	}
	return value
}

func ParseBoolArgument(cmd *cobra.Command, argument_name string) bool {
	value, err := cmd.Flags().GetBool(argument_name)
	if err != nil {
		src.ErrorLog(err)
	}
	return value
}

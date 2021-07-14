package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kugo",
	Short: "Parse kubernetes json resources",
}

package cmd

import (
	"github.com/spf13/cobra"
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo 后端API",
	Long:  "demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

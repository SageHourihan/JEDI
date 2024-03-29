package cmd

import (
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use: "validate",
	Short: "Validate EDI.",
	Long: "Validate if the EDI follows the correct X12 format.",
	Run: func(cmd *cobra.Command, args []string){
		path, _ := cmd.Flags().GetString("file")
		validate(path)
	},
}

func init() {
    validateCmd.Flags().String("file", "", "File path")
	validateCmd.MarkPersistentFlagRequired("file")
}
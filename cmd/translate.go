package cmd

import (
	"github.com/spf13/cobra"
)

var translateCmd = &cobra.Command{
	Use: "translate",
	Short: "Translate EDI.",
	Long: "Translate EDI to JSON.",
	Run: func(cmd *cobra.Command, args []string){
		path, _ := cmd.Flags().GetString("file")
		cmd.Printf("Translated. %s", path)
		translate(path)
	},
}

func init() {
	translateCmd.Flags().String("file", "", "File path")
	translateCmd.MarkPersistentFlagRequired("file")
}
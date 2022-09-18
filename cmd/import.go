package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import OTP so that codes can be generated",
	Long: `There are several ways to import OTP codes. They are specified in the table below

` + importHelpTable,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")
		fmt.Println(importHelpTable)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

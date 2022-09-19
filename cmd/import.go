package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var importShortHelp = "Import OTP so that codes can be generated by otp-cli \n\n"

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: importShortHelp,
	Long: importShortHelp + `The import commands take two arguments, the TYPE that you are importing and the PATH. The TYPE is the OTP input format and the PATH is where you can find the type to import. The details for the types and their paths can be found in the table below:

` + importHelpTable + " \nNote that you should only need to use the import command when initalising. To add new otp-codes, use the \"add\" command instead",
	Args:    cobra.ExactArgs(2),
	Example: "  otp-cli import TYPE PATH",
	RunE: func(cmd *cobra.Command, args []string) error {

		importTypeParsed := args[0]
		importPathParsed := args[1]

		// parse the individual types. should match what is found in cmd/const.go in validImportTypeArray
		switch importTypeParsed {
		case "otpauth-uri":
			fmt.Println("found otpauth-uri")
		case "otpauth-file":
			fmt.Println("found otpauth-file")
		default:
			errorMessage := "invalid import type, the valid types are below: \n\n" + importHelpTable
			return errors.New(errorMessage)
		}
		// fmt.Println(lib.Apple)
		// errorString := fmt.Sprintf("for the type %s, we were unable to correctly parse %s. Please check the type and path and try again", importTypeParsed, importPathParsed)
		// fmt.Println(errorString)
		return nil
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

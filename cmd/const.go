package cmd

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/term"
)

func generateTable(validImportTypeArray []string, validImportTypeArrayDescription []string) string {

	var tableString = &strings.Builder{}
	var table = tablewriter.NewWriter(tableString)
	var tableHeaders = []string{
		"Valid Type",
		"Description",
	}

	var terminalFd = int(os.Stdin.Fd())
	width, _, _ := term.GetSize(terminalFd)

	table.SetAutoWrapText(false)
	// https://github.com/olekukonko/tablewriter/issues/86
	// tablewriter.WrapString()
	table.SetHeader(tableHeaders)
	for i := 0; i < len(validImportTypeArray); i++ {
		table.Append([]string{
			validImportTypeArray[i],
			validImportTypeArrayDescription[i],
		})
	}

	table.SetColWidth((width / 2))
	table.Render()
	return tableString.String()
}

var validImportTypeArray = []string{
	"otpauth-uri",
	"otpauth-file",
}
var validImportTypeArrayDescription = []string{
	"A single otpauth URI",
	"A path to a file where each line contains an otpauth URI.",
}

// a table that contains the valid types for import
var importHelpTable = generateTable(validImportTypeArray, validImportTypeArrayDescription)

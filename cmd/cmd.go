package cmd

import (
	"fmt"
	"os"

	"github.com/appno/stripe/schema"
	"github.com/spf13/cobra"
)

var filePath string

var part1Example = `
stripe part1 -f $HOME/data.json

stripe part1 \
'{
		"first_name": "Violet",
		"last_name": "Beauregarde",
		"id": "12345",
		"tax_id": "111-22-3333",
		"address": {
				"street": "111 W Third",
				"city": "Chicago",
				"state": "IL",
				"postal_code": "60619",
				"country": "US"
		}
}'`

// Part1Cmd : Run Part 1 of the coding challenge
var Part1Cmd = &cobra.Command{
	Use:     "part1 [DATA | -f FILE]",
	Short:   "stripe coding challenge part 1",
	Args:    cobra.RangeArgs(0, 1),
	Example: part1Example,
	Run: func(cmd *cobra.Command, args []string) {
		str, err := readInput(args, 0, filePath)
		if err != nil {
			os.Exit(1)
		}

		data, err := unmarshalInput(str)
		if err != nil {
			os.Exit(1)
		}

		validator := schema.DefaultValidator()
		result, err := validator.IsCompliant(data)

		if err != nil {
			os.Exit(1)
		}

		json, err := result.JSONString()

		if err != nil {
			os.Exit(1)
		}

		fmt.Println(json)
	},
}

var part2Example = `
stripe part2 -f $HOME/data.json

stripe part2 \
'{
		"first_name": "Violet",
		"last_name": "Beauregarde",
		"id": "12345",
		"tax_id": "111-22-3333",
		"address": {
				"street": "111 W Third",
				"city": "Chicago",
				"state": "IL",
				"postal_code": "60619",
				"country": "US"
		}
}'`

// Part2Cmd : Run Part 2 of the coding challenge
var Part2Cmd = &cobra.Command{
	Use:     "part2 [DATA | -f FILE]",
	Short:   "stripe coding challenge part 2",
	Args:    cobra.RangeArgs(0, 1),
	Example: part2Example,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readInput(args, 0, filePath)
		if err != nil {
			fmt.Printf("1: ERR: %+v\n", err)
			os.Exit(1)
		}

		document, err := schema.DocumentFromBytes(input)

		if err != nil {
			fmt.Printf("2: ERR: %+v\n", err)
			os.Exit(1)
		}

		compliance := document.GetPastDueCompliance()

		json, err := compliance.JSONString()

		if err != nil {
			fmt.Printf("3: ERR: %+v\n", err)
			os.Exit(1)
		}

		fmt.Println(json)
	},
}

// MainCmd : Run main application
var MainCmd = &cobra.Command{
	Use:   "stripe",
	Short: "stripe coding demo",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("Stripe coding demo.\n")
	},
}

// ConfigCmd : Display configuration variables
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "display application configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		// fmt.Printf("Stripe coding demo.\n")
		fmt.Println(schema.GetConfigString())
	},
}

// Execute : Execute command
func Execute() {
	Part1Cmd.Flags().StringVarP(&filePath, "file", "f", "", "project data file")
	Part2Cmd.Flags().StringVarP(&filePath, "file", "f", "", "project data file")

	MainCmd.AddCommand(Part1Cmd)
	MainCmd.AddCommand(Part2Cmd)
	MainCmd.AddCommand(ServerCmd)
	MainCmd.AddCommand(ConfigCmd)

	if err := MainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package day00

import (
	"fmt"

	"github.com/spf13/cobra"
)

const day = "0"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "1":
		fmt.Printf("Solve day %s part 1", day)
	case "2":
		fmt.Printf("Solve day %s part 2", day)
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

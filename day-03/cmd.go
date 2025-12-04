package day03

import (
	"fmt"
	"strconv"

	"github.com/nosnikliw/advent2025/input"
	"github.com/spf13/cobra"
)

const day = "3"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	banks, err := input.ReadLines(args[1])
	cobra.CheckErr(err)
	switch args[0] {
	case "1":
		fmt.Printf("Total: %d\n", sumJoltages(banks, maxJoltage))
	case "2":
		fmt.Printf("Total: %d\n", sumJoltages(banks, maxJoltage2))
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func sumJoltages(banks []string, getJoltage func(string) int64) int64 {
	var total int64 = 0
	for _, bank := range banks {
		total += getJoltage(bank)
	}
	return total
}

func maxJoltage2(bank string) int64 {
	joltage := maxJoltageExt(bank, 12, 0)
	return joltage
}

func maxJoltageExt(bank string, cells int, current int64) int64 {
	if cells <= 2 {
		return 100*current + maxJoltage(bank)
	}
	pos := 0
	var max int64 = 0
	for i, v := range bank[:len(bank)-cells+1] {
		val, err := strconv.ParseInt(string(v), 10, 64)
		cobra.CheckErr(err)
		if val > max {
			pos = i
			max = val
		}
	}
	return maxJoltageExt(bank[pos+1:], cells-1, 10*current+max)
}

func maxJoltage(bank string) int64 {
	var max int64 = 0
	maxPos := 0
	for i, v := range bank[:len(bank)-1] {
		val, err := strconv.ParseInt(string(v), 10, 64)
		cobra.CheckErr(err)
		if val > max {
			max = val
			maxPos = i
		}
	}
	var max2 int64 = 0
	for _, v := range bank[maxPos+1:] {
		val, err := strconv.ParseInt(string(v), 10, 64)
		cobra.CheckErr(err)
		if val > max2 {
			max2 = val
		}
	}
	joltage := max*10 + max2
	return joltage
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

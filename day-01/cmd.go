package day01

import (
	"fmt"
	"strconv"

	"github.com/nosnikliw/advent2025/input"
	"github.com/spf13/cobra"
)

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "day-1 [part] [input file]",
	Short: "Solve day 1",
	Long:  `Solve day 1 with the given input file`,
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	dial := getDial(args[1])
	switch args[0] {
	case "1":
		fmt.Printf("Password: %d\n", dial.GetPassword())
	case "2":
		fmt.Printf("Password: %d\n", dial.GetTempPassword())
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func getDial(inputFile string) Dial {
	lines, err := input.ReadLines(inputFile)
	cobra.CheckErr(err)

	dial := Dial{50, 0, 0}

	for _, instuction := range lines {
		dial.Rotate(instuction)
	}
	return dial
}

type Dial struct {
	position        int64
	zeroCount       int64
	passedZeroCount int64
}

func (d *Dial) Rotate(val string) {
	direction := val[0]
	distance, err := strconv.ParseInt(val[1:], 10, 64)
	cobra.CheckErr(err)
	switch direction {
	case 'L':
		d.passedZeroCount += (distance + (100 - d.position)) / 100
		if d.position == 0 {
			d.passedZeroCount -= 1
		}
		position := (d.position - distance) % 100
		if position < 0 {
			position += 100
		}
		d.position = position
	case 'R':
		d.passedZeroCount += (distance + d.position) / 100
		position := (d.position + distance) % 100
		d.position = position
	default:
		panic("invalid input")
	}
	if d.position == 0 {
		d.zeroCount += 1
	}
}

func (d *Dial) GetPassword() int64 {
	return d.zeroCount
}

func (d *Dial) GetTempPassword() int64 {
	return d.passedZeroCount
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

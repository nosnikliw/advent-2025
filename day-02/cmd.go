package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nosnikliw/advent2025/input"
	"github.com/spf13/cobra"
)

const day = "2"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	ids, err := input.ReadFile(args[1])
	cobra.CheckErr(err)
	switch args[0] {
	case "1":
		result, err := sumInvalid(ids, isInvalid)
		cobra.CheckErr(err)
		fmt.Printf("Result: %d\n", result)
	case "2":
		result, err := sumInvalid(ids, isInvalid2)
		cobra.CheckErr(err)
		fmt.Printf("Result: %d\n", result)
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func sumInvalid(ids string, test func(int64) bool) (result int64, err error) {
	rangeStrings := strings.Split(ids, ",")
	for _, r := range rangeStrings {
		var min, max int64
		min, max, err = parseRange(r)
		if err != nil {
			return
		}
		for id := min; id <= max; id++ {
			if test(id) {
				result += id
			}
		}
	}
	return
}

func parseRange(val string) (min int64, max int64, err error) {
	parts := strings.Split(val, "-")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid range: %s", val)
		return
	}
	min, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return
	}
	max, err = strconv.ParseInt(parts[1], 10, 64)
	return
}

func isInvalid(intid int64) bool {
	id := fmt.Sprint(intid)
	length := len(id)
	half := length / 2
	return id[half:] == id[:half]
}

func isInvalid2(intId int64) bool {
	id := fmt.Sprint(intId)
	length := len(id)
	for repeats := length; repeats >= 2; repeats-- {
		if length%repeats != 0 {
			continue
		}
		partLength := length / repeats
		pattern := id[:partLength]
		test := ""
		for i := 0; i < repeats; i++ {
			test = test + pattern
		}
		if test == id {
			return true
		}
	}
	return false
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

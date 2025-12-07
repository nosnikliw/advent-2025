package day07

import (
	"fmt"
	"strings"

	"github.com/nosnikliw/advent2025/input"
	"github.com/nosnikliw/advent2025/sets"
	"github.com/spf13/cobra"
)

const day = "7"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	manifold, err := input.ReadLines(args[1])
	cobra.CheckErr(err)
	switch args[0] {
	case "1":
		fmt.Printf("Split count: %d\n", countSplits(manifold))
	case "2":
		fmt.Printf("Split count: %d\n", countPaths(manifold))
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func countSplits(manifold []string) int64 {
	startPos := strings.Index(manifold[0], "S")
	beams := sets.NewIntSet()
	beams.Add(int64(startPos))
	splitCount := int64(0)
	for i := 1; i < len(manifold); i++ {
		currentBeams := beams.Members()
		for _, v := range currentBeams {
			if manifold[i][v] == '^' {
				splitCount++
				beams.Add(v - 1)
				beams.Add(v + 1)
				beams.Remove(v)
			}
		}
	}
	return splitCount
}

func countPaths(manifold []string) int64 {
	startPos := strings.Index(manifold[0], "S")
	paths := map[int64]int64{}
	paths[int64(startPos)] = 1
	for i := 1; i < len(manifold); i++ {
		newPaths := map[int64]int64{}
		for pos, count := range paths {
			if manifold[i][pos] == '^' {
				add(newPaths, pos-1, count)
				add(newPaths, pos+1, count)
			} else {
				add(newPaths, pos, count)
			}
		}
		paths = newPaths
	}
	pathCount := int64(0)
	for _, v := range paths {
		pathCount += v
	}
	return pathCount
}

func add(paths map[int64]int64, position int64, count int64) {
	current, found := paths[position]
	if found {
		paths[position] = current + count
	} else {
		paths[position] = count
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

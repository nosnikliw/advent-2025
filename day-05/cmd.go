package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nosnikliw/advent2025/input"
	"github.com/spf13/cobra"
)

const day = "5"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	ranges, ids := loadInput(args[1])
	switch args[0] {
	case "1":
		fmt.Printf("Fresh: %d\n", countFreshIngredients(ranges, ids))
	case "2":
		fmt.Printf("Fresh IDs: %d\n", countFreshIngredientIDs(ranges))
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func countFreshIngredientIDs(ranges []IntRange) int {
	merged := getMergedRanges(ranges)
	count := 0
	for _, r := range merged {
		count += int(r.Size())
	}
	return count
}

func getMergedRanges(ranges []IntRange) (merged []IntRange) {
	merged = append(merged, ranges[0])
	anyMerged := false
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		wasMerged := false
		for j := 0; j < len(merged); j++ {
			m := merged[j]
			if m.Overlaps(r) {
				merged[j] = m.Merge(r)
				wasMerged = true
				anyMerged = true
				break
			}
		}
		if !wasMerged {
			merged = append(merged, r)
		}
	}
	if anyMerged {
		merged = getMergedRanges(merged)
	}
	return
}

func countFreshIngredients(ranges []IntRange, ids []int64) int {
	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r.IsInRange(id) {
				count++
				break
			}
		}
	}
	return count
}

func loadInput(path string) (ranges []IntRange, ids []int64) {
	lines, err := input.ReadLines(path)
	cobra.CheckErr(err)

	index := 0
	for _, line := range lines {
		index++
		if line == "" {
			break
		}
		r, err := parseIntRange(line)
		cobra.CheckErr(err)
		ranges = append(ranges, r)
	}
	for i := index; i < len(lines); i++ {
		id, err := strconv.ParseInt(lines[i], 10, 64)
		cobra.CheckErr(err)
		ids = append(ids, id)
	}
	return
}

func parseIntRange(val string) (result IntRange, err error) {
	parts := strings.Split(val, "-")
	if len(parts) != 2 {
		err = fmt.Errorf("unable to parse int range: %s", val)
		return
	}
	min, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return
	}
	max, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return
	}
	result = IntRange{Min: min, Max: max}
	return
}

type IntRange struct {
	Min int64
	Max int64
}

func (r *IntRange) IsInRange(val int64) bool {
	return val >= r.Min && val <= r.Max
}

func (r *IntRange) Overlaps(other IntRange) bool {
	return r.IsInRange(other.Min) || r.IsInRange(other.Max) || other.IsInRange(r.Min) || other.IsInRange(r.Max)
}

func (r *IntRange) Merge(other IntRange) IntRange {
	min := r.Min
	if min > other.Min {
		min = other.Min
	}
	max := r.Max
	if max < other.Max {
		max = other.Max
	}
	return IntRange{
		Min: min,
		Max: max,
	}
}

func (r *IntRange) Size() int64 {
	return r.Max - r.Min + 1
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

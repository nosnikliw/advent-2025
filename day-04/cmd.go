package day04

import (
	"fmt"

	"github.com/nosnikliw/advent2025/input"
	"github.com/nosnikliw/advent2025/vectors"
	"github.com/spf13/cobra"
)

const day = "4"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

func solve(cmd *cobra.Command, args []string) {
	grid := loadGrid(args[1])
	switch args[0] {
	case "1":
		fmt.Printf("Count: %d\n", grid.CountAccessible())
	case "2":
		removed, _ := grid.RemoveAccessible()
		fmt.Printf("Count: %d\n", removed)
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func loadGrid(path string) Grid {
	data, err := input.ReadLines(path)
	cobra.CheckErr(err)
	return Grid{data}
}

type Grid struct {
	grid []string
}

func (g *Grid) SizeX() int {
	return len(g.grid[0])
}

func (g *Grid) SizeY() int {
	return len(g.grid)
}

func (g *Grid) HasPaper(x int, y int) bool {
	if x < 0 || x >= g.SizeX() || y < 0 || y >= g.SizeY() {
		return false
	}
	return g.grid[y][x] == '@'
}

func (g *Grid) CountNeighbours(x int, y int) int {
	neihbours := []vectors.Vector{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 1, Y: -1},
		{X: 0, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
	}
	count := 0
	for _, v := range neihbours {
		if g.HasPaper(x+v.X, y+v.Y) {
			count++
		}
	}
	return count
}

func (g *Grid) CountAccessible() int {
	count := 0
	for i := 0; i < g.SizeX(); i++ {

		for j := 0; j < g.SizeY(); j++ {
			if !g.HasPaper(i, j) {
				continue
			}
			if g.CountNeighbours(i, j) < 4 {
				count++
			}
		}
	}
	return count
}

func (g *Grid) Print() {
	for _, row := range g.grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func (g *Grid) RemoveAccessible() (int, *Grid) {
	count := 0
	updated := [][]rune{}
	for j := 0; j < g.SizeY(); j++ {
		updated = append(updated, []rune{})
		for i := 0; i < g.SizeX(); i++ {
			if !g.HasPaper(i, j) {
				updated[j] = append(updated[j], '.')
				continue
			}
			if g.CountNeighbours(i, j) < 4 {
				updated[j] = append(updated[j], '.')
				count++
			} else {
				updated[j] = append(updated[j], '@')
			}
		}
	}
	if count == 0 {
		return count, g
	}
	newData := []string{}
	for _, row := range updated {
		newData = append(newData, string(row))
	}
	newGrid := Grid{newData}

	removed, finalGrid := newGrid.RemoveAccessible()

	return count + removed, finalGrid
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

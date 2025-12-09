package day08

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nosnikliw/advent2025/input"
	"github.com/nosnikliw/advent2025/sets"
	"github.com/nosnikliw/advent2025/vectors"
	"github.com/spf13/cobra"
)

const day = "8"

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   fmt.Sprintf("day-%s [part] [input file]", day),
	Short: fmt.Sprintf("Solve day %s", day),
	Long:  fmt.Sprintf(`Solve day %s with the given input file`, day),
	Run:   solve,
	Args:  cobra.ExactArgs(2),
}

var connectionCount int = 10

func solve(cmd *cobra.Command, args []string) {
	vectorList := loadVectors(args[1])
	switch args[0] {
	case "1":
		fmt.Printf("Result %d\n", connect(vectorList, connectionCount))
	case "2":
		fmt.Printf("Result %d\n", connectAll(vectorList))
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

type connection struct {
	Distance float64
	A        vectors.Vector3
	B        vectors.Vector3
}

func connectAll(list []vectors.Vector3) int {
	circuits := []sets.Set[vectors.Vector3]{}
	for _, v := range list {
		circuit := sets.NewSet[vectors.Vector3]()
		circuit.Add(v)
		circuits = append(circuits, *circuit)
	}
	connections := getOrderedConnections(list)
	for _, con := range connections {
		for _, circuit := range circuits {
			if circuit.Contains(con.A) {
				circuit.Add(con.B)
				break
			}
			if circuit.Contains(con.B) {
				circuit.Add(con.A)
				break
			}
		}
		circuits = mergeIntersectingSets(circuits)
		if len(circuits) == 1 {
			return int(con.A.X) * int(con.B.X)
		}
	}
	return -1
}

func getOrderedConnections(list []vectors.Vector3) []connection {
	pairs := []connection{}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			pairs = append(pairs, connection{
				A:        list[i],
				B:        list[j],
				Distance: list[i].Subtract(list[j]).Magnitude(),
			})
		}
	}
	sort.Slice(pairs, func(a, b int) bool { return pairs[a].Distance < pairs[b].Distance })
	return pairs
}

func connect(list []vectors.Vector3, connections int) int {
	pairs := getOrderedConnections(list)

	circuits := []sets.Set[vectors.Vector3]{}
	connectionsMade := 0
	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		alreadyConnected := false
		newConnection := false
		for _, circuit := range circuits {
			if circuit.Contains(p.A) && circuit.Contains(p.B) {
				alreadyConnected = true
				break
			}
			if circuit.Contains(p.A) {
				newConnection = true
				circuit.Add(p.B)
				break
			}
			if circuit.Contains(p.B) {
				newConnection = true
				circuit.Add(p.A)
				break
			}
		}
		if !alreadyConnected && !newConnection {
			newCircuit := sets.NewSet[vectors.Vector3]()
			newCircuit.Add(p.A)
			newCircuit.Add(p.B)
			circuits = append(circuits, *newCircuit)
		}
		circuits = mergeIntersectingSets(circuits)
		connectionsMade++
		if connectionsMade >= connections {
			break
		}
	}
	sort.Slice(circuits, func(a, b int) bool { return circuits[a].Count() > circuits[b].Count() })
	// for _, v := range circuits {
	// 	fmt.Println(v)
	// }
	return circuits[0].Count() * circuits[1].Count() * circuits[2].Count()
}

func mergeIntersectingSets(setList []sets.Set[vectors.Vector3]) []sets.Set[vectors.Vector3] {
	merged := []sets.Set[vectors.Vector3]{}
	anyMerged := false
	for _, s := range setList {
		wasMerged := false
		for i := 0; i < len(merged); i++ {
			m := merged[i]
			if s.Intersects(m) {
				merged[i] = *m.Union(s)
				wasMerged = true
				anyMerged = true
			}
		}
		if !wasMerged {
			merged = append(merged, s)
		}
	}
	if anyMerged {
		merged = mergeIntersectingSets(merged)
	}
	return merged
}

func loadVectors(path string) (vs []vectors.Vector3) {
	lines, err := input.ReadLines(path)
	cobra.CheckErr(err)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		v := vectors.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		}
		v.X, err = strconv.ParseFloat(parts[0], 64)
		cobra.CheckErr(err)
		v.Y, err = strconv.ParseFloat(parts[1], 64)
		cobra.CheckErr(err)
		v.Z, err = strconv.ParseFloat(parts[2], 64)
		cobra.CheckErr(err)
		vs = append(vs, v)
	}
	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	Cmd.Flags().IntVarP(&connectionCount, "connections", "c", 10, "The number of connections to make")
}

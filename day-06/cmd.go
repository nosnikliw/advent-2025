package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nosnikliw/advent2025/input"
	"github.com/spf13/cobra"
)

const day = "6"

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
		table := loadTable(args[1])
		fmt.Printf("Result: %d\n", doHomework(table))
	case "2":
		opStream := loadAsOpStream(args[1])
		fmt.Printf("Result: %d\n", processOpStream(opStream))
	default:
		fmt.Printf("Invalid part: %s\n", args[0])
	}
}

func doHomework(table [][]string) (result int64) {
	problemCount := len(table[0])
	opRow := table[len(table)-1]
	for i := 0; i < problemCount; i++ {
		op := opRow[i]
		switch op {
		case "+":
			sum := int64(0)
			for j := 0; j < len(table)-1; j++ {
				val, err := strconv.ParseInt(table[j][i], 10, 64)
				cobra.CheckErr(err)
				sum += val
			}
			result += sum
		case "*":
			prod := int64(1)
			for j := 0; j < len(table)-1; j++ {
				val, err := strconv.ParseInt(table[j][i], 10, 64)
				cobra.CheckErr(err)
				prod *= val
			}
			result += prod
		default:
			cobra.CheckErr(fmt.Errorf("unknown operator: %s", op))
		}
	}
	return
}

func loadTable(path string) (table [][]string) {
	lines, err := input.ReadLines(path)
	cobra.CheckErr(err)
	for _, row := range lines {
		table = append(table, strings.Fields(row))
	}
	return
}

func processOpStream(stream []string) (result int64) {
	operator := ""
	opResult := int64(0)
	for _, v := range stream {
		if v == "+" || v == "*" {
			result += opResult
			operator = v
			opResult = 0
			if operator == "*" {
				opResult = 1
			}
			continue
		}
		val, err := strconv.ParseInt(v, 10, 64)
		cobra.CheckErr(err)
		if operator == "*" {
			opResult *= val
		} else {
			opResult += val
		}
	}
	result += opResult
	return
}

func loadAsOpStream(path string) []string {
	lines, err := input.ReadLines(path)
	cobra.CheckErr(err)

	//assume all lines are the same length
	stream := []string{}
	for i := 0; i < len(lines[0]); i++ {
		op := ""
		val := ""
		for j := 0; j < len(lines); j++ {
			char := []rune(lines[j])[i]
			switch {
			case char == ' ':
				//do nothing
			case char == '+' || char == '*':
				op = string(char)
			default:
				val += string(char)
			}
		}
		if op != "" {
			stream = append(stream, op)
		}
		if val != "" {
			stream = append(stream, val)
		}
	}
	return stream
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
}

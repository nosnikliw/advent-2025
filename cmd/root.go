/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	day01 "github.com/nosnikliw/advent2025/day-01"
	day02 "github.com/nosnikliw/advent2025/day-02"
	day03 "github.com/nosnikliw/advent2025/day-03"
	day04 "github.com/nosnikliw/advent2025/day-04"
	day05 "github.com/nosnikliw/advent2025/day-05"
	day06 "github.com/nosnikliw/advent2025/day-06"
	day07 "github.com/nosnikliw/advent2025/day-07"
	day08 "github.com/nosnikliw/advent2025/day-08"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent2025",
	Short: "A CLI solving the advent of code 2025",
	Long:  `A CLI solving the advent of code 2025`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(day01.Cmd)
	rootCmd.AddCommand(day02.Cmd)
	rootCmd.AddCommand(day03.Cmd)
	rootCmd.AddCommand(day04.Cmd)
	rootCmd.AddCommand(day05.Cmd)
	rootCmd.AddCommand(day06.Cmd)
	rootCmd.AddCommand(day07.Cmd)
	rootCmd.AddCommand(day08.Cmd)
}

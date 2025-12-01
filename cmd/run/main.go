package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gustavLindqvist/aoc-2025/internal/days/day01"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day02"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day03"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day04"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day05"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day06"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day07"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day08"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day09"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day10"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day11"
	"github.com/gustavLindqvist/aoc-2025/internal/days/day12"
	"github.com/gustavLindqvist/aoc-2025/pkg/aoc"
)

type dayBundle struct {
	meta    aoc.Day
	example string
	input   string
}

var days = map[string]dayBundle{
	"01": {meta: day01.Day, example: day01.Example, input: day01.Input},
	"02": {meta: day02.Day, example: day02.Example, input: day02.Input},
	"03": {meta: day03.Day, example: day03.Example, input: day03.Input},
	"04": {meta: day04.Day, example: day04.Example, input: day04.Input},
	"05": {meta: day05.Day, example: day05.Example, input: day05.Input},
	"06": {meta: day06.Day, example: day06.Example, input: day06.Input},
	"07": {meta: day07.Day, example: day07.Example, input: day07.Input},
	"08": {meta: day08.Day, example: day08.Example, input: day08.Input},
	"09": {meta: day09.Day, example: day09.Example, input: day09.Input},
	"10": {meta: day10.Day, example: day10.Example, input: day10.Input},
	"11": {meta: day11.Day, example: day11.Example, input: day11.Input},
	"12": {meta: day12.Day, example: day12.Example, input: day12.Input},
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: run <day> <star> [example|input]")
	}

	dayKey := normalizeDay(os.Args[1])
	bundle, ok := days[dayKey]
	if !ok {
		log.Fatalf("unknown day %s", dayKey)
	}

	solver, starLabel := selectStar(bundle.meta, os.Args[2])
	input, inputLabel, ok := selectInput(bundle, os.Args[3:])
	if !ok {
		fmt.Printf("Day %s Star %s (%s): <no %s data>\n", dayKey, starLabel, inputLabel, inputLabel)
		return
	}

	result := solver(input)
	fmt.Printf("Day %s Star %s (%s): %v\n", dayKey, starLabel, inputLabel, result)
}

func normalizeDay(arg string) string {
	num, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatalf("invalid day %q", arg)
	}
	if num < 1 || num > 25 {
		log.Fatalf("day must be between 1 and 25, got %d", num)
	}
	return fmt.Sprintf("%02d", num)
}

func selectStar(day aoc.Day, arg string) (aoc.Solver, string) {
	switch strings.TrimPrefix(strings.ToLower(arg), "star") {
	case "1":
		return day.Star1, "1"
	case "2":
		return day.Star2, "2"
	default:
		log.Fatalf("unknown star %q", arg)
		return nil, ""
	}
}

func selectInput(bundle dayBundle, args []string) (string, string, bool) {
	if len(args) == 0 {
		if bundle.input == "" {
			return "", "input", false
		}
		return bundle.input, "input", true
	}
	selector := strings.ToLower(args[0])
	switch selector {
	case "example", "sample", "ex":
		if bundle.example == "" {
			return "", "example", false
		}
		return bundle.example, "example", true
	case "input", "real", "normal", "main":
		if bundle.input == "" {
			return "", "input", false
		}
		return bundle.input, "input", true
	default:
		log.Fatalf("unknown input selector %q", selector)
		return "", "", false
	}
}

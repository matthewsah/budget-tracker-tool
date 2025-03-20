package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const TRANSACTION_DATE = 0
const DESCRIPTION = 2
const CATEGORY = 3
const AMOUNT = 5

// func removeYear(date string) string {
// 	substrings := strings.Split(date, "/")
// 	if len(substrings) != 3 {
// 		panic("Should be 3 inputs for date!")
// 	}
// 	month := substrings[0]
// 	day := substrings[1]
// 	return fmt.Sprintf("%s/%s", month, day)
// }

func makeCamelCase(str string) string {
	substrings := strings.Split(str, " ")
	var output strings.Builder
	for substring_idx, s := range substrings {
		if len(s) > 0 {
			for idx, r := range s {
				if idx == 0 {
					output.WriteRune(unicode.ToUpper(r))
				} else {
					output.WriteRune(unicode.ToLower(r))
				}
			}
			if substring_idx < len(substrings)-1 {
				output.WriteString(" ")
			}
		}
	}
	return output.String()
}

var categoryMapping = map[string]string{
	"Automotive":            "Needs",
	"Bills & Utilities":     "Needs",
	"Education":             "Personal Investment",
	"Entertainment":         "Entertainment",
	"Fees & Adjustments":    "UNSELECTED",
	"Food & Drink":          "Dining",
	"Gas":                   "Needs",
	"Gifts & Donations":     "Gifts",
	"Groceries":             "Groceries",
	"Health & Wellness":     "Wellness",
	"Home":                  "Personal Investment",
	"Miscellaneous":         "Wants",
	"Personal":              "UNSELECTED",
	"Professional Services": "UNSELECTED",
	"Shopping":              "Wants",
	"Travel":                "Travel",
}

func mapCategory(category string) string {
	value, ok := categoryMapping[category]
	if ok {
		return value
	} else {
		fmt.Printf("Unable to map category %s", category)
		return "UNSELECTED"
	}
}

func getAmount(amount string) string {
	var neg byte = '-'
	if len(amount) > 1 && amount[0] == neg {
		return amount[1:]
	} else {
		return amount
	}
}

func run(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
	}

	var output strings.Builder

	for i := len(lines) - 1; i >= 1; i-- {
		line := lines[i]
		if len(line) != 7 {
			fmt.Printf("Line %s was not the correct length", line)
			continue
		}
		output.WriteString(fmt.Sprintf("%s,%s,%s", makeCamelCase(line[DESCRIPTION]), getAmount(line[AMOUNT]), mapCategory(line[CATEGORY])))
		output.WriteString("\n")
	}

	fmt.Print(output.String())
}

func main() {
	run("")
}

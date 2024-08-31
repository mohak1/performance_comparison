package main

import (
	"strings"
)

const FILES_DIR = "./text_files/smaller_files/" // dir where log files are present

var WORDS_TO_SEARCH = [...]string{
	"Session",
	"Warning",
	"Failed",
	"CBS",
	"CSI",
}

func perform_case_insensitive_search(
	words string,
	line string,
	tracker map[string]int,
) {
	for _, word := range words {
		var lower_case_word = strings.ToLower(string(word))
		var lower_case_line = strings.ToLower(string(line))
		tracker[string(word)] += strings.Count(lower_case_line, lower_case_word)
	}
	return
}

func main() {
	perform_case_insensitive_search()
}

package main

import (
	"log"
	"os"
	"strings"
)

const FILES_DIR = "./text_files/smaller_files/" // dir where log files are present

var WORDS_TO_SEARCH = []string{
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

func validate_path_and_words() {
	log.Print("Validating FILES_DIR...")
	info, err := os.Stat(FILES_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The location '%v' does not exist. Please "+
				"check the path and try again.", FILES_DIR)
		}
		log.Fatal("Error checking the path: %v", FILES_DIR)
	}
	log.Print("%v is a valid path", FILES_DIR)

	if info.IsDir() {
		log.Print("%v is a valid directory", FILES_DIR)
	} else {
		log.Fatal("The location '%v' is not a directory. Please "+
			"check the path and try again.", FILES_DIR)
	}

	if len(WORDS_TO_SEARCH) < 1 {
		log.Fatal("WORDS_TO_SEARCH cannot be an empty list. Please add " +
			"some words to this list and try again.")
	}
	log.Print("WORDS_TO_SEARCH list is not empty")

}


func main() {
	validate_path_and_words()
}

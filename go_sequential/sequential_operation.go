package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const FILES_DIR = "../text_files/smaller_files/" // dir where log files are present

var WORDS_TO_SEARCH = []string{
	"Session",
	"Warning",
	"Failed",
	"CBS",
	"CSI",
}

func perform_case_insensitive_search(
	words []string,
	line string,
	tracker map[string]int,
) {
	lower_case_line := strings.ToLower(line)
	for _, word := range words {
		tracker[word] += strings.Count(lower_case_line, strings.ToLower(word))
	}
}

func validate_path_and_words() {
	log.Print("Validating FILES_DIR...")
	info, err := os.Stat(FILES_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("The location '%v' does not exist. Please "+
				"check the path and try again.", FILES_DIR)
		}
		log.Fatalf("Error checking the path: %v", FILES_DIR)
	}
	log.Printf("%v is a valid path", FILES_DIR)

	if info.IsDir() {
		log.Printf("%v is a valid directory", FILES_DIR)
	} else {
		log.Fatalf("The location '%v' is not a directory. Please "+
			"check the path and try again.", FILES_DIR)
	}

	if len(WORDS_TO_SEARCH) < 1 {
		log.Fatal("WORDS_TO_SEARCH cannot be an empty list. Please add " +
			"some words to this list and try again.")
	}
	log.Print("WORDS_TO_SEARCH list is not empty")

}

func process_file(
	file_name string,
	dir_path string,
	words_to_search []string,
) {
	file_path := path.Join(dir_path, file_name)
	count_tracker := make(map[string]int)
	for _, word := range words_to_search {
		count_tracker[word] = 0
	}

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v\n", file_path, err)
	}
	defer file.Close()

	var start_time = time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		perform_case_insensitive_search(
			words_to_search,
			line,
			count_tracker,
		)
	}
	var duration = time.Since(start_time)
	fmt.Printf("Completed searching %v. Results: %v. Time taken: %.2f seconds\n",
		file_name, count_tracker, duration.Seconds())

}

func entrypoint() {
	validate_path_and_words()

	files, err := os.ReadDir(FILES_DIR)
	if err != nil {
		log.Fatalf("Error in reading files from %s; %s", FILES_DIR, err)
	}

	for _, file := range files {
		process_file(
			file.Name(),
			FILES_DIR,
			WORDS_TO_SEARCH,
		)
	}
}

func main() {
	var process_start_time = time.Now()
	entrypoint()
	var process_end_time = time.Since(process_start_time)
	log.Printf("The process took %.2f seconds", process_end_time.Seconds())
}

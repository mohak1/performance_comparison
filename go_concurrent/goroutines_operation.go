package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
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
	for _, word := range words {
		var lower_case_word = strings.ToLower(string(word))
		var lower_case_line = strings.ToLower(string(line))
		tracker[word] += strings.Count(lower_case_line, lower_case_word)
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
	wait_group *sync.WaitGroup,
) {
	defer wait_group.Done()

	var file_path = path.Join(dir_path, file_name)
	var count_tracker = make(map[string]int)
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

	var wait_group sync.WaitGroup

	for _, file := range files {
		wait_group.Add(1)
		go process_file(
			file.Name(),
			FILES_DIR,
			WORDS_TO_SEARCH,
			&wait_group,
		)
	}

	wait_group.Wait()
}

func main() {
	var CPU_CORES = runtime.NumCPU() * 12
	log.Printf("Setting GOMAXPROCS at %v", CPU_CORES)
	runtime.GOMAXPROCS(CPU_CORES)

	var process_start_time = time.Now()
	entrypoint()
	var process_end_time = time.Since(process_start_time)
	log.Printf("The process took %.2f seconds", process_end_time.Seconds())
}

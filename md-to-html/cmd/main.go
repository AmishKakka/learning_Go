package main
import (
	"fmt"
	"io/fs"
	"md-to-html/internal"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	root := "/Users/amish130437/Desktop/learning_Go/md-to-html"
	workers := 5
	var allJobs []logic.Job
	
	fmt.Println("Searching directory: ", root)
	err := filepath.WalkDir(root, func(path string, f fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, e := f.Info()
		if !f.IsDir() && e == nil && strings.HasSuffix(info.Name(), ".md") {
			jobs, err := logic.ParseFileForLinks(path)
			if err != nil {
				fmt.Println("Failed to parse file: ", path)
				return nil
			}
			// adding all the jobs found in this file
			allJobs = append(allJobs, jobs...)
		}
		return err
	})

	// early exit if error
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Now, the Jobs channel is populated we need to now check which links ar Broken and Ok
	if len(allJobs) == 0 {
		fmt.Println("No links found in markdown files.")
		return
	}
	fmt.Printf("Found %v links...\n", len(allJobs))
	
	// creating channels for communication
	jobsChannel := make(chan logic.Job, len(allJobs))
	resChannel := make(chan logic.Result, len(allJobs))
	var wg sync.WaitGroup

	// the Go runtime instantly puts this worker goroutine to sleep (blocks it), when there are no Jobs
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go logic.CheckLinks(jobsChannel, resChannel, &wg)
	}

	// filling the Jobs channel
	for _, job := range allJobs {
		jobsChannel <- job
	}
	close(jobsChannel)

	// this tells main goroutine that there is no more data in the Result channel
	go func ()  {
		wg.Wait()
		close(resChannel)
	}()

	// printing out results
	for res := range resChannel {
		if res.Broken == true {
			fmt.Printf("\033[31m[BROKEN]\033[0m Error: %s %s\n", res.Status, res.Link)
		} else {
			fmt.Printf("\033[32m[OK]\033[0m %s\n", res.Link)
		}
	}
}
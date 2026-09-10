package logic
import (
	"bufio"
	// "fmt"
	"os"
	"regexp"
)

// custom data type to store URL and the file location it exists in
type Job struct {
	link string
	location string
}

// This function gets a file path, and it parses the file to find any links present
func ParseFileForLinks(fileName string) ([]Job, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// making sure the file is closed after this function completes
	defer file.Close()

	// to find strings starting with pattern: [](https://)
	re := regexp.MustCompile(`\[.*?\]\((https?://.*?)\)`)
	// will save the links found in this Job struct 
	var jobs []Job

	// scanner object
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// applyting regex expression over a single line in the file
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// fmt.Println(match)
			jobs = append(jobs, Job{
				link: match[1],
				location: fileName,
			})
		}
	}
	// scanner.Err() is a non-EOF error, means an error encountered while reading the file
	return jobs, scanner.Err()
}
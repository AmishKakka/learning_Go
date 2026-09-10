package logic
import (
	"time"
	"net/http"
	"sync"
)

// custom data type to store URL, if the link is broken or not, and status of the link
type Result struct {
	Link string
	Status string
	Broken bool
}

// This function get a channel (list) of Jobs, checks if link is working or not and adds it to Result channel.
// This is function that will be called at the same time for multiple Job channels.
func CheckLinks(jobs <-chan Job, res chan<- Result, wg *sync.WaitGroup) {
	// we will be consuming data from jobs channel, so, <-chan
	// and putting data into res channel, so, chan<-
	defer wg.Done()
	// create a HTTP client for this particular goroutine
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	for job := range jobs {
		resp, err := client.Head(job.link)
		// directly adding to Results channel (chan objects are in heap memory)
		// That's why we don't return it
		if err != nil {
			res <- Result{
				Link: job.link,
				Status: err.Error(),
				Broken: true,
			}
			continue
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			res <- Result{
				Link: job.link,
				Status: "OK",
				Broken: false,
			}
			continue
		} else {
			res <- Result{
				Link: job.link,
				Status: resp.Status,
				Broken: true,
			}
		}
	}
}
package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a job to be done by a worker
type Job struct {
	id       int
	workload int
}

// Result represents the result of a job
type Result struct {
	job    Job
	output int
}

// worker processes jobs from the jobs channel and sends results to the results channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.id)
		time.Sleep(time.Duration(job.workload) * time.Millisecond) // Simulate work
		result := Result{job, job.workload * 2}                    // Simulate processing
		fmt.Printf("Worker %d finished job %d\n", id, job.id)
		results <- result
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		job := Job{j, j}
		jobs <- job
	}
	fmt.Println("All jobs sent to workers")
	close(jobs) // Close the jobs channel since no more jobs will be sent
	// Wait for all workers to finish
	wg.Wait()
	close(results) // Close the results channel since no more results will be sent
	fmt.Println("All jobs processed by workers")

	// Collect and print results
	for result := range results {
		fmt.Printf("Job %d processed with result: %d\n", result.job.id, result.output)
	}

	val, ok := <-results
	fmt.Printf("val: %v, ok: %v\n", val, ok)
}

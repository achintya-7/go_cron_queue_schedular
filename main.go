package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	maxConcurrent = 3              // Maximum number of concurrent processes
	requestQueue  = make([]string, 0) // Slice to hold request IDs
	queueMutex    sync.Mutex       // Mutex to synchronize access to the queue
	wg            sync.WaitGroup   // WaitGroup to synchronize goroutines
	count         = 0
)

func main() {
	r := gin.Default()

	r.GET("/", enqueueProcess)

	// Start the cron scheduler in a separate goroutine
	go cronScheduler()

	r.Run(":8080")
}

func enqueueProcess(c *gin.Context) {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(400, gin.H{"message": "Missing name parameter"})
		return
	}

	requestQueue = append(requestQueue, name)
	c.JSON(200, gin.H{"message": "Process added to the queue"})
}

func cronScheduler() {
	// Schedule the task to run every minute
	cron := time.NewTicker(time.Second)

	for range cron.C {
		processQueue()
	}

}

func processQueue() {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	// Check if the number of running instances is less than the maximum and if there are requests in the queue
	for count < maxConcurrent && len(requestQueue) > 0 {
		// Retrieve the first request from the queue
		requestID := requestQueue[0]

		// Run the request and remove it from the queue
		count++
		wg.Add(1)
		go processRequest(requestID)
		requestQueue = requestQueue[1:]
	}
}

func processRequest(requestID string) {
	length := len(requestID)
	fmt.Printf("Processing request %d..\n", length)
	// Simulate some heavy processing time
	time.Sleep(time.Second * time.Duration(length))

	fmt.Printf("Processing request %s..\n", requestID)

	count--
	wg.Done()
}

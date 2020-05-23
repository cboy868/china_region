package engine

import (
	"log"
)

// ConcurrentEngine 1
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler 1
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

//Run 1
func (c *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out, c)
	}
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		parseResult := <-out

		for _, item := range parseResult.Items {
			log.Printf("Got items: %v", item)
		}

		for _, request := range parseResult.Requests {
			c.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, c *ConcurrentEngine) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				log.Printf("fetch url:%v error:%v 重新把请求推入到requests中", request.Url, err)
				// requests = append(requests, r)
				c.Scheduler.Submit(request)
				continue
			}

			out <- result
		}
	}()
}

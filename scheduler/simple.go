package scheduler

import "github.com/cboy868/china_regions/engine"

// SimpleScheduler 1
type SimpleScheduler struct {
	workerChan chan engine.Request
}

//Submit 1
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

// ConfigureMasterWorkerChan 1
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

package types

import (
	"encoding/json"
	"net/http"
	"sync"

	"hermes/constants"
	"hermes/util"
)

// JobManager contains the job queue and provides methods to queue and run jobs
type JobManager struct {
	jobQueue *[constants.DefaultJobQueueSize]job
	mut      sync.Mutex
}

// Init must be called before using an instance of JobManager to initialize internal fields
func (manager *JobManager) Init() {
	manager.jobQueue = new([constants.DefaultJobQueueSize]job)
}

// Run infinitely loops to check for jobs on the job queue, and synchronously runs each job found
func (manager *JobManager) Run() {
	for {
		manager.mut.Lock()

		manager.mut.Unlock()
	}
}

// ProcessHTTPRequest provides a route handler func for a mux that can process a job request and queue the job
func (manager *JobManager) ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) {
	if contentTypeHeader := r.Header["Content-Type"]; contentTypeHeader == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\":\"missing Content-Type header\"}"))

		return
	}

	if contentTypeHeader := r.Header["Content-Type"]; !util.SliceContainsString(contentTypeHeader, "application/json") {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\":\"content type given by header not supported\"}"))

		return
	}

	var data job
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\":\"content provided in request body could not be decoded into json\"}"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"request successful!\"}"))
}

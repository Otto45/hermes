package jobmanager

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"hermes/constants"
	"hermes/util"
)

// JobManager contains the job queue and provides methods to queue and run jobs
type JobManager struct {
	jobQueue *jobQueue
	mut      sync.Mutex
}

// Init must be called before using an instance of JobManager to initialize internal fields
func (manager *JobManager) Init() {
	manager.jobQueue = newJobQueue()
}

// Run infinitely loops to check for jobs on the job queue, and synchronously runs each job found
func (manager *JobManager) Run() {
	// TODO: Check if Init() was called, and error if not

	for {
		var nextJob *job
		manager.mut.Lock()
		nextJob = manager.jobQueue.dequeue()
		manager.mut.Unlock()

		if nextJob == nil {
			continue
		}

		log.Printf("Running job written in %s\n", nextJob.Language)

		// TODO: Remove placeholder for simulating a job being ran
		seconds, _ := time.ParseDuration("10s")
		time.Sleep(seconds)

		log.Printf("Job complete.\n")

		// TODO: Unzip code temporarily stored on disk, and move to input folder used for mounting to containers
		// TODO: Get container containing an appropriate environment for the specified job's language
		// TODO: Run container, mounting input folder containing unzipped code and output folder
		// TODO: Zip up contents of output folder and move them to another temporary location on disk to return when CLI client queries this daemon for job completion
	}
}

// ProcessHTTPRequest provides a mux route handler to process and queue job requests
func (manager *JobManager) ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if Init() was called, and error if not

	if contentTypeHeader := r.Header[constants.ContentTypeHeaderName]; contentTypeHeader == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set(constants.ContentTypeHeaderName, constants.ContentTypeApplicationJSON)
		w.Write([]byte(util.BuildHTTPResponseBodyForError("missing Content-Type header")))

		return
	}

	if contentTypeHeader := r.Header["Content-Type"]; !util.SliceContainsString(contentTypeHeader, "application/json") {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set(constants.ContentTypeHeaderName, constants.ContentTypeApplicationJSON)
		w.Write([]byte(util.BuildHTTPResponseBodyForError("content type given by header not supported")))

		return
	}

	var data job
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set(constants.ContentTypeHeaderName, constants.ContentTypeApplicationJSON)
		w.Write([]byte(util.BuildHTTPResponseBodyForError("content provided in request body could not be decoded into json")))

		return
	}

	// TODO: Save zipped up source code from request body to temporary location on disk and update job object

	manager.mut.Lock()
	manager.jobQueue.queue(&data)
	manager.mut.Unlock()

	log.Printf("Job queued successfully!\n")

	w.WriteHeader(http.StatusOK)
	w.Header().Set(constants.ContentTypeHeaderName, constants.ContentTypeApplicationJSON)
	w.Write([]byte(util.BuildHTTPResponseBodyForSuccess("success")))
}

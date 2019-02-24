package jobmanager

import (
	"hermes/constants"
)

type jobQueue struct {
	jobs           []*job
	head           int
	tail           int
	queueSize      int
	queuedJobCount int
}

func newJobQueue() *jobQueue {
	return &jobQueue{
		make([]*job, constants.DefaultJobQueueSize),
		1,
		1,
		constants.DefaultJobQueueSize,
		0}
}

func (jq *jobQueue) queue(job *job) {
	jq.jobs[jq.tail-1] = job
	jq.queuedJobCount++
	jq.tail = jq.tail%len(jq.jobs) + 1

	if jq.head == jq.tail {
		// TODO: Grow queue
	}
}

func (jq *jobQueue) dequeue() *job {
	if jq.queuedJobCount == 0 {
		return nil
	}

	nextJob := jq.jobs[jq.head-1]
	jq.queuedJobCount--
	jq.head = jq.head%len(jq.jobs) + 1

	if ratio := float32(jq.queuedJobCount) / float32(jq.queueSize); ratio <= 0.25 {
		// TODO: Shrink queue
	}

	return nextJob
}

func (jq *jobQueue) grow() {
	newQueueSize := jq.queueSize << 1
	newJobQueue := make([]*job, newQueueSize)
	copy(newJobQueue, jq.jobs)
	jq.tail = jq.queueSize + 1

	// Need to preserve queue order after resizing slice and adjusting tail pointer
	if jq.head > 1 {
		for i := 0; i < jq.head; i++ {
			job := newJobQueue[i]

		}
	}
}

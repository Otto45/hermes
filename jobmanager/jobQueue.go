package jobmanager

import (
	"hermes/constants"
	"log"
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
		jq.grow()
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
		jq.shrink()
	}

	return nextJob
}

func (jq *jobQueue) grow() {
	newQueueSize := jq.queueSize << 1

	log.Printf("Growing queue from size %d to size %d.\n", jq.queueSize, newQueueSize)

	newJobQueue := make([]*job, newQueueSize)

	// Preserve queue order when copying elements to resized queue
	if jq.head == 1 {
		copy(newJobQueue, jq.jobs)
	} else {
		head := jq.head - 1
		numberItemsCopied := copy(newJobQueue, jq.jobs[head:])
		copy(newJobQueue[numberItemsCopied:], jq.jobs[0:head])
	}

	jq.head = 1
	jq.tail = jq.queueSize + 1
	jq.queueSize = newQueueSize
	jq.jobs = newJobQueue

	log.Printf("Queue grown successfully.\n")
	log.Printf("Number of remaining jobs: %d\n", jq.queuedJobCount)
}

func (jq *jobQueue) shrink() {
	newQueueSize := jq.queueSize >> 1

	log.Printf("Shrinking queue from size %d to size %d.\n", jq.queueSize, newQueueSize)

	if newQueueSize <= constants.DefaultJobQueueSize {
		return
	}

	newJobQueue := make([]*job, newQueueSize)

	// Preserve queue order when copying elements to resized queue
	if jq.head < jq.tail {
		copy(newJobQueue, jq.jobs[jq.head-1:jq.tail])
	} else {
		numberItemsCopied := copy(newJobQueue, jq.jobs[jq.head-1:])
		copy(newJobQueue[numberItemsCopied:], jq.jobs[0:jq.tail])
	}

	jq.head = 1
	jq.tail = jq.queuedJobCount + 1
	jq.queueSize = newQueueSize
	jq.jobs = newJobQueue

	log.Printf("Queue shrunken successfully.\n")
	log.Printf("Number of remaining jobs: %d\n", jq.queuedJobCount)
}

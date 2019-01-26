package jobmanager

// Job contains the deserialized request body from the CLI client
type job struct {
	Language string `json:"language"`
}

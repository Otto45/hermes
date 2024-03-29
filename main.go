package main

import (
	"log"
	"net/http"

	"hermes/jobmanager"

	"github.com/gorilla/mux"
)

func main() {
	jobManager := jobmanager.NewManager()

	// Launches infinite goroutine to concurrently check for and run jobs placed on the queue
	go jobManager.Run()

	// Main thread will infinitely listen for POSTed jobs to be placed on the queue
	router := mux.NewRouter()
	router.HandleFunc("/", jobManager.ProcessHTTPRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// func redisExample() error {
// 	// create a new client connected to the default socket path for containerd
// 	client, err := containerd.New("/run/containerd/containerd.sock")
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	// create a new context with an "example" namespace
// 	ctx := namespaces.WithNamespace(context.Background(), "example")

// 	// pull the redis image from DockerHub
// 	image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
// 	if err != nil {
// 		return err
// 	}

// 	// create a container
// 	container, err := client.NewContainer(
// 		ctx,
// 		"redis-server",
// 		containerd.WithImage(image),
// 		containerd.WithNewSnapshot("redis-server-snapshot", image),
// 		containerd.WithNewSpec(oci.WithImageConfig(image)),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer container.Delete(ctx, containerd.WithSnapshotCleanup)

// 	// create a task from the container
// 	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
// 	if err != nil {
// 		return err
// 	}
// 	defer task.Delete(ctx)

// 	// make sure we wait before calling start
// 	exitStatusC, err := task.Wait(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// call start on the task to execute the redis server
// 	if err := task.Start(ctx); err != nil {
// 		return err
// 	}

// 	// sleep for a lil bit to see the logs
// 	time.Sleep(10 * time.Second)

// 	// kill the process and get the exit status
// 	if err := task.Kill(ctx, syscall.SIGTERM); err != nil {
// 		return err
// 	}

// 	// wait for the process to fully exit and print out the exit status

// 	status := <-exitStatusC
// 	code, _, err := status.Result()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("redis-server exited with status: %d\n", code)

// 	return nil
// }

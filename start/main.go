package main

import (
	"context"
	"fmt"
	"log"

	app "github.com/luansapelli/hello-world-temporal"
	"go.temporal.io/sdk/client"
)

func main () {
    // Create the client object just once per process
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("unable to create Temporal client", err)
    }
    defer c.Close()

	options := client.StartWorkflowOptions{
		ID: "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}
	
	// Start the Workflow
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the Workflow result
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	fmt.Printf("\nWorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
	fmt.Println(greeting)
	
}
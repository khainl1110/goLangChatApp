package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func publish(projectID, topicID, msg string) error {
	// projectID = "api-project-70002766628"
	// topicID = "test-topic"
	// msg = "Hello World"
	fmt.Print("test")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Print("error!")
		return err
	}
	defer client.Close()

	// messages := []struct {
	// 	message     string
	// 	orderingKey string
	// }{
	// 	{
	// 		message:     "message1",
	// 		orderingKey: "key1",
	// 	},
	// 	{
	// 		message:     "message2",
	// 		orderingKey: "key2",
	// 	},
	// }
	fmt.Print("test1")
	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		// can only send byte?
		Data: []byte(msg),
		// Data: []byte (messages),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Errorf("pubsub: result.Get: %w", err)
		fmt.Println(err)
	}

	fmt.Print("test11 " + id)
	fmt.Println("Published a message; msg ID: %v\n", id)
	fmt.Print("published a message")
	return nil
}

func pullMsgsCustomAttributes(projectID, subID string) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	print("wassup")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Receive blocks until the context is cancelled or an error occurs.
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		print("Got message ", string(msg.Data))
		print("Attributes: ")
		for key, value := range msg.Attributes {
			print("key", key, "value", value)
		}
		// fmt.Fprintf(w, "Got message :%q\n", string(msg.Data))
		// fmt.Fprintln(w, "Attributes:")
		// for key, value := range msg.Attributes {
		// 	fmt.Fprintf(w, "%s = %s\n", key, value)
		// }
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	return nil
}

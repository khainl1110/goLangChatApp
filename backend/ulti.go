package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func publishMsg(projectID, topicID, msg string) error {
	// fmt.Print("test")
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
	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		// can only send byte?
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Errorf("pubsub: result.Get: %w", err)
	}

	fmt.Println("Published a message; msg ID: ", id)
	return nil
}

func pullMsg(projectID, subID string) string {
	// print("wassup")
	result := "empty"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return err.Error()
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
		// print("Attributes: ")
		// for key, value := range msg.Attributes {
		// 	print("key", key, "value", value)
		// }
		result = string(msg.Data)
		msg.Ack()

	})
	if err != nil {
		return err.Error()
	}

	return result
}

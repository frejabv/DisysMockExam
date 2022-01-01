package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"DisysMockExam/Mock/protobuf"

	"google.golang.org/grpc"
)

func main() {
	LOG_FILE := "../log.txt"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	conn1, err1 := grpc.Dial(":8081", grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err1)
	}

	conn2, err2 := grpc.Dial(":8082", grpc.WithInsecure(), grpc.WithBlock())
	if err2 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err2)
	}

	defer conn.Close()
	defer conn1.Close()
	defer conn2.Close()

	client := protobuf.NewMockClient(conn)
	client1 := protobuf.NewMockClient(conn1)
	client2 := protobuf.NewMockClient(conn2)

	fmt.Println("Press enter to increment, results will be written to log file")
	go TakeInput(client, client1, client2)
	time.Sleep(1000 * time.Second)

}

func TakeInput(client protobuf.MockClient, client1 protobuf.MockClient, client2 protobuf.MockClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		if input != "\n" {
			log.Fatal("User typed something else")
		}

		message, error := client.Increment(context.Background(), &protobuf.IncrementRequest{})
		message1, error1 := client1.Increment(context.Background(), &protobuf.IncrementRequest{})
		message2, error2 := client2.Increment(context.Background(), &protobuf.IncrementRequest{})

		var values []int32
		if error == nil {
			values = append(values, message.NewValue)
		}
		if error1 == nil {
			values = append(values, message1.NewValue)
		}
		if error2 == nil {
			values = append(values, message2.NewValue)
		}

		var highestValue int32
		for i := 0; i < len(values); i++ {
			if values[i] > highestValue {
				highestValue = values[i]
			}
		}

		//syncValues
		if error == nil && message.NewValue != highestValue {
			client.SetValue(context.Background(), &protobuf.SetValueRequest{Value: highestValue})
		}
		if error1 == nil && message1.NewValue != highestValue {
			client1.SetValue(context.Background(), &protobuf.SetValueRequest{Value: highestValue})
		}
		if error2 == nil && message2.NewValue != highestValue {
			client2.SetValue(context.Background(), &protobuf.SetValueRequest{Value: highestValue})
		}

		log.Println("The new value is:", highestValue)
	}
}

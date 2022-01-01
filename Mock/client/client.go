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
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	/*var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}*/

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	//defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	//c := myPackage.NewGetXXXClient(conn)

	//SendRequest(c)

	// log to custom file
	LOG_FILE := "../log.txt"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
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
	//make below into go routine and use: time.Sleep(1000 * time.Second)
	fmt.Println("Press enter to increment, results will be written to log file")
	go TakeInput(client, client1, client2)
	time.Sleep(1000 * time.Second)

}

func TakeInput(client protobuf.MockClient, client1 protobuf.MockClient, client2 protobuf.MockClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		if input == "omg lol" {
			log.Fatal("Something went wrong")
		}

		message, error := client.Increment(context.Background(), &protobuf.IncrementRequest{})
		message1, error1 := client1.Increment(context.Background(), &protobuf.IncrementRequest{})
		message2, error2 := client2.Increment(context.Background(), &protobuf.IncrementRequest{})
		if error != nil && error1 != nil && error2 != nil { //error can not establish connection
			log.Fatalf("failed while incrementing: %v", error)
		}
		values := []int32{message.NewValue, message1.NewValue, message2.NewValue}

		var highestValue int32
		for i := 0; i < len(values); i++ {
			if values[i] > highestValue {
				highestValue = values[i]
			}
		}

		log.Println("The new value is:", highestValue)
	}
}

/*func Dial(port int, name string) (protobuf.ReplicationClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(":"+strconv.Itoa(port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	frontend := protobuf.NewReplicationClient(conn)
	message, userAlreadyExistsError := frontend.NewNode(context.Background(), &protobuf.NewNodeRequest{Name: name, Type: *protobuf.NewNodeRequest_FrontEnd.Enum()})
	if userAlreadyExistsError != nil {
		//Error handling
		if message == nil {
			log.Println("Username is already in use")
		}
	} else {
		//Start to do stuff here
		log.Println("Dial to " + strconv.Itoa(port) + " was succesful")
		return frontend, conn
	}
	return nil, nil
}*/

/*func SendRequest(c myPackage.XXXClient) {
    // Between the curly brackets are nothing, because the .proto file expects no input.
	message := myPackage.somethingsomethingRequest{}

    response, err := c.somethingsomething(context.Background(), &message)
    if err != nil {
        log.Fatalf("Error when calling XXX: %s", err)
    }

    log.Printf("Response from the Server: %s \n", response.Reply)
}*/

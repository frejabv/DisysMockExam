package main

import (
	"context"
	"log"

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
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err)
	}

	conn1, err1:= grpc.Dial(":8081", grpc.WithInsecure(), grpc.WithBlock())
	if err1 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err1)
	}

	conn2, err2:= grpc.Dial(":8081", grpc.WithInsecure(), grpc.WithBlock())
	if err2 != nil { //error can not establish connection
		log.Fatalf("did not connect: %v", err2)
	}

	defer conn1.Close()
	defer conn2.Close()
	defer conn3.Close()

	client := protobuf.NewMockClient(conn)
	//make below into go routine and use: time.Sleep(1000 * time.Second)
	message, error := client.Increment(context.Background(), &protobuf.IncrementRequest{})
	if error != nil { //error can not establish connection
		log.Fatalf("failed while incrementing: %v", error)
	}
	log.Println("The new value is:", message.NewValue)
}

/*func SendRequest(c myPackage.XXXClient) {
    // Between the curly brackets are nothing, because the .proto file expects no input.
	message := myPackage.somethingsomethingRequest{}

    response, err := c.somethingsomething(context.Background(), &message)
    if err != nil {
        log.Fatalf("Error when calling XXX: %s", err)
    }

    fmt.Printf("Response from the Server: %s \n", response.Reply)
}*/

# protobuf example
A simple example of using protobuf in golang

# Summary
If you are not using gRPC for server to server communications in your golang app, you are missing out on one of the best optimizations that you can make for the performance of the code and the development team.

1. It makes understanding the data model easier by strictly enforcing the proto 3 syntax. Json is easy to get started with and make a prototype, but protobuf will be more strict about syntax and standardize your communication between services.

2. Updating the schema for a gRPC request is simple, update the protobuf file and run code generation. Then it becomes like using any other package written in go to call your external service. No more boilerplate client code.

3. Protobuf marshaling is faster, uses less memory and produces a smaller payload than Json or xml Marshalling. This directly impacts infrastructure costs and will save money.

The time and cost savings you will see from using gRPC in your golang app is worth the time investment it will take to learn and implement it.

# creating data structures using protobuf
write the protobuf file
```proto
syntax = "proto3";
package publish;
option go_package  = "/publish";

message Event {
  bytes content = 1;
  repeated string tags = 2;
}

message EventList {
  repeated Event events = 1;
}
```

This creates two structs, one is an event struct that contains `content` that is a byte array and `tags` that is a string array of strings
There is also a struct for an array of events.

Now run the protoc command to generate the go code we will use in `main.go`
```
protoc --proto_path=. --go_out=. publish.proto
```
This creates the directory `publish` that contains the package `publish` with the go code that we can import to our `main.go` file.

Now we can import it using `github.com/kfelter/protobuf-example/publish`

There are other advantages to using protobuf that can be found in the protobuf documentation.

References:

install protoc           https://grpc.io/docs/protoc-installation/

proto3 language guide    https://developers.google.com/protocol-buffers/docs/proto3#simple

protobuf golang tutorial https://developers.google.com/protocol-buffers/docs/gotutorial


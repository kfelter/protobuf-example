# protobuf example
A simple example of using protobuf in golang

# Summary
Protobuf is a method of structuring data so that it can be stored as a byte array.
JSON could also fit this description but there are many differences between json and protobuf.

The main difference demonstrated by this example is the size of the byte array when converting a data struct to a byte array.
Protobuf is a much more efficient conversion method than json.

In `main.go` we build an EventList object and marshal it into json and protobuf. 
```go
start := time.Now()
buf, _ := proto.Marshal(&e)
os.WriteFile("events.protobuf", buf, os.ModePerm)
fmt.Println("marshal events.protobuf", time.Since(start))

start = time.Now()
buf, _ = json.Marshal(&e)
os.WriteFile("events.json", buf, os.ModePerm)
fmt.Println("marshal events.json", time.Since(start))
```
```txt
$ go run main.go
marshal events.protobuf 397.726µs
marshal events.json 269.558µs
```

This indicates it takes slightly longer to marshal data into protobuf, but the file size of events.protobuf is much smaller.
```
$ ls -al events*
-rwxr-xr-x  1 kylefelter  staff  144 Apr 30 22:35 events.json
-rwxr-xr-x  1 kylefelter  staff   71 Apr 30 22:35 events.protobuf
```

Now checking the unmarshalling it takes a bit longer to unmarshal protobuf also.
```go
	events := publish.EventList{}
buf, _ = os.ReadFile("events.protobuf")
start = time.Now()
proto.Unmarshal(buf, &events)
fmt.Println("unmarshal protobuf", events.String(), time.Since(start))

events = publish.EventList{}
buf, _ = os.ReadFile("events.json")
start = time.Now()
json.Unmarshal(buf, &events)
fmt.Println("unmarshal json", events.String(), time.Since(start))
```
```txt
$ go run main.go
marshal events.protobuf 397.726µs
marshal events.json 269.558µs
unmarshal protobuf events:{content:"some event content!"  tags:"tag1"  tags:"tag2"}  events:{content:"some event content2!"  tags:"tag3"  tags:"tag4"} 98.805µs
unmarshal json events:{content:"some event content!"  tags:"tag1"  tags:"tag2"}  events:{content:"some event content2!"  tags:"tag3"  tags:"tag4"} 40.391µs
```

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

Referances:

install protoc           https://grpc.io/docs/protoc-installation/

proto3 language guide    https://developers.google.com/protocol-buffers/docs/proto3#simple

protobuf golang tutorial https://developers.google.com/protocol-buffers/docs/gotutorial


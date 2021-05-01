package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kfelter/protobuf-example/publish"
	"google.golang.org/protobuf/proto"
)

func main() {
	e := publish.EventList{
		Events: []*publish.Event{
			{
				Content: []byte(`some event content!`),
				Tags:    []string{"tag1", "tag2"},
			},
			{
				Content: []byte(`some event content2!`),
				Tags:    []string{"tag3", "tag4"},
			},
		},
	}

	start := time.Now()
	buf, _ := proto.Marshal(&e)
	os.WriteFile("events.protobuf", buf, os.ModePerm)
	fmt.Println("marshal events.protobuf", time.Since(start))

	start = time.Now()
	buf, _ = json.Marshal(&e)
	os.WriteFile("events.json", buf, os.ModePerm)
	fmt.Println("marshal events.json", time.Since(start))

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
}

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

	// 	: []*publish.Event{
	// 		{
	// 			Content: []byte(`some event content!`),
	// 			Tags:    []string{"tag1", "tag2"},
	// 		},
	// 		{
	// 			Content: []byte(`some event content2!`),
	// 			Tags:    []string{"tag3", "tag4"},
	// 		},
	// 	},
	// }

	start := time.Now()
	buf, err := proto.Marshal(&e)
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile("events.protobuf", buf, os.ModePerm); err != nil {
		panic(err)
	}
	fmt.Println("marshal events.protobuf", time.Since(start))
	start = time.Now()
	buf, err = json.Marshal(&e)
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile("events.json", buf, os.ModePerm); err != nil {
		panic(err)
	}
	fmt.Println("marshal events.json", time.Since(start))

	events := publish.EventList{}
	buf, err = os.ReadFile("events.protobuf")
	if err != nil {
		panic(err)
	}
	start = time.Now()
	err = proto.Unmarshal(buf, &events)
	if err != nil {
		panic(err)
	}
	fmt.Println("unmarshal protobuf", events.String(), time.Since(start))

	events = publish.EventList{}
	buf, err = os.ReadFile("events.json")
	if err != nil {
		panic(err)
	}
	start = time.Now()
	err = json.Unmarshal(buf, &events)
	if err != nil {
		panic(err)
	}
	fmt.Println("unmarshal json", events.String(), time.Since(start))
}

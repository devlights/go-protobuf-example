package main

import (
	"fmt"
	"log"
	"protobuftest/internal/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	// ---------------------------------------------
	// Protocol Buffers を使ったシリアライズとデシリアライズのサンプルです.
	//
	// REFERENCES:
	//   - https://developers.google.com/protocol-buffers/docs/gotutorial
	//   - https://dev.to/heerthees/protobuf-with-go-4hb
	// ---------------------------------------------

	// ---------------------------------------------
	// シリアライズ
	// ---------------------------------------------
	p := &pb.Person{
		Name: "hoge",
		Age:  99,
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	// - https://devlights.hatenablog.com/entry/2020/08/18/014703
	fmt.Printf("[marshal  ] [%10T] %[1]v\n", data)

	// ---------------------------------------------
	// デシリアライズ
	// ---------------------------------------------
	p2 := new(pb.Person)

	err = proto.Unmarshal(data, p2)
	if err != nil {
		log.Fatal(err)
	}

	// - https://devlights.hatenablog.com/entry/2020/08/18/014703
	fmt.Printf("[unmarshal] [%10T] %[1]v\n", p2)
}

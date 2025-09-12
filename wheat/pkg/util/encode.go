package util

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

func PbEncode() {
	b, err := proto.Marshal(&durationpb.Duration{
		Nanos: 1024,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("125ns encoded into %d bytes of Protobuf wire format:\n% x\n", len(b), b)

	var dur durationpb.Duration
	if err := proto.Unmarshal(b, &dur); err != nil {
		panic(err)
	}
	fmt.Printf("Protobuf wire format decoded to duration %v\n", dur.AsDuration())
}

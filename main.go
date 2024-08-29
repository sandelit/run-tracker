package main

import (
	"fmt"
	"os"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
)

func main() {
	f, err := os.Open("activity.fit")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := decoder.New(f)

	fit, err := dec.Decode()
	if err != nil {
		panic(err)
	}

	fmt.Printf("FileHeader DataSize: %d\n", fit.FileHeader.DataSize)
	fmt.Printf("Messages count: %d\n", len(fit.Messages))
	// FileId is always the first message; 4 = activity
	fmt.Printf("File Type: %v\n",
		fit.Messages[0].FieldValueByNum(fieldnum.FileIdType).Any())
}

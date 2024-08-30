package main

import (
	"fmt"
	"os"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/filedef"
)

func main() {
	f, err := os.Open("rasmus-test.fit")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := decoder.New(f)

	if err != nil {
		panic(err)
	}
	/*
		fmt.Printf("FileHeader DataSize: %d\n", fit.FileHeader.DataSize)
		fmt.Printf("Messages count: %d\n", len(fit.Messages))
		// FileId is always the first message; 4 = activity
		fmt.Printf("File Type: %v\n",
			fit.Messages[0].FieldValueByNum(fieldnum.FileIdType).Any())*/

	for dec.Next() {
		fit, err := dec.Decode()
		if err != nil {
			panic(err)
		}

		activity := filedef.NewActivity(fit.Messages...)

		fmt.Printf("File Type: %s\n", activity.FileId.Type)
		fmt.Printf("Sessions count: %d\n", len(activity.Sessions))
		fmt.Printf("Laps count: %d\n", len(activity.Laps))
		fmt.Printf("Records count: %d\n", len(activity.Records))

		i := len(activity.Records) - 1

		fmt.Printf("\nSample value of record[%d]:\n", i)
		fmt.Printf("  Distance: %g m\n", activity.Records[i].DistanceScaled())
		fmt.Printf("  Lat: %d semicircles\n", activity.Records[i].PositionLat)
		fmt.Printf("  Long: %d semicircles\n", activity.Records[i].PositionLong)
		fmt.Printf("  Speed: %g m/s\n", activity.Records[i].SpeedScaled())
		fmt.Printf("  Average HeartRate: %d bpm\n", activity.Records[i].HeartRate)
		fmt.Printf("  Average Cadence: %d rpm\n", activity.Records[i].Cadence*2)
	}

}

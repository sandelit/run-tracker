package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/profile/filedef"
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	f, err := os.Open("data/fit/rasmus-test.fit")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	printFitFile(f)
}

// TODO: Move somewhere later
func printFitFile(f *os.File) {
	dec := decoder.New(f)
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

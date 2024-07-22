package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
	for {
		f, err := os.Open("sound1.mp3")
		if err != nil {
			log.Fatal(err)
		}
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

		speaker.Play(streamer)

		time.Sleep(2 * time.Minute)
		streamer.Close()
		f.Close()
	}
}

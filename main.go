package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("No name provided!")
		return
	}
	name := args[0]
	x := GetXML(name)
	episodes := ParseXML(x)

	var errs []string

	for i, episode := range episodes {
		if i == 5 {
			break
		}
		if episode.HD.FileLocation != "" {
			fmt.Printf("Downloading HD video %d of %d\n", i+1, len(episodes))
			if err := Download(episode.HD.FileLocation, name); err == nil {
				continue
			}
		}
		if episode.SD.FileLocation != "" {
			fmt.Printf("Downloading SD video %d of %d\n", i+1, len(episodes))
			if err := Download(episode.SD.FileLocation, name); err == nil {
				continue
			}
		}
		errs = append(errs, fmt.Sprintf("Episode #%d: %s", i+1, episode.Title))
	}

	if len(errs) > 0 {
		fmt.Printf("Unable to download the following...\n\n")
		for _, s := range errs {
			fmt.Println(s)
		}
	}
}

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
	fmt.Println("Getting playlist XML...")
	x := GetXML(name)
	fmt.Println("Parsing playlist XML")
	episodes := ParseXML(x)

	var errs []string
	fmt.Println("Starting episodes download...")
	for i, episode := range episodes {
		if episode.HD.FileLocation != "" {
			fmt.Printf("Downloading HD video %d of %d...\n", i+1, len(episodes))
			if err := Download(episode.HD.FileLocation, name); err == nil {
				continue
			}
		}
		if episode.SD.FileLocation != "" {
			fmt.Printf("Downloading SD video %d of %d...\n", i+1, len(episodes))
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

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var WGET_ERR *os.File
var APP_LOG *os.File
var FAILED *os.File

var start int
var end int

func main() {

	var err error

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No name provided!")
		return
	}

	name := args[0]

	if len(args) > 1 {
		start, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Start must Be a number!\n")
		}
		start -= 1
	}

	if len(args) > 2 {
		end, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatalf("End must Be a number!\n")
		}
	}

	if end != 0 && start > end {
		log.Fatalf("Start must be less than end!\n")
	}

	WGET_ERR, err = os.OpenFile(name+"_wget.err", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening %s_wget.err\n%v\n\n", name, err)
	}
	defer WGET_ERR.Close()

	APP_LOG, err = os.OpenFile(name+"_app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening %s_app.err\n%v\n\n", name, err)
	}
	defer APP_LOG.Close()

	FAILED, err = os.OpenFile(name+"_failures.err", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening %s_failures.err\n%v\n\n", name, err)
	}
	defer FAILED.Close()

	fmt.Println("Getting playlist XML...")
	x := GetXML(name)
	fmt.Println("Parsing playlist XML")
	episodes := ParseXML(x)

	if end == 0 {
		end = len(episodes)
	}

	var errs []string
	fmt.Println("Starting episodes download...")
	episodes = episodes[start:end]
	fmt.Println(len(episodes))
	for i, episode := range episodes {
		if episode.HD.FileLocation != "" {
			fmt.Printf("Downloading HD video %d of %d...\n", i+1, len(episodes))
			err := Download(episode.HD.FileLocation, name)
			if err == nil {
				continue
			}
			Logger(APP_LOG, err.Error())
		}
		Logger(APP_LOG, "Episode %s has no HD link", episode.Title)
		if episode.SD.FileLocation != "" {
			fmt.Printf("Downloading SD video %d of %d...\n", i+1, len(episodes))
			err := Download(episode.SD.FileLocation, name)
			if err == nil {
				continue
			}
			Logger(APP_LOG, err.Error())
		}
		Logger(APP_LOG, err.Error())
		Logger(FAILED, "Episode #%d: %s", i+1, episode.Title)
		errs = append(errs, fmt.Sprintf("Episode #%d: %s", i+1, episode.Title))
	}

	if len(errs) > 0 {
		fmt.Printf("Unable to download the following...\n\n")
		for _, s := range errs {
			fmt.Println(s)
		}
	}
}

func Logger(f *os.File, s string, v ...interface{}) {
	log.SetOutput(f)
	log.Printf("\n"+s+"\n\n", v...)
}

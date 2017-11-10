package main

import (
	"bytes"
	"encoding/xml"
	"log"
)

func ParseXML(x []byte) []Episode {

	dec := xml.NewDecoder(bytes.NewReader(x))
	dec.Strict = false
	var show Show
	if err := dec.Decode(&show); err != nil {
		log.Fatalf("xmlParser.go >> ParseXML() >> dec.Decode() >> \n%v\n\n", err)
	}

	var episodes []Episode

	for _, item := range show.Items {
		var episode Episode
		episode.Title = item.Title
		for _, source := range item.Sources {
			if source.Quality == "HD" {
				episode.HD = source
			} else if source.Quality == "SD" {
				episode.SD = source
			}
		}
		episodes = append(episodes, episode)
	}

	return episodes
}

type Show struct {
	Items []Item `xml:"channel>item"`
}

type Item struct {
	Title   string   `xml:"title"`
	Sources []Source `xml:"source"`
}

type Source struct {
	FileLocation string `xml:"file,attr"`
	FileType     string `xml:"type,attr"`
	Quality      string `xml:"label,attr"`
}

type Episode struct {
	Title string
	HD    Source
	SD    Source
}

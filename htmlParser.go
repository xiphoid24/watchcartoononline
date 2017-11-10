package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func GetXML(name string) []byte {
	resp, err := http.Get("https://www.watchcartoononline.io/playlist-cat/" + name)
	if err != nil {
		log.Fatalf("htmlParser.go >> GetXML() >> http.Get() >> \n%v\n\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("htmlParser.go >> GetXML() >> ioutil.ReadAll() >> \n%v\n\n", err)
	}
	i := bytes.Index(body, []byte("jw.setup("))
	if i == -1 {
		log.Fatalf("htmlParser.go >> GetXML() >> bytes.Index >> \nUnable to find jw.setup(\n\n")
	}

	i += 9

	j := bytes.Index(body[i:], []byte(");"))
	if j == -1 {
		log.Fatalf("htmlParser.go >> GetXML() >> bytes.Index >> \nUnable to find END of jw.setup(\n\n")
	}
	j += i

	body = body[i:j]

	i = bytes.Index(body, []byte(`playlist: "`))
	if i == -1 {
		log.Fatalf("htmlParser.go >> GetXML() >> bytes.Index >> \nUnable to find playlist\n\n")
	}
	i += 11

	j = bytes.Index(body[i:], []byte(`",`))
	if j == -1 {
		log.Fatalf("htmlParser.go >> GetXML() >> bytes.Index >> \nUnable to find END of playlist\n\n")
	}
	j += i

	resp, err = http.Get(string(body[i:j]))
	if err != nil {
		log.Fatalf("htmlParser.go >> GetXML() >> http.Get() >> \n%v\n\n", err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("htmlParser.go >> GetXML() >> ioutil.ReadAll() >> \n%v\n\n", err)
	}

	return body
}

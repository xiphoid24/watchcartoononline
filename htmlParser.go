package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GetXML(name string) []byte {
	resp, err := http.Get("https://www.watchcartoononline.io/playlist-cat/" + name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	i := bytes.Index(body, []byte("jw.setup("))
	if i == -1 {
		panic("Unable to find jw.setup(")
	}

	i += 9

	j := bytes.Index(body[i:], []byte(");"))
	if j == -1 {
		panic("Unable to find end of jw.setup(")
	}
	j += i

	body = body[i:j]

	i = bytes.Index(body, []byte(`playlist: "`))
	if i == -1 {
		panic("Unable to find playlist")
	}
	i += 11

	j = bytes.Index(body[i:], []byte(`",`))
	if j == -1 {
		panic("Unable to find playlist")
	}
	j += i

	resp, err = http.Get(string(body[i:j]))
	if err != nil {
		panic(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

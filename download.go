package main

import (
	"os"
	"os/exec"
)

func Download(url, folder string) error {
	path := folder + "/"
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}
	cmd := exec.Command("wget", url)
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}

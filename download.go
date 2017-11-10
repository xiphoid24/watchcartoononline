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

	if b, err := cmd.CombinedOutput(); err != nil {
		Logger(WGET_ERR, "%s", b)
		return err
	}

	return nil

}

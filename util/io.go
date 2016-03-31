package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// const MaxOutputFileSize = 30 * 1024 * 1024
const AVConvPath = "/usr/bin/avconv"
const M4aExt = ".m4a"

func ReadFile(f *os.File) ([]byte, error) {
	return ioutil.ReadFile(f.Name())
}

func WriteFile(dst string, d []byte) error {
	fmt.Printf("Writing %d bytes to file\n", len(d))
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
	}
	return err
}

func CleanTempFiles(fs ...string) error {
	for _, f := range fs {
		err := os.Remove(f)
		if err != nil {
			return err
		}
	}
	return nil
}

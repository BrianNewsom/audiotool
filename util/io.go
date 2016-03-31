package util

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(f *os.File) ([]byte, error) {
	return ioutil.ReadFile(f.Name())
}

func WriteFile(dst string, d []byte) error {
	log.Printf("Writing %d bytes to file %s\n", len(d), dst)
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

package audiotool

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/briannewsom/audiotool/util"
)

func ConcatenateToBytes(ext, f1, f2 string) ([]byte, error) {
	fName, err := Concatenate(ext, f1, f2)

	if err != nil {
		return nil, err
	}

	f, err := os.Open(fName)
	if err != nil {
		return nil, err
	}

	return util.ReadFile(f)
}

func Concatenate(ext string, fs ...string) (string, error) {
	// TODO: Make this work both directions at the same time
	if len(fs) < 2 {
		return "", errors.New("Please provide more than one file to concatenate!")
	}

	outputFile, err := ioutil.TempFile("/tmp", "concat-")

	if err != nil {
		return "", err
	}

	outputFileName := outputFile.Name() + ext
	os.Rename(outputFile.Name(), outputFileName)

	log.Printf("Concatenating files to %s", outputFileName)

	if err != nil {
		return "", err
	}

	cmd := exec.Command(getAVConvPath(), "-y", "-i",
		buildConcatCmd(fs...), "-c", "copy", outputFileName)

	err = cmd.Run()

	if err != nil {
		return "", err
	}

	return outputFileName, nil
}

func buildConcatCmd(fs ...string) string {
	core := ""
	for i, f := range fs {
		if i == len(fs)-1 {
			core += f
		} else {
			core += f + "|"
		}
	}
	return "concat:" + core
}

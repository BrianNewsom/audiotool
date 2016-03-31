package audiotool

import (
	"errors"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/briannewsom/audiotool/util"
)

func Compress(content []byte, uuid string, bitrate string) ([]byte, error) {
	var inputFileName string
	var outputFileName string

	if uuid != "" {
		inputFileName = "/tmp/raw-" + uuid + ".m4a"
		outputFileName = "/tmp/" + uuid + ".m4a"
	} else {
		return nil, errors.New("No uuid given")
	}

	defer util.CleanTempFiles(inputFileName, outputFileName)

	if content != nil {
		util.WriteFile(inputFileName, content)
	} else {
		return nil, errors.New("No content supplied to compress")
	}

	err := ChangeBitrate(inputFileName, outputFileName, bitrate)

	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadFile(outputFileName)

	if err != nil {
		return nil, err
	}

	return d, nil

}

func ChangeBitrate(inputFileName string, outputFileName string, bitrate string) error {
	// Returns outputFileName
	cmd := exec.Command(getAVConvPath(), "-y", "-i", inputFileName,
		"-strict", "experimental", "-b", bitrate, outputFileName)

	err := cmd.Start()
	if err != nil {
		return err
	}

	log.Printf("Changing Bitrate of audio file %s to %s", inputFileName, bitrate)

	err = cmd.Wait()

	if err != nil {
		return err
	}

	log.Print("Successfully processed audio file.\n")

	return nil

}

package audiotool

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/briannewsom/audiotool/util"
)

func ConvertWavToM4aBytes(f string) ([]byte, error) {
	fName, err := ConvertWavToM4a(f)

	if err != nil {
		return nil, err
	}

	fi, err := os.Open(fName)
	if err != nil {
		return nil, err
	}

	util.CleanTempFiles(f)
	return util.ReadFile(fi)
}

func ConvertWavToM4a(f string) (string, error) {
	outputFile, _ := ioutil.TempFile("/tmp", "convert-")

	outputFileName := outputFile.Name() + M4aExt

	os.Rename(outputFile.Name(), outputFileName)

	log.Printf("Converting wav file %s to m4a file %s", f, outputFileName)

	cmd := exec.Command(getAVConvPath(), "-y", "-i", f, outputFileName)

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return outputFileName, nil
}

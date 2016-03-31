package audiotool

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/briannewsom/audiotool/util"
)

func ConcatenateToBytes(f1 string, f2 string, ext string) ([]byte, error) {
	fName, err := Concatenate(f1, f2, ext)

	if err != nil {
		return nil, err
	}

	f, err := os.Open(fName)
	if err != nil {
		return nil, err
	}

	return util.ReadFile(f)
}
func Concatenate(f1 string, f2 string, ext string) (string, error) {
	outputFile, err := ioutil.TempFile("/tmp", "concat-")

	if err != nil {
		return "", err
	}

	outputFileName := outputFile.Name() + ext
	os.Rename(outputFile.Name(), outputFileName)

	log.Printf("Concatenating files %s and %s to %s", f1, f2, outputFileName)

	if err != nil {
		return "", err
	}

	cmd := exec.Command(getAVConvPath(), "-y", "-i",
		"concat:"+f1+"|"+f2, "-c", "copy", outputFileName)

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}

package audiotool

import (
	"io/ioutil"
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

	outputFileName := outputFile.Name() + util.M4aExt

	os.Rename(outputFile.Name(), outputFileName)

	cmd := exec.Command(util.AVConvPath, "-y", "-i", f, outputFileName)

	cmd.CombinedOutput()

	return outputFileName, nil
}

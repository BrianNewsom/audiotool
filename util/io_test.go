package util_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/briannewsom/audiotool/util"
)

func TestCleanTempFiles(t *testing.T) { // inputFileName string, outputFileName string) error {
	f, _ := ioutil.TempFile("/tmp", "test")
	g, _ := ioutil.TempFile("/tmp", "test")

	util.CleanTempFiles(f.Name(), g.Name())

	f1, _ := os.Open(f.Name())

	g1, _ := os.Open(g.Name())

	if f1 != nil || g1 != nil {
		t.Error("CleanTempFiles(f,g) should delete files f and g")
	}
}

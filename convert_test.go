package audiotool_test

import (
	"testing"

	"github.com/briannewsom/audiotool"
	"github.com/briannewsom/audiotool/util"
)

func TestConvertWavToM4a(t *testing.T) {
	f, err := audiotool.ConvertWavToM4a("_test-res/hi.wav")

	if err != nil {
		t.Error("Failed to convert .wav to .m4a ")
	}

	util.CleanTempFiles(f)
}
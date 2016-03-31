package audiotool_test

import (
	"testing"

	"github.com/briannewsom/audiotool"
	"github.com/briannewsom/audiotool/util"
)

func ExampleConcatenate() {
	audiotool.Concatenate("_test-res/hi.wav", "_test-res/how-are-you.wav", ".wav")
}

func ExampleConcatenateToBytes() {
	audiotool.ConcatenateToBytes("_test-res/hi.wav", "_test-res/how-are-you.wav", ".wav")
}

func TestConcatenate(t *testing.T) {
	f, err := audiotool.Concatenate("_test-res/hi.wav", "_test-res/how-are-you.wav", ".wav")

	if err != nil {
		t.Error("Failed to concatenate audio files - " + err.Error())
	}

	util.CleanTempFiles(f)
}

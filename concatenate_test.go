package audiotool_test

import (
	"testing"

	"github.com/briannewsom/audiotool"
	// "github.com/briannewsom/audiotool/util"
)

func ExampleConcatenate() {
	audiotool.Concatenate(".wav", "_test-res/hi.wav", "_test-res/how-are-you.wav")
}

func ExampleConcatenateToBytes() {
	audiotool.ConcatenateToBytes(".wav", "_test-res/hi.wav", "_test-res/how-are-you.wav")
}

func TestConcatenate(t *testing.T) {
	_, err := audiotool.Concatenate(".wav", "_test-res/hi.wav", "_test-res/how-are-you.wav")

	if err != nil {
		t.Error("Failed to concatenate audio files - " + err.Error())
	}

	/* Concatenate more than 2 files */
	_, err = audiotool.Concatenate(".wav", "/home/brian/Downloads/be6b9d72-0550-45fa-ae99-157693ac2500-0-t2s.wav",
		"/home/brian/Downloads/be6b9d72-0550-45fa-ae99-157693ac2500-1-t2s.wav",
		"/home/brian/Downloads/be6b9d72-0550-45fa-ae99-157693ac2500-2-t2s.wav")

	if err != nil {
		t.Error("Failed to concatenate four audio files - " + err.Error())
	}

	// util.CleanTempFiles(f, f2)
}

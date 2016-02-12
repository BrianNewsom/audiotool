package compress_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/satori/go.uuid"

	"github.com/briannewsom/audiotool/compress"
)

var content, err = ioutil.ReadFile("_test-res/audio.mp4")

func TestCompress(t *testing.T) {
	id := uuid.NewV4().String()
	_, err := compress.Compress(nil, id, "64k")

	if err == nil {
		print(err)
		t.Error("Compress(nil, uuid, bitrate) should throw an error, got nil")
	}

	_, err = compress.Compress(content, "", "64k")

	if err == nil {
		print(err)
		t.Error("Compress(content, \"\", bitrate) should throw an error, got nil")
	}

	_, err = compress.Compress(content, id, "bogus")

	if err == nil {
		print(err)
		t.Error("Compress(content, uuid, \"bogus\") should throw an error, got nil")
	}

	/* Sucessful compression
	c, err := compress.Compress(content, id, "128k")

	if c == nil {
		t.Error("Compress(content, uuid, bitrate) should return content, got nil")
	}
	*/

}

func TestCleanTempFiles(t *testing.T) { // inputFileName string, outputFileName string) error {
	f, _ := ioutil.TempFile("/tmp", "test")
	g, _ := ioutil.TempFile("/tmp", "test")

	compress.CleanTempFiles(f.Name(), g.Name())

	f1, _ := os.Open(f.Name())

	g1, _ := os.Open(g.Name())

	if f1 != nil || g1 != nil {
		t.Error("CleanTempFiles(f,g) should delete files f and g")
	}
}

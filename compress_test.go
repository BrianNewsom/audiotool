package audiotool_test

import (
	"io/ioutil"
	"testing"

	"github.com/satori/go.uuid"

	"github.com/briannewsom/audiotool"
)

var content, err = ioutil.ReadFile("_test-res/audio.mp4")

func TestCompress(t *testing.T) {
	id := uuid.NewV4().String()
	_, err := audiotool.Compress(nil, id, "64k")

	if err == nil {
		print(err)
		t.Error("Compress(nil, uuid, bitrate) should throw an error, got nil")
	}

	_, err = audiotool.Compress(content, "", "64k")

	if err == nil {
		print(err)
		t.Error("Compress(content, \"\", bitrate) should throw an error, got nil")
	}

	_, err = audiotool.Compress(content, id, "bogus")

	if err == nil {
		print(err)
		t.Error("Compress(content, uuid, \"bogus\") should throw an error, got nil")
	}

	/* Sucessful compression
	c, err := audiotool.Compress(content, id, "128k")

	if c == nil {
		t.Error("Compress(content, uuid, bitrate) should return content, got nil")
	}
	*/

}

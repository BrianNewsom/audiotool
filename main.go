package audiotool

import (
	"os"
)

const DefaultAVConvPath = "/usr/bin/avconv"
const DefaultFFMpegPath = "/usr/local/bin/ffmpeg"
const M4aExt = ".m4a"

func getFFMpegPath() string {
	e := os.Getenv("FFMPEG_PATH")

	if e != "" {
		return e
	}

	return DefaultFFMpegPath
}

func getAVConvPath() string {
	e := os.Getenv("AVCONV_PATH")

	if e != "" {
		return e
	}

	return DefaultAVConvPath
}

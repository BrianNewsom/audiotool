package audiotool

import (
	"os"
)

const DefaultAVConvPath = "/usr/bin/avconv"
const M4aExt = ".m4a"

func getAVConvPath() string {
	e := os.Getenv("AVCONV_PATH")

	if e != "" {
		return e
	}

	return DefaultAVConvPath
}

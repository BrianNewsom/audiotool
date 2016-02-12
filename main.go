package audiotool

import "flag"

func main(){
	var avconvPath = flag.String("avconvPath", "/usr/bin/avconv", "The location of the avconv executable on this machine")

	flag.Parse()

	print("You gave ", *avconvPath)
}
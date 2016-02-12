package compress

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"audiotool/common"
)

const avconvPath = "/usr/bin/avconv"
const tmpDir = "/tmp"

func Compress(content []byte, uuid string, bitrate string) ([]byte, error) {
	var inputFileName string
	var outputFileName string

	if uuid != "" {
		inputFileName = "/tmp/raw-" + uuid + ".mp4"
		outputFileName = "/tmp/" + uuid + ".mp4"
	} else {
		return nil, errors.New("No uuid given")
	}

	defer CleanTempFiles(inputFileName, outputFileName)

	if content != nil {
		common.WriteFile(inputFileName, content)
	} else {
		return nil, errors.New("No content supplied to compress")
	}

	err := ChangeBitrate(inputFileName, outputFileName, bitrate)

	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadFile(outputFileName)

	if err != nil {
		return nil, err
	}

	return d, nil

}

func CleanTempFiles(inputFileName string, outputFileName string) error {
	err := os.Remove(inputFileName)
	err = os.Remove(outputFileName)

	return err
}

func ChangeBitrate(inputFileName string, outputFileName string, bitrate string) error {
	// Returns outputFileName
	cmd := exec.Command(avconvPath, "-y", "-i", inputFileName, "-strict", "experimental", "-b", bitrate, outputFileName)

	err := cmd.Start()
	if err != nil {
		return err
	}

	fmt.Printf("Processing Audio File...\n")

	err = cmd.Wait()

	if err != nil {
		return err
	}

	fmt.Printf("Successfully processed audio file.\n")

	return nil

}

/* For converting previously ignored files - to be used later

func DownloadToFile(uri string, dst string) {
	fmt.Printf("DownloadToFile From: %s.\n", uri)
	if d, err := HTTPDownload(uri); err == nil {
		fmt.Printf("downloaded %s.\n", uri)
		if WriteFile(dst, d) == nil {
			fmt.Printf("saved %s as %s\n", uri, dst)
		}
	}
}

func HTTPDownload(uri string) ([]byte, error) {
	fmt.Printf("HTTPDownload From: %s.\n", uri)
	res, err := http.Get(uri)
	if err != nil {
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
	}
	fmt.Printf("ReadFile: Size of download: %d\n", len(d))
	return d, err
}

func DownloadAndCompress(Uuid string) error {
	audioUrl := db.S3BucketPath + "/" + Uuid + ".mp4"
	inputFileName := "/tmp/raw-" + Uuid + ".mp4"
	outputFileName := "/tmp/" + Uuid + ".mp4"

	fmt.Printf("%s - %s", inputFileName, outputFileName)

	DownloadToFile(audioUrl, inputFileName)

	err := ChangeBitrate(inputFileName, outputFileName, bitrate)

	if err != nil {
		return err
	}

	// Upload w/ s3

	// First read file into data slice
	d, _ := ioutil.ReadFile(outputFileName)

	db.UploadToS3(d, outputFileName)

	// Change and update url

	// Clean up files

	return nil
}

/*
func DownloadAudio(url string) []byte {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	file := resp.body
}
*/

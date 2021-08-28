package main

import (
	"io"
	"os"
	"fmt"
	"time"
	"context"
	"strings"
	"strconv"

	"os/exec"
	"io/ioutil"

	"github.com/kkdai/youtube"
        "gopkg.in/vansante/go-ffprobe.v2"
)

func GetDuration(path string) (string, error) {
        ctx, cancelFn := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancelFn()

        data, err := ffprobe.ProbeURL(ctx, path)
        if err != nil {
                return "", err
        }

        duration := fmt.Sprintf("%f", data.Format.Duration().Seconds())
        parts := strings.Split(duration, ".")

        return parts[0], nil
}

func CleanURL(url *string) {
        parts := strings.Split(*url, "=")
        *url = parts[1] 
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <File_or_Youtube_URL>")
		fmt.Println("Ex: " + os.Args[0] + " https://www.youtube.com/watch?v=NTT-SXaP4EA")
		fmt.Println("")
		return
	}

	videoID := os.Args[1]

	if _, err := os.Stat(videoID); os.IsNotExist(err) {
		CleanURL(&videoID)

		client := youtube.Client{}

		video, err := client.GetVideo(videoID)
		if err != nil {
			panic(err)
		}

		cnt := -1
		keep := -1
		for _, format := range video.Formats {
			cnt++
			if strings.Contains(format.MimeType, "audio/mp4") {
				keep = cnt
				break
			}
		}

		if keep == -1 {
			fmt.Println("No Audio Found")
			return
		}

		stream, _, err := client.GetStream(video, &video.Formats[keep])
		if err != nil {
			panic(err)
		}

		file, err := os.Create("/tmp/video.mp4")
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}
		file.Close()

		duration, err := GetDuration("/tmp/video.mp4")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dur, err := strconv.Atoi(duration)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for {
			cmd := exec.Command("/usr/bin/ffplay", "-nodisp", "/tmp/video.mp4")
			go func() {
				time.Sleep(time.Duration(dur) * time.Second)
				cmd.Process.Kill()
			}()
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(output)
				fmt.Println(err.Error())
			}

			_ = output
		}
	} else {
		dat, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		lines := strings.Split(string(dat), "\n")
		for _, line := range lines {
			videoID := line
			CleanURL(&videoID)

			client := youtube.Client{}

			video, err := client.GetVideo(videoID)
			if err != nil {
				panic(err)
			}

			cnt := -1
			keep := -1
			for _, format := range video.Formats {
				cnt++
				if strings.Contains(format.MimeType, "audio/mp4") {
					keep = cnt
					break
				}
			}

			if keep == -1 {
				fmt.Println("No Audio Found")
				return
			}

			stream, _, err := client.GetStream(video, &video.Formats[keep])
			if err != nil {
				panic(err)
			}

			file, err := os.Create("/tmp/video.mp4")
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(file, stream)
			if err != nil {
				panic(err)
			}
			file.Close()

			duration, err := GetDuration("/tmp/video.mp4")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			dur, err := strconv.Atoi(duration)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			cmd := exec.Command("/usr/bin/ffplay", "-nodisp", "/tmp/video.mp4")
			go func() {
				time.Sleep(time.Duration(dur) * time.Second)
				cmd.Process.Kill()
			}()

			output, err := cmd.Output()
			if err != nil {
				fmt.Println(output)
				fmt.Println(err.Error())
			}

			_ = output

		}
	}
}


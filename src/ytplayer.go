
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
	"math/rand"

	"github.com/kkdai/youtube"
        "gopkg.in/vansante/go-ffprobe.v2"
)

// I kill Giants. When I swing the Coveleski and Crush

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

	if os.Args[1] == "--random" || os.Args[1] == "-R" {
		for {
			var urllist []string
			var already []int
			
			already = nil
			urllist = nil

			urllist = append(urllist, "https://www.youtube.com/watch?v=2vlrUVtb6aY")
			urllist = append(urllist, "https://www.youtube.com/watch?v=2l8N3rYpKYg")
			urllist = append(urllist, "https://www.youtube.com/watch?v=qwz8TK-yBlI")
			urllist = append(urllist, "https://www.youtube.com/watch?v=LIME1erteP0")
			urllist = append(urllist, "https://www.youtube.com/watch?v=89gnG9ErOaA")
			urllist = append(urllist, "https://www.youtube.com/watch?v=QZzIafncdTs")
			urllist = append(urllist, "https://www.youtube.com/watch?v=rl9FFZZnWWo")
			urllist = append(urllist, "https://www.youtube.com/watch?v=gD_0U5l7jk4")
			urllist = append(urllist, "https://www.youtube.com/watch?v=MH9FyLsfDzw")
			urllist = append(urllist, "https://www.youtube.com/watch?v=2hfxB3FAey4")
			urllist = append(urllist, "https://www.youtube.com/watch?v=NBiQPZxd81s")
			urllist = append(urllist, "https://www.youtube.com/watch?v=xFMPBPOy9SI")
			urllist = append(urllist, "https://www.youtube.com/watch?v=J5X0hGtIEp4")
			urllist = append(urllist, "https://www.youtube.com/watch?v=KQcY1aj9XTI")
			urllist = append(urllist, "https://www.youtube.com/watch?v=fgmpWkUcpjo")
			urllist = append(urllist, "https://www.youtube.com/watch?v=RhJWk2g0lfo")
			urllist = append(urllist, "https://www.youtube.com/watch?v=9ywS7DH5dAQ")
			urllist = append(urllist, "https://www.youtube.com/watch?v=AmOiIs_v6EE")
			urllist = append(urllist, "https://www.youtube.com/watch?v=EFfofP8Q1q8")
			urllist = append(urllist, "https://www.youtube.com/watch?v=GwkUq3cXe8o")
			urllist = append(urllist, "https://www.youtube.com/watch?v=-vkyV84Ho7Q")
			urllist = append(urllist, "https://www.youtube.com/watch?v=wEERFBI9eCg") // Old loves, they die hard.  Old lies, they die harder.
			urllist = append(urllist, "https://www.youtube.com/watch?v=rFgpxEFySig")
			urllist = append(urllist, "https://www.youtube.com/watch?v=EpCyqq-jBw4")
			urllist = append(urllist, "https://www.youtube.com/watch?v=ZiD0cX7M6m4")
			urllist = append(urllist, "https://www.youtube.com/watch?v=RZx8_qb8qm8")
			urllist = append(urllist, "https://www.youtube.com/watch?v=akv_L3_qWYA")
			urllist = append(urllist, "https://www.youtube.com/watch?v=r6582oInn18")
			urllist = append(urllist, "https://www.youtube.com/watch?v=nQN9CITV-T4")
			urllist = append(urllist, "https://www.youtube.com/watch?v=ta7xntwQ1vA")
			urllist = append(urllist, "https://www.youtube.com/watch?v=OJuHjsAk_YA")
			urllist = append(urllist, "https://www.youtube.com/watch?v=0qYt1dp818I")
			urllist = append(urllist, "https://www.youtube.com/watch?v=CPEa8mWxYkQ")
			urllist = append(urllist, "https://www.youtube.com/watch?v=YuYumJwIqz8")
			urllist = append(urllist, "https://www.youtube.com/watch?v=LqOfPkHGq9U")
			urllist = append(urllist, "https://www.youtube.com/watch?v=NjLH1McZkFo")
			urllist = append(urllist, "https://www.youtube.com/watch?v=RBuceU6ufiw")
			urllist = append(urllist, "https://www.youtube.com/watch?v=jaKPIp_mO-E")
			urllist = append(urllist, "https://www.youtube.com/watch?v=U5XBtbGkqig")
			urllist = append(urllist, "https://www.youtube.com/watch?v=vaXt0rxC2No")
			urllist = append(urllist, "https://www.youtube.com/watch?v=_ThRAOpaBk8")
			urllist = append(urllist, "https://www.youtube.com/watch?v=v-6zGLtqS_s")
			urllist = append(urllist, "https://www.youtube.com/watch?v=20-txQdOo4Q")
			urllist = append(urllist, "https://www.youtube.com/watch?v=z69_ftoUSwM")
			urllist = append(urllist, "https://www.youtube.com/watch?v=AjhgvLeAxR0")
			urllist = append(urllist, "https://www.youtube.com/watch?v=Z4nsIi4a1TU")
			urllist = append(urllist, "https://www.youtube.com/watch?v=Yiv3ExEr3Nw")
			urllist = append(urllist, "https://www.youtube.com/watch?v=4RF9-efhHuE")
			urllist = append(urllist, "https://www.youtube.com/watch?v=5IUXcOjkOA0")
			urllist = append(urllist, "https://www.youtube.com/watch?v=aUd4cnyHUfE")
			urllist = append(urllist, "https://www.youtube.com/watch?v=p-qnsZhQ_iI")
			urllist = append(urllist, "https://www.youtube.com/watch?v=5dmQ3QWpy1Q")
			urllist = append(urllist, "https://www.youtube.com/watch?v=VDvr08sCPOc")

			rand.Seed(time.Now().UnixNano())
			for {
				if len(already) >= len(urllist) {
					break
				}

				index := rand.Intn(len(urllist))
				flag := false
				for _, al := range already {
					if index == al {
						flag = true
						break
					}
				}

				if flag {
					continue
				}

				already = append(already, index)

				videoID := urllist[index]
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


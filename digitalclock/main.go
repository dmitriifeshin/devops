//go:build !solution

package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	layoutTime  = "15:04:05"
	height      = 12
	width       = 8
	width_colon = 4
)

var nums = map[rune]string{
	'0': Zero,
	'1': One,
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	':': Colon,
}

func printPixel(img *image.RGBA, shiftx, shifty, scale int, col color.RGBA) {
	for i := range scale {
		for j := range scale {
			img.Set(j+shifty, i+shiftx, col)
		}
	}
}

func printDigit(img *image.RGBA, digit rune, scale int, shift int) {
	i := 0
	j := shift
	for _, ch := range nums[digit] {
		if ch == '\n' {
			i += scale
			j = shift
			continue
		}
		color := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
		if ch == '1' {
			color = Cyan
		}
		printPixel(img, i, j, scale, color)
		j += scale
	}
}

func printTime(img *image.RGBA, timeStr string, scale int) {
	shiftW := 0
	for _, digit := range timeStr {
		printDigit(img, digit, scale, shiftW)
		if digit == ':' {
			shiftW += width_colon * scale
		} else {
			shiftW += width * scale
		}
	}
}

func main() {

	setError := func(w http.ResponseWriter, msg string) {
		w.WriteHeader(400)
		w.Write([]byte(msg))
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		// panic(r.URL.RawQuery)
		m, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			setError(w, "Invalid parameters")
			return
		}
		k := 1
		if strK, ok := m["k"]; ok {
			k, err = strconv.Atoi(strK[0])
			if err != nil || k < 1 || k > 30 {
				setError(w, "Invalid k")
				return
			}
		}
		timeStr := time.Now().Format(layoutTime)
		if strTime, ok := m["time"]; ok {
			timeStr = strTime[0]
			timeSplit := strings.Split(timeStr, ":")
			if len(timeSplit) != 3 || len(timeStr) != len(layoutTime) {
				setError(w, "Invalid time0")
				return
			}
			for i, j := range timeSplit {
				num, err := strconv.Atoi(j)
				if err != nil {
					setError(w, "Invalid time1")
					return
				}
				if i == 0 && (num > 23 || num < 0) {
					setError(w, "Invalid time2")
					return
				}
				if (i == 1 || i == 2) && (num > 59 || num < 0) {
					setError(w, "Invalid time3")
					return
				}
			}
		}

		img := image.NewRGBA(image.Rect(0, 0, (6*width+2*width_colon)*k, height*k))
		printTime(img, timeStr, k)

		buffer := new(bytes.Buffer)
		if err := png.Encode(buffer, img); err != nil {
			log.Println("unable to encode image.")
		}

		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(buffer.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
	}

	go http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	// flag.StringVar(&port, "port", "8000", "Port app")
	// flag.Parse()

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x94, 0x00, 0xd3, 0xff},
	color.RGBA{0x4b, 0x00, 0x82, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x7f, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
}

// e.g., http://localhost:8000/?delay=10
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		delays, ok := r.Form["delay"]
		delay := 8
		if ok && len(delays) > 0 {
			d, err := strconv.Atoi(delays[0])
			if err == nil {
				delay = d
			}
		}
		lissajous(w, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, delay int) {
	const (
		cycles  = 7     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	numColors := len(palette) - 1
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			index := uint8(t/2/math.Pi) % uint8(numColors)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				index+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

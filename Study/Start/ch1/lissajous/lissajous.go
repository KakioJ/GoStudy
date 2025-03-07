// Lissajous generates GIF animations of random Lissajous figures.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	// "os"
	// "time"
)

var palette = []color.Color{color.White, color.RGBA{43, 25, 0, 100}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// func main() {
// 	// The sequence of images is deterministic unless we seed
// 	// the pseudo-random number generator using the current time.
// 	// Thanks to Randall McPherson for pointing out the omission.
// 	// 设置随机数种子，使用当前时间的Unix时间戳（纳秒级）来确保每次运行生成的图像序列都是随机的
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	// 调用lissajous函数，将生成的图像输出到标准输出（os.Stdout）
// 	lissajous(os.Stdout)
// }

func Lissajous(out io.Writer, cycles int) {
	// 定义常量
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	randColor := color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}

	palette = append(palette, randColor)

	// 生成一个随机频率
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	// 生成每一帧动画
	for i := 0; i < nframes; i++ {
		// 定义图像的矩形区域
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 创建一个新的调色板图像
		img := image.NewPaletted(rect, palette)
		// 计算每个点的位置并设置颜色
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(len(palette)-1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

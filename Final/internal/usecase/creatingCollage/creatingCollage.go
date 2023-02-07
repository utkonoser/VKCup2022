package creatingCollage

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// imagesCache кеш изображений
var imagesCache = sync.Map{}

// removeRedGreen функция принимающая изображение и возвращающая новое с удалением -Red и -Green
// составляющих каждого пикселя
func removeRedGreen(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewNRGBA(bounds)

	rowCh := make(chan int)
	resultCh := make(chan *image.NRGBA)

	go func() {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rowCh <- y
		}
		close(rowCh)
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for y := range rowCh {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					oldPixel := img.At(x, y)
					_, _, b, a := oldPixel.RGBA()
					newPixel := color.RGBA{B: uint8(b >> 8), A: uint8(a >> 8)}
					newImg.Set(x, y, newPixel)
				}
				resultCh <- newImg
			}
		}()
	}

	for i := 0; i < bounds.Max.Y-bounds.Min.Y; i++ {
		<-resultCh
	}

	return newImg
}

// downloadImage загружает url-ссылку на изображение и возвращает само изображение в виде image.Image
func downloadImage(url string) (image.Image, error) {
	if val, ok := imagesCache.Load(url); !ok {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err = resp.Body.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		img, err := png.Decode(resp.Body)
		if err != nil {
			return nil, err
		}

		result := removeRedGreen(img)

		imagesCache.Store(url, result)

		return result, nil
	} else {
		return val.(image.Image), nil
	}
}

// Collage функция принимающая канал с url-ссылками на изображение и возвращающая коллаж из всех обработанных изображений
func Collage(imageUrls chan string) image.Image {

	var chanImages = make(chan image.Image, 100_000)
	var wg sync.WaitGroup

	startChangeImages := time.Now()
	log.Println("Loading images and changing pixels in them...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		for imageUrl := range imageUrls {
			newImg, err := downloadImage(imageUrl)
			if err != nil {
				log.Fatal("incorrect download images")
			}

			chanImages <- newImg
		}
	}()
	wg.Wait()
	close(chanImages)
	log.Printf("End of changing images in: %s\n", time.Since(startChangeImages))

	log.Println("Creating a collage...")
	startCollage := time.Now()

	cols := 32
	rows := len(chanImages) / cols
	if len(chanImages)%32 != 0 {
		rows++
	}

	collage := image.NewNRGBA(image.Rect(0, 0, cols*512, rows*512))

	i := 0
	for src := range chanImages {
		x := (i % cols) * 512
		y := (i / cols) * 512
		i++
		draw.Draw(collage, image.Rect(x, y, x+512, y+512), src, image.Point{}, draw.Src)
	}

	log.Printf("The creation of the collage took: %s\n", time.Since(startCollage))

	return collage
}

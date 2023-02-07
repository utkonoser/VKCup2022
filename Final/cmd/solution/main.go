package main

import (
	"VKCupFinal/internal/usecase/creatingCollage"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("please add url and filename in subcommands")
	}
	url, filename := os.Args[1], os.Args[2]

	log.Println("Application start...")
	start := time.Now()
	go func() {
		defer close(creatingCollage.ChanPics)
		creatingCollage.ProcessingAllPages(url)
	}()

	collage := creatingCollage.Collage(creatingCollage.ChanPics)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("PNG image encoding...")
	if err = png.Encode(file, collage); err != nil {
		log.Fatal(err)
	}
	log.Printf("Application running time %s \n", time.Since(start))
}

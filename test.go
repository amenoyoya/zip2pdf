package main

import (
	_ "io"
	_ "os"
	"fmt"
	"log"

	"archive/zip"
	
	"image"
	_ "image/png"
	_ "image/jpeg"
	_ "image/gif"
	_ "golang.org/x/image/webp" // $ go get golang.org/x/image/webp

	"github.com/signintech/gopdf" // $ go get github.com/signintech/gopdf
)

func main() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// PDF handler
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		fmt.Printf("- File: %#v\n", f)

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- Resource: %#v\n", rc)
		defer rc.Close()

		// _, err = io.Copy(os.Stdout, rc)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// Decode image
		img, _, err := image.Decode(rc)
		if err != nil {
			fmt.Printf("File is not a image\n")
		} else {
			fmt.Printf("Image file: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())

			// Add image to pdf
			pdf.AddPageWithOption(gopdf.PageOption{
				PageSize: &gopdf.Rect{
					W: float64(img.Bounds().Dx()) * 0.562,
					H: float64(img.Bounds().Dy()) * 0.562,
				},
			})
			pdf.ImageFrom(img, 0, 0, nil)
		}
		fmt.Println()
	}
	// Save pdf file
	pdf.WritePdf("test.pdf")
}
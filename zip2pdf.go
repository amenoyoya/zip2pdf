package main

import (
    "flag"
    "fmt"
    "log"
    "path/filepath"
    "archive/zip"

    "image"
    _ "image/png"
    _ "image/jpeg"
    _ "image/gif"
    _ "golang.org/x/image/webp"

    "github.com/signintech/gopdf"
)

// Append image file to pdf object
func appendImageToPdf(pdf *gopdf.GoPdf, img image.Image) error {
    pdf.AddPageWithOption(gopdf.PageOption{
        PageSize: &gopdf.Rect{
            W: float64(img.Bounds().Dx()) * 0.562,
            H: float64(img.Bounds().Dy()) * 0.562,
        },
    })
    return pdf.ImageFrom(img, 0, 0, nil)
}

// Create image.Image object from zip.File object
func createImageFromFileInZip(f *zip.File) (img image.Image, err error) {
    rc, err := f.Open()
    if err != nil {
        return nil, err
    }
    defer rc.Close()
    
    // Decode image
    img, _, err = image.Decode(rc)
    return img, err
}

// Convert zip file to pdf file
func convertZipToPdf(zipFile string) error {
    pdfFile := filepath.Base(zipFile)
    pdfFile = pdfFile[0:len(pdfFile) - len(filepath.Ext(zipFile))] + ".pdf"

    fmt.Printf("Pdf: %s\n", pdfFile)

    // Open a zip archive for reading
    r, err := zip.OpenReader(zipFile)
    if err != nil {
        return fmt.Errorf("The target may not be a zip file: %s", zipFile)
    }
    defer r.Close()

    // PDF handler
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

    // Iterate through the files in the zip file
    for _, f := range r.File {
        img, err := createImageFromFileInZip(f)
        if err != nil {
            fmt.Printf("- Skip: %s\n", f.Name)
            continue
        }

        err = appendImageToPdf(&pdf, img)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("- Add: %s\n", f.Name)
    }

    // Save pdf file
    return pdf.WritePdf(pdfFile)
}

// main process (entry point)
func main() {
    flag.Parse()
    for _, zipFile := range flag.Args() {
        err := convertZipToPdf(zipFile)
        if err != nil {
            fmt.Printf("- Error: %s\n", err)
        }
    }
}
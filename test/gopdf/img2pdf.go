package main

import (
    "github.com/signintech/gopdf"
    "os"
    "image"
    _ "image/jpeg"
)

// 画像読み込み
func ReadImage(filename string) image.Image {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        panic(err)
    }
    return img
}

func main() {
    img := ReadImage("./card.jpg")

    // PDF作成
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{
        PageSize: *gopdf.PageSizeA4,
    })
    pdf.AddPageWithOption(gopdf.PageOption{
        PageSize: &gopdf.Rect{
            W: float64(img.Bounds().Dx()) * 0.562,
            H: float64(img.Bounds().Dy()) * 0.562,
        },
    })
    pdf.ImageFrom(img, 0, 0, nil)

    // 2ページ目
    img = ReadImage("./halloween.jpg")
    pdf.AddPageWithOption(gopdf.PageOption{
        PageSize: &gopdf.Rect{
            W: float64(img.Bounds().Dx()) * 0.562,
            H: float64(img.Bounds().Dy()) * 0.562,
        },
    })
    pdf.ImageFrom(img, 0, 0, nil)

    // PDF保存
    pdf.WritePdf("img.pdf")
}
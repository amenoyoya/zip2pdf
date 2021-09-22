package main

import (
    "github.com/signintech/gopdf"
)

func main() {
    // gopdf のオブジェクトを作成 --- (*1)
    pdf := gopdf.GoPdf{}
    // A4のページを作る --- (*2)
    pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
    // TTFフォントを取り込む --- (*3)
    err := pdf.AddTTFFont("ipaexg", "./ipaexg.ttf"    )
    if err != nil {
        panic(err)
    }
    err = pdf.SetFont("ipaexg", "", 26) // フォントサイズを選択
    if err != nil {
        panic(err)
    }
    //　文字を書き込む --- (*4)
    pdf.AddPage()
    pdf.SetX(80)
    pdf.SetY(130)
    pdf.Cell(nil, "洞察力があればすぐに怒ることはない")
    pdf.SetX(80)
    pdf.SetY(180)
    pdf.Cell(nil, "過ちを見過ごす人は美しい")
    // PDFをファイルに書き出す --- (*5)
    pdf.WritePdf("hello.pdf")
}
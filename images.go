package main

// NOTE this will only work in the tour.golang.org site, as it uses 
// tour specific package to create the actual PNG
import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
    w,h int
    genR func(x,y int) (val uint8)
    genG func(x,y int) (val uint8)
    genB func(x,y int) (val uint8)
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{
        img.genR(x,y),
        img.genG(x,y),
        img.genB(x,y),
        255,
    }
}

func main() {
    m := Image{
        128,
        128,
        func(x,y int) (val uint8) {return uint8(4*(x+y))},
        func(x,y int) (val uint8) {return uint8(x)},
        func(x,y int) (val uint8) {return uint8(y)},
    }
    pic.ShowImage(m)
}

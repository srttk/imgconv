package main

import (
	"log"
	"flag"

	"github.com/srttk/imgconv/converter"
)

var srcType string
var outType string
var srcDir  string

func init() {
	flag.StringVar(&srcType, "src", "jpeg", "set source image type")
	flag.StringVar(&srcType, "s", "jpeg", "shorthand 'src flag'")
	flag.StringVar(&outType, "out", "png", "set output image type")
	flag.StringVar(&outType, "o", "png", "shorthand 'out flag")
}

func main() {
	flag.Parse()
	srcDir = flag.Arg(0)
	converter, err := converter.NewImages(srcType, srcDir)
	if err != nil {
		log.Fatal(err)
	}
	if err := converter.ConvertTo(outType); err != nil {
		log.Fatal(err)
	}
}

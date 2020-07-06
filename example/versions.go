package main

import (
	"log"

	"github.com/limestone0o0/goav2/avcodec"
	"github.com/limestone0o0/goav2/avdevice"
	"github.com/limestone0o0/goav2/avfilter"
	"github.com/limestone0o0/goav2/avformat"
	"github.com/limestone0o0/goav2/avutil"
	"github.com/limestone0o0/goav2/swresample"
	"github.com/limestone0o0/goav2/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}

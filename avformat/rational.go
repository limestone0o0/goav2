// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"
import "github.com/limestone0o0/goav2/avcodec"

func newRational(r C.struct_AVRational) avcodec.Rational {
	return avcodec.NewRational(int(r.num), int(r.den))
}

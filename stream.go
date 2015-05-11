package golibrdf

// #cgo linux pkg-config: redland
// #cgo LDFLAGS: -lrdf
// #include <librdf.h>
import "C"

import (
	"runtime"
)

type Stream struct {
	librdf_stream *C.librdf_stream
}

func createStream(s *C.librdf_stream) *Stream {
	if s == nil {
		return nil
	}
	stream := &Stream{s}
	runtime.SetFinalizer(stream, (*Stream).Free)
	return stream
}

func (stream *Stream) Free() {
	if stream.librdf_stream != nil {
		stream.librdf_stream = nil
		C.librdf_free_stream(stream.librdf_stream)
	}
}

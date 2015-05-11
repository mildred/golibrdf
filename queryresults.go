package golibrdf

// #cgo linux pkg-config: redland
// #cgo LDFLAGS: -lrdf
// #include <librdf.h>
import "C"

import (
	"runtime"
)

type Results struct {
	librdf_results *C.librdf_query_results
}

func createResults(r *C.librdf_query_results) *Results {
	if r == nil {
		return nil
	}
	results := &Results{r}
	runtime.SetFinalizer(results, (*Results).Free)
	return results
}

// Test if the result is a variable bindings format
func (res *Results) IsBindings() bool {
	return C.librdf_query_results_is_bindings(res.librdf_results) != 0
}

// Test if result is boolean format
func (res *Results) IsBoolean() bool {
	return C.librdf_query_results_is_boolean(res.librdf_results) != 0
}

// Test if result is RDF graph format
func (res *Results) IsGraph() bool {
	return C.librdf_query_results_is_graph(res.librdf_results) != 0
}

// Test if results is a syntax. If this function returns true, the only
// available form this query is a syntax can be serialized using one of the
// query_result_formater ckass lethods or with
// librdf_query_results_to_counted_strings, librdf_query_results_to_strings,
// librdf_query_results_to_file_handle or librdf_query_results_to_file
func (res *Results) IsSyntax() bool {
	return C.librdf_query_results_is_syntax(res.librdf_results) != 0
}

// Only meaningful if IsGraph() is true
func (res *Results) AsStream() *Stream {
	return createStream(C.librdf_query_results_as_stream(res.librdf_results))
}

func (res *Results) Free() {
	if res.librdf_results != nil {
		C.librdf_free_query_results(res.librdf_results)
		res.librdf_results = nil
	}
}

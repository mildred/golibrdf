/*
*
* This file forms part of the golibrdf package containing go language bindings,
* tests and examples for the Redland RDF library.
*
* Please refer to http://librdf.org for copyright and licence information 
* on the Redland libraries that this package wraps 
*
* This golibrdf package is: 
* 	Copyright (C) 2013, Phillip Pettit http://ppettit.net/
* 
* This package is licensed under the following three licenses as alternatives:
* 1. GNU Lesser General Public License (LGPL) V2.1 or any newer version
* 2. GNU General Public License (GPL) V2 or any newer version
* 3. Apache License, V2.0 or any newer version
*
* You may not use this file except in compliance with at least one of
* the above three licenses.
*
 */

package golibrdf

// #cgo linux pkg-config: redland raptor2
// #cgo LDFLAGS: -lrdf
// #include <stdlib.h>
// #include <string.h>
// #include <strings.h>
// #include <librdf.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

type Hash struct {
	librdf_hash *C.librdf_hash
}

func NewHash(world *World) (*Hash, error) {
	hash := Hash{}
	hash.librdf_hash = C.librdf_new_hash(world.librdf_world, nil)

	if hash.librdf_hash == nil {
		return nil, errors.New("Unable to make new hash.  Call to librdf_new_hash failed.")
	}

	// set the finalizer so that free call occurs as required
	runtime.SetFinalizer(&hash, (*Hash).Free)

	return &hash, nil
}

func (hash *Hash) PutStrings(key, value string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	if C.librdf_hash_put_strings(hash.librdf_hash, cKey, cValue) != 0 {
		return errors.New("Could not add strings in hash")
	}
	return nil
}

func (hash *Hash) Free() {
	if hash.librdf_hash != nil {
		C.librdf_free_hash(hash.librdf_hash)
		hash.librdf_hash = nil
	}
}

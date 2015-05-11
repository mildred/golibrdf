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
// int go_call_librdf_log_level_func(void *user_data, char *message);
// int cgo_call_librdf_log_level_func(void *user_data, char *message, va_list arguments) {
//   go_call_librdf_log_level_func(user_data, message);
// }
import "C"

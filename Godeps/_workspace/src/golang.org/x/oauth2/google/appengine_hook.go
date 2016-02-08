// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package google

import "github.com/nildev/lib/Godeps/_workspace/src/google.golang.org/appengine"

func init() {
	appengineTokenFunc = appengine.AccessToken
}

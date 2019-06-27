// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"crypto/x509"
)

// Note(kyle): not sure how this works on windows, or if this does.
func initSystemRoots() []*x509.Certificate {
	return nil
}

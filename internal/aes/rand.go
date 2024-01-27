// Copyright (c) 2024 Alexej Disterhoft
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT

package aes

import "crypto/rand"

// randReader is an interface that abstracts the rand.Reader used by the AES
// implementation. This is mostly used to allow mocking the rand.Read function
// in tests.
type randReader interface {
	Read(p []byte) (n int, err error)
}

// cryptoRandReader is a randReader that uses crypto/rand.Read to read random
// bytes in a cryptographically secure way.
type cryptoRandReader struct{}

// Read implements the randReader interface by calling crypto/rand.Read. See
// the documentation of crypto/rand.Read for more details.
func (r *cryptoRandReader) Read(p []byte) (n int, err error) {
	return rand.Read(p)
}

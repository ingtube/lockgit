// Copyright © 2018 Jesse Swidler <jswidler@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package content

import (
	"fmt"
)

type KeyLoadError struct {
	message string
}

func (err *KeyLoadError) Error() string {
	return err.message
}

func IsKeyLoadError(err error) bool {
	if err != nil {
		switch err.(type) {
		case *KeyLoadError:
			return true
		}
	}
	return false
}

type ManifestLoadError struct {
	path   string
	reason string
}

func (err *ManifestLoadError) Error() string {
	return fmt.Sprintf("error loading manifest at %s: %s", err.path, err.reason)
}

func IsManifestLoadError(err error) bool {
	if err != nil {
		switch err.(type) {
		case *ManifestLoadError:
			return true
		}
	}
	return false
}

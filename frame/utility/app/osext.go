// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Extensions to the standard "os" package.
package app

import (
	"path/filepath"
)

// GetAppPath returns an absolute path that can be used to
// re-invoke the current program.
// It may not be valid after the current program exits.
func GetAppPath() (string, error) {
	p, err := executable()
	return filepath.Clean(p), err
}

// Returns same path as GetAppPath, returns just the folder
// path. Excludes the executable name.
func GetAppFolder() (string, error) {
	p, err := GetAppPath()
	if err != nil {
		return "", err
	}
	folder, _ := filepath.Split(p)
	return folder, nil
}

// Returns the executable file name.
func GetAppName() (string, error) {
	p, err := GetAppPath()
	if err != nil {
		return "", err
	}
	return filepath.Base(p), nil
}

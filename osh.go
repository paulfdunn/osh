// Package osh extends the base os package.
// osh is hosted at https://github.com/paulfdunn/osh; please see the repo
// for more information
package osh

import (
	"os"
	"path/filepath"
)

// RemoveAllFiles will remove all files meeting the Glob pattern.
func RemoveAllFiles(pattern string) error {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}

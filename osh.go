// Package osh extends the base os package.
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

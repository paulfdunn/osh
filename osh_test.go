package osh

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveAllFiles(t *testing.T) {
	tempDir := t.TempDir()
	testFiles := []string{"killme01.txt", "killme02.txt", "killme03.txt"}
	for _, v := range testFiles {
		_, err := os.Create(filepath.Join(tempDir, v))
		if err != nil {
			t.Errorf("error creating file, error:%v", err)
			return
		}
	}

	RemoveAllFiles(filepath.Join(tempDir, "killme01*"))
	files, err := filepath.Glob(filepath.Join(tempDir, "killme*"))
	if err != nil {
		t.Errorf("getting file list, error: %v", err)
		return
	}
	if len(files) != 2 {
		t.Errorf("wrong number of files, len=%d", len(files))
		return
	}

	RemoveAllFiles(filepath.Join(tempDir, "killme0*"))
	files, err = filepath.Glob(filepath.Join(tempDir, "killme*"))
	if err != nil {
		t.Errorf("getting file list, error: %v", err)
		return
	}
	if len(files) != 0 {
		t.Errorf("wrong number of files, len=%d", len(files))
		return
	}

}

package controller

import (
	"os"
	"path/filepath"
)

func ClearDirectory(dir string) error {
	// Open the directory.
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	// Read all file names from the directory.
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Loop over the file names and delete each one.
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}

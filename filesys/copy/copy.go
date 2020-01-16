package copy

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Copy a file or directory recursively.
// Supports symlinks and maintains Permissions
func Copy(src, dst string) error {
	var inf os.FileInfo
	var err error

	if inf, err = os.Lstat(src); err != nil {
		return err
	}

	return copy(src, dst, inf)
}

// copy called recursively, checks the file mode and calls the appropriate function
// Only call this function from copy.go, otherwise use the Copy() function
func copy(src, dst string, inf os.FileInfo) error {
	if inf.Mode()&os.ModeSymlink != 0 {
		return symlink(src, dst)
	}

	if inf.IsDir() {
		return dir(src, dst, inf)
	}

	return file(src, dst, inf)
}

// symlink copies the symlink destination of src and creates a new symlink (dst) pointing to the same location as src
// ie. It copies the symlink, not files
func symlink(src, dst string) error {
	var err error
	var symLinkTarget string

	if symLinkTarget, err = os.Readlink(src); err != nil {
		return err
	}

	return os.Symlink(symLinkTarget, dst)
}

// dir recursively copies a directory
// This function should only be called as a result of calling the exported Copy() function
func dir(src, dst string, inf os.FileInfo) error {
	var err error
	var directories []os.FileInfo

	if err = os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	defer os.Chmod(dst, inf.Mode())

	if directories, err = ioutil.ReadDir(src); err != nil {
		return err
	}

	for _, dirItem := range directories {
		srcPath := path.Join(src, dirItem.Name())
		dstPath := path.Join(dst, dirItem.Name())

		if err := copy(srcPath, dstPath, dirItem); err != nil {
			return err
		}
	}

	return nil
}

// file copies a file (src) to a new location (dst) including permissions
// This function should only be called as a result of calling the exported Copy() function
func file(src, dst string, inf os.FileInfo) error {
	var err error
	var srcFile *os.File
	var newFile *os.File

	if srcFile, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()

	if newFile, err = os.Create(dst); err != nil {
		return err
	}
	defer newFile.Close()

	if _, err = io.Copy(newFile, srcFile); err != nil {
		return err
	}

	return os.Chmod(dst, inf.Mode())
}

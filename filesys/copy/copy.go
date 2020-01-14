package copy

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Dir recursively copies a directory (src) and files to a new directory (dst).
// An error is returned if the src doesn't exist or if there is a problem creating the dst.
// Individual files that fail to copy are skipped over and provided in the errors array
func Dir(src, dst string) ([]error, error) {
	var err error
	var srcinf os.FileInfo
	var directories []os.FileInfo

	if srcinf, err = os.Stat(src); err != nil {
		return []error{}, err
	}

	if err = os.MkdirAll(dst, srcinf.Mode()); err != nil {
		return []error{}, err
	}

	if directories, err = ioutil.ReadDir(src); err != nil {
		return []error{}, err
	}

	var failedList []error

	for _, dirFile := range directories {
		srcPath := path.Join(src, dirFile.Name())
		dstPath := path.Join(dst, dirFile.Name())

		if dirFile.IsDir() {
			if errList, err := Dir(srcPath, dstPath); err != nil {
				failedList = append(failedList, err)
				for _, subErr := range errList {
					failedList = append(failedList, subErr)
				}
			}

		} else {
			if err = File(srcPath, dstPath); err != nil {
				failedList = append(failedList, err)
			}
		}

	}

	return failedList, nil
}

// File copies a file (src) to a new location (dst) including permissions
func File(src, dst string) error {
	var err error
	var srcFile *os.File
	var newFile *os.File
	var srcInf os.FileInfo

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

	if srcInf, err = os.Stat(src); err != nil {
		return err
	}

	return os.Chmod(dst, srcInf.Mode())
}

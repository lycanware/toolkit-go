package copy

import (
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	var err error
	var newFile *os.File

	// Create test directory and files
	if err = os.Mkdir("tmp_test_orig", 0700); err != nil {
		t.Fatal("Error creating test directory")
	}

	if newFile, err = os.Create("tmp_test_orig/1.txt"); err != nil {
		t.Fatal("Error creating test file")
	}

	if err = Copy("tmp_test_orig", "tmp_test_copy"); err != nil {
		t.Fatal("Copy failed: ", err)
	}

	newFile.Close()

	if err = os.RemoveAll("tmp_test_orig"); err != nil {
		t.Fatal("Passed tests but error removing test files. Please delete manually")
	}

	if err = os.RemoveAll("tmp_test_copy"); err != nil {
		t.Fatal("Passed tests but error removing test files. Please delete manually")
	}

}

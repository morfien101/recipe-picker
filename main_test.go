package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Flaque/filet"
)

func prepFiles(t *testing.T, numberOfFiles int) (tempDir string, cleanup func()) {
	tempDir = filet.TmpDir(t, "./")
	for i := 0; i < numberOfFiles; i++ {
		f, err := ioutil.TempFile(tempDir, "example.*.pdf")
		if err != nil {
			t.Logf("Failed to create a file for testing. Error: %s", err)
			t.FailNow()
		}

		_, err = f.Write([]byte("This is a test recipe. Add water make wine!"))
		if err != nil {
			t.Logf("Failed to write into test file. Error: %s", err)
			t.FailNow()
		}
		err = f.Close()
		if err != nil {
			t.Logf("Failed to close test file. Error: %s", err)
			t.FailNow()
		}
	}

	cleanup = func() {
		os.RemoveAll(tempDir)
	}

	return tempDir, cleanup
}

func TestCollectFiles(t *testing.T) {
	filecount := 10

	tempDir, cleanup := prepFiles(t, filecount)
	defer cleanup()

	files, err := collectFilesNames(tempDir)
	if err != nil {
		t.Logf("Failed to collect files. Error: %s", err)
		t.Fail()
	}

	if len(files) != filecount {
		t.Logf("Didn't find enough files. Got: %d, Want: %d", len(files), filecount)
		t.Fail()
	}
}

func TestSelection(t *testing.T) {
	testRuns := 25
	tempDir, cleanup := prepFiles(t, 50)
	defer cleanup()

	files, err := collectFilesNames(tempDir)
	if err != nil {
		t.Logf("Failed to collect files. Error: %s", err)
		t.Fail()
	}

	for x := 0; x < testRuns; x++ {
		expectedFiles := 5
		selectedFiles := selectFiles(expectedFiles, files)
		if len(selectedFiles) != expectedFiles {
			t.Logf("Unexpected number of files selected. Got: %d, Want: %d", len(selectedFiles), expectedFiles)
			t.Fail()
		}

		dupList := []string{}
		for currentPostion, filename := range selectedFiles {
			for index, potentialDuplicate := range selectedFiles {
				if index == currentPostion {
					continue
				}
				if filename == potentialDuplicate {
					dupList = append(dupList, filename)
				}
			}
		}

		if len(dupList) > 0 {
			t.Logf("Found duplicate selections. Got: %s", strings.Join(dupList, ","))
			t.Fail()
		}
	}
}

func TestStripping(t *testing.T) {
	tempDir, cleanup := prepFiles(t, 50)
	defer cleanup()

	files, err := collectFilesNames(tempDir)
	if err != nil {
		t.Logf("Failed to collect files. Error: %s", err)
		t.Fail()
	}
	selectedFiles := selectFiles(10, files)
	for _, file := range selectedFiles {
		if strings.Contains(file, "/") {
			t.Logf("file path contains a /. Got: %s", file)
			t.Fail()
		}
	}
}

func TestBody(t *testing.T) {
	// This is a visual test to see if the body comes out correctly.
	conf := &config{
		From:     "test@test.local",
		Password: "Password123",
		Prefix:   "https://fileserver.com/recipes",
	}
	tempDir, cleanup := prepFiles(t, 50)
	defer cleanup()
	recipes, err := pick(tempDir, 5)
	if err != nil {
		t.Logf("Failed to get test recipes. Error: %s", err)
		t.FailNow()
	}

	output := makeBody(conf, recipes)
	t.Logf(string(output))
}

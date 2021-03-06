package pkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// EditFile will edit the content of a file
func EditFile(filepath, old, new string) error {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.Wrapf(err, "Failed to read the given file, due to:%v", err)
	}
	lines := strings.Split(string(fileData), "\n")

	for i, line := range lines {
		if strings.Contains(line, old) {
			lines[i] = strings.Replace(lines[i], old, new, -1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filepath, []byte(output), 0644)
	if err != nil {
		return errors.Wrapf(err, "Failed to write the data in the given file, due to:%v", err)
	}

	return nil
}

// EditKeyValue will edit the value according to key content of the file
func EditKeyValue(filepath, key, oldvalue, newvalue string) error {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.Wrapf(err, "Failed to read the given file, due to:%v", err)
	}
	lines := strings.Split(string(fileData), "\n")

	for i, line := range lines {
		if strings.Contains(line, key) {
			lines[i+1] = strings.Replace(lines[i+1], oldvalue, newvalue, -1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filepath, []byte(output), 0644)
	if err != nil {
		return errors.Wrapf(err, "Failed to write the data in the given file, due to:%v", err)
	}

	return nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fail to get the data: %w", err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("fail to create the file: %w", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("fail to write the data in file: %w", err)
	}
	return nil
}

package utils

import "errors"

// Update will load a json file in, merge fields where possible, returns true if changes were made, the json and any errors
func Update(json []byte, version string, pretty bool) (updated bool, output []byte, err error) {
	if version != "v1" {
		err = errors.New("version mismatch")
		return updated, output, err
	}

	return updated, output, err
}

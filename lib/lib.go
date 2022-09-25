package lib

import (
	"errors"
	"fmt"
	"os"
)

// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func WriteToConfig(otpStruct OTPStruct, configFilePath string) (bool, error) {

	var configFileByteArray []byte

	if _, err := os.Stat(configFilePath); err == nil {
		// path/to/whatever exists

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		configFileByteArray, err = os.ReadFile(configFilePath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	fmt.Println(configFileByteArray)
	fmt.Println(otpStruct)
	return true, nil
}

// package name must match folder name (except main). File name is not important
// there can not be more than one package per folder, this is why this fileops package has to be in its own folder
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//only functions starting with capital letters are allowed to be exported
func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	//[]byte(balanceText) transforms the string into the byte type required by the method below
	//0644 is a file permission type, permits for owner to read and write and others to only view
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	// _ means that we know the var exists but we don't want to work with it now
	data, err := os.ReadFile(fileName)

	// nil stands for when we don't have anhy useful value, err will be an error if it equals nil
	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	valueText := string(data)                       // from []byte to string
	value, err := strconv.ParseFloat(valueText, 64) // from string to float64

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil //returning no errors
}
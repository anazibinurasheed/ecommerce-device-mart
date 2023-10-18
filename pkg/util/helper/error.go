package helper

import "fmt"

// IsErr return prints the error and true if it is an error.
func IsErr(err error) bool {
	if err != nil {
		Logger("has error:", err)

	}
	return err != nil
}

// Logger can print any amount of data with any type.
func Logger(any ...any) {
	fmt.Println("\n" + fmt.Sprint(any...) + "\n")
}

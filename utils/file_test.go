package utils

import "testing"
import "fmt"

func TestReadFile(t *testing.T) {
	t.Run("file", func(t *testing.T) {
		got, err := ReadFile("../data/location.json")
		fmt.Println("ReadFile() error = %v, want = %v", err, string(got))
	})
}

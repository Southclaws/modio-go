package modio_test

import (
	"testing"

	"fmt"

	modio "github.com/Southclaws/modio-go"
)

func TestSimpleFile(t *testing.T) {
	fmt.Println("test start")
	_, err := modio.Open("tests/file1.dat")
	if err != nil {
		t.Error(err)
	}
}

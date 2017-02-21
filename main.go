package modio

import (
	"fmt"
)

// Open opens a modio format file by name, returns handle.
func Open(file string) (*Session, error) {
	fmt.Println("creating")
	session := Session{Path: file}
	fmt.Println("interpret")
	session.interpret()
	fmt.Println("done")

	return &session, nil
}

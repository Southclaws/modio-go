package modio

import (
	"fmt"
	"os"
	"unsafe"
)

// Session represents an open session with a modio file.
type Session struct {
	Path    string
	Mode    int32
	Version int32
	Length  int32
	Tags    int32
	Data    []int32
	Tagdict map[string]Tag
	Handle  *os.File
}

func (s Session) interpret() error {
	fmt.Println("interpreting...")

	file, err := os.Open(s.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	header := make([]byte, 12)
	_, err = file.Read(header)
	if err != nil {
		return err
	}

	s.Version = bytesToInt32(header[0:3])
	s.Length = bytesToInt32(header[4:7])
	s.Tags = bytesToInt32(header[8:11])

	type tagList struct {
		tagname int32
		physpos int32
	}
	tmp := make([]byte, 8)
	tl := make([]tagList, s.Tags)

	for i := int32(0); i < s.Tags; i++ {
		bytes, err := file.Read(tmp)
		if err != nil {
			return err
		}
		if bytes != 8 {
			return fmt.Errorf("while looping over tags, bytes != 2")
		}
		tl[i].tagname = bytesToInt32(tmp[0:3])
		tl[i].physpos = bytesToInt32(tmp[4:7])
	}

	for i := int32(0); i < s.Tags; i++ {
	}

	return nil
}

func bytesToInt32(b []byte) int32 {
	return *(*int32)(unsafe.Pointer(&b[0]))
}

package addr

import (
	"fmt"
	"strings"
)

type ResourceAddr struct {
	Type string
	Name string
}

func ParseAddress(addr string) (*ResourceAddr, error) {
	segs := strings.Split(addr, ".")
	if len(segs) != 2 {
		return nil, fmt.Errorf("invalid resource address found: %s", addr)
	}
	return &ResourceAddr{Type: segs[0], Name: segs[1]}, nil
}

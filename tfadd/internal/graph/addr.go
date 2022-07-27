package graph

import "strings"

type PropertyAddr []string

func MustParseAddr(addr string) PropertyAddr {
	return strings.Split(addr, ".")
}

func (addr PropertyAddr) Belongs(oaddr PropertyAddr) bool {
	if len(oaddr) > len(addr) {
		return false
	}
	for i := range oaddr {
		if oaddr[i] != addr[i] {
			return false
		}
	}
	return true
}

func (addr PropertyAddr) Equals(oaddr PropertyAddr) bool {
	if len(oaddr) != len(addr) {
		return false
	}
	return addr.Belongs(oaddr)
}

func (addr PropertyAddr) String() string {
	return strings.Join(addr, ".")
}

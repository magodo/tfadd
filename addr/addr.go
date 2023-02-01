package addr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ModuleStep struct {
	Name string

	// At most one of below is not nil
	Key   *string
	Index *int
}

func (step ModuleStep) String() string {
	out := "module." + step.Name
	switch {
	case step.Key != nil:
		out += `["` + *step.Key + `"]`
	case step.Index != nil:
		out += `[` + strconv.Itoa(*step.Index) + `]`
	}
	return out
}

type ModuleAddr []ModuleStep

func (addr ModuleAddr) String() string {
	var segs []string
	for _, ms := range addr {
		segs = append(segs, ms.String())
	}
	if len(segs) == 0 {
		return ""
	}
	return strings.Join(segs, ".")
}

func ParseModuleAddr(addr string) (ModuleAddr, error) {
	segs := strings.Split(addr, ".")
	if len(segs)%2 != 0 {
		return nil, fmt.Errorf("invalid module address")
	}

	var maddr ModuleAddr
	p := regexp.MustCompile(`^([^\[\]]+)(\[(.+)\])?$`)
	for i := 0; i < len(segs); i += 2 {
		if segs[i] != "module" {
			return nil, fmt.Errorf(`expect "module", got %q`, segs[i])
		}
		moduleSeg := segs[i+1]
		matches := p.FindStringSubmatch(moduleSeg)
		if len(matches) == 0 {
			return nil, fmt.Errorf("invalid module segment: %s", moduleSeg)
		}
		ms := ModuleStep{
			Name: matches[1],
		}
		if matches[3] == "" {
			if matches[2] != "" {
				return nil, fmt.Errorf("invalid module segment: %s", moduleSeg)
			}
		} else {
			idxLit := matches[3]
			if strings.HasPrefix(idxLit, `"`) && strings.HasSuffix(idxLit, `"`) {
				key, err := strconv.Unquote(idxLit)
				if err != nil {
					return nil, fmt.Errorf("unquoting module key %s: %v", idxLit, err)
				}
				ms.Key = &key
			} else {
				idx, err := strconv.Atoi(idxLit)
				if err != nil {
					return nil, fmt.Errorf("converting module index to number %s: %v", idxLit, err)
				}
				ms.Index = &idx
			}
		}
		maddr = append(maddr, ms)
	}
	return maddr, nil
}

type ResourceAddr struct {
	ModuleAddr ModuleAddr
	Type       string
	Name       string
}

func (addr ResourceAddr) String() string {
	raddr := addr.Type + "." + addr.Name
	if moduleAddr := addr.ModuleAddr.String(); moduleAddr != "" {
		raddr = moduleAddr + "." + raddr
	}
	return raddr
}

func ParseResourceAddr(addr string) (*ResourceAddr, error) {
	segs := strings.Split(addr, ".")

	if len(segs)%2 != 0 {
		return nil, fmt.Errorf("invalid resource address")
	}

	raddr := &ResourceAddr{
		Type: segs[len(segs)-2],
		Name: segs[len(segs)-1],
	}

	if len(segs) == 2 {
		return raddr, nil
	}

	maddr, err := ParseModuleAddr(strings.Join(segs[:len(segs)-2], "."))
	if err != nil {
		return nil, err
	}

	raddr.ModuleAddr = maddr
	return raddr, nil

}

func MustParseResourceAddr(addr string) *ResourceAddr {
	out, err := ParseResourceAddr(addr)
	if err != nil {
		panic(err)
	}
	return out
}

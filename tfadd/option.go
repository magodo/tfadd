package tfadd

import (
	"github.com/magodo/tfadd/addr"
)

type StateOption interface {
	configureState(*stateConfig)
}

var _ StateOption = fullOption(true)

type fullOption bool

func Full(b bool) fullOption {
	return fullOption(b)
}

func (opt fullOption) configureState(cfg *stateConfig) {
	cfg.full = bool(opt)
}

var _ StateOption = targetOption{}

type targetOption addr.ResourceAddr

func Target(raddr string) targetOption {
	// Validation for the resource address is guaranteed in flag parsing.
	addr, _ := addr.ParseResourceAddr(raddr)
	return targetOption(*addr)
}

func (opt targetOption) configureState(cfg *stateConfig) {
	target := addr.ResourceAddr(opt)
	cfg.target = &target
}

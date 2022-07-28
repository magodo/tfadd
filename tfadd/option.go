package tfadd

import (
	"github.com/magodo/tfadd/addr"
)

type StateOption interface {
	configureState(*stateConfig)
}

var _ StateOption = ignoreAttrConstraintsOption(true)

type ignoreAttrConstraintsOption bool

func IgnoreAttrConstraints(b bool) ignoreAttrConstraintsOption {
	return ignoreAttrConstraintsOption(b)
}

func (opt ignoreAttrConstraintsOption) configureState(cfg *stateConfig) {
	cfg.tuneOption.IgnoreAttrConstraints = bool(opt)
}

var _ StateOption = targetOption{}

type targetOption addr.ResourceAddr

func Target(raddr string) targetOption {
	// Validation for the resource address is guaranteed in flag parsing.
	addr, _ := addr.ParseAddress(raddr)
	return targetOption(*addr)
}

func (opt targetOption) configureState(cfg *stateConfig) {
	raddr := addr.ResourceAddr(opt)
	if !cfg.targetMap[raddr] {
		cfg.targets = append(cfg.targets, raddr)
		cfg.targetMap[raddr] = true
	}
}

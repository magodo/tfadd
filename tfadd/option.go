package tfadd

import (
	"github.com/magodo/tfadd/addr"
)

type FailableOption interface {
	Error() error
}

type RunOption interface {
	configureRun(*runConfig)
}

var _ RunOption = fullOption(true)

type fullOption bool

func Full(b bool) fullOption {
	return fullOption(b)
}

func (opt fullOption) configureRun(cfg *runConfig) {
	cfg.full = bool(opt)
}

var _ RunOption = targetOption{}
var _ FailableOption = targetOption{}

type targetOption struct {
	err  error
	addr *addr.ResourceAddr
}

func Target(raddr string) targetOption {
	addr, err := addr.ParseAddress(raddr)
	return targetOption{
		err:  err,
		addr: addr,
	}
}

func (opt targetOption) configureRun(cfg *runConfig) {
	cfg.target = opt.addr
}

func (opt targetOption) Error() error {
	return opt.err
}

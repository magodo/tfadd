package tfadd

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

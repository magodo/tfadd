package tfadd

type Option struct {
	// Whether the generated config contains all the non-computed properties?
	// Set via Full option.
	full bool

	// Whether to mask the sensitive attributes in the generated config?
	// Set via MaskSensitive option.
	maskSensitive bool
}

type OptionSetter interface {
	configureState(*Option)
}

var _ OptionSetter = fullOption(true)

type fullOption bool

func Full(b bool) fullOption {
	return fullOption(b)
}

func (opt fullOption) configureState(cfg *Option) {
	cfg.full = bool(opt)
}

var _ OptionSetter = maskSensitiveOption(true)

type maskSensitiveOption bool

func MaskSenstitive(b bool) maskSensitiveOption {
	return maskSensitiveOption(b)
}

func (opt maskSensitiveOption) configureState(cfg *Option) {
	cfg.maskSensitive = bool(opt)
}

package tfadd

type Option struct {
	// Whether the generated config contains all the non-computed properties?
	// Equivalent to enabling keepOC, keepZero and keepDefault all at once.
	// Set via Full option.
	full bool

	// Whether to mask the sensitive attributes in the generated config?
	// Set via MaskSensitive option.
	maskSensitive bool

	// Whether to keep Optional+Computed (O+C) attributes/blocks rather than
	// trimming them. Set via KeepOC option.
	keepOC bool

	// Whether to keep Optional attributes whose value equals the type's
	// "zero" value (used when the schema does not define a default).
	// Set via KeepZero option.
	keepZero bool

	// Whether to keep Optional attributes whose value equals the
	// schema-defined default value. Set via KeepDefault option.
	keepDefault bool
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

var _ OptionSetter = keepOCOption(true)

type keepOCOption bool

// KeepOC keeps Optional+Computed (O+C) attributes/blocks rather than
// trimming them.
func KeepOC(b bool) keepOCOption {
	return keepOCOption(b)
}

func (opt keepOCOption) configureState(cfg *Option) {
	cfg.keepOC = bool(opt)
}

var _ OptionSetter = keepZeroOption(true)

type keepZeroOption bool

// KeepZero keeps Optional attributes whose value equals the type's
// "zero" value (used when the schema does not define a default).
func KeepZero(b bool) keepZeroOption {
	return keepZeroOption(b)
}

func (opt keepZeroOption) configureState(cfg *Option) {
	cfg.keepZero = bool(opt)
}

var _ OptionSetter = keepDefaultOption(true)

type keepDefaultOption bool

// KeepDefault keeps Optional attributes whose value equals the
// schema-defined default value.
func KeepDefault(b bool) keepDefaultOption {
	return keepDefaultOption(b)
}

func (opt keepDefaultOption) configureState(cfg *Option) {
	cfg.keepDefault = bool(opt)
}

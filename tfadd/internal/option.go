package internal

type Option struct {
	// Whether to mask the sensitive attributes in the generated config?
	// Set via MaskSensitive option.
	MaskSensitive bool
}

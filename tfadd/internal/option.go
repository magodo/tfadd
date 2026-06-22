package internal

type Option struct {
	// Whether to mask the sensitive attributes in the generated config?
	// Set via MaskSensitive option.
	MaskSensitive bool
}

type TuneOption struct {
	// Whether to remove O+C attributes/blocks, as long as it doesn't violate the cross property constraint?
	RemoveOC bool

	// The O+C attributes/blocks to keep as otherwise it is deemed to be removed.
	// The key is the string representation of the attribute's/block's address.
	// For attribute, the address is separated by ".".
	// For block, the address is separated by ".0.".
	OCToKeep map[string]bool

	// Whether to remove optional attributes whose value equals the type's zero value (used when the schema does not define a default).
	// A null attribute value also counts as zero and is removed when this is enabled.
	RemoveOZeroAttribute bool

	// Whether to remove optional attributes whose value equals the schema-defined default value.
	RemoveODefaultAttribute bool
}

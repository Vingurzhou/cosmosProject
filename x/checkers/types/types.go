package types

import sdkerrors "cosmossdk.io/errors"

// x/checkers module sentinel errors
var (
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1103, "black address is invalid: %s")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
)

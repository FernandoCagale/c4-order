package event

import "github.com/google/wire"

var Set = wire.NewSet(NewOrder)
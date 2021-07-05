package lib

import (
	"github.com/google/uuid"
)

// BinaryUUID -> binary uuid wrapper over uuid.UUID
type BinaryUUID uuid.UUID

// ParseUUID -> parses string uuid to binary uuid
func ParseUUID(id string) BinaryUUID {
	return BinaryUUID(uuid.MustParse(id))
}

func (b BinaryUUID) String() string {
	return uuid.UUID(b).String()
}

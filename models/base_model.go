package models

import "clean_gin_api/lib"

type Base struct {
	ID lib.BinaryUUID `json:"id"`
}

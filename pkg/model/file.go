package model

import (
	"github.com/google/uuid"
)

// File 文件信息
type File struct {
	Model

	OID  uint64
	UUID uuid.UUID
}


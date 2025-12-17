package googlebooks

import (
	"github.com/RonitSamaddar/BookYu/internal/types"
)

type GoogleBooksGetVolumeResponse struct {
	Kind       string                      `json:"kind"`
	TotalItems int                         `json:"totalItems"`
	Items      []GoogleBooksGetVolumesItem `json:"items"`
}

type GoogleBooksGetVolumesItem struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo types.Book `json:"volumeInfo"`
}

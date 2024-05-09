package model

import (
	"movie-service/metadata/pkg/model"
)

type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}

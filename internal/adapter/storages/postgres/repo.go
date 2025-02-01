package postgres

import (
	"errors"

	"url/internal/config"
)

var NotFound = errors.New("long url not found")

type Repo struct {
	cluster *config.Cluster
}

func NewRepoDb(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}

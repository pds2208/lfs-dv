package db

import (
	"fmt"
	"github.com/ONSDigital/lfs-imports/config"
	"github.com/rs/zerolog/log"
	"lfs-dv/db/postgres"
	"lfs-dv/types"
	"sync"
)

type DVPersistence interface {
	Connect() error
	Close()

	GetDVConfiguration() ([]types.DV, error)
}

var cachedConnection DVPersistence
var connectionMux = &sync.Mutex{}

func GetDVPersistenceImpl() (DVPersistence, error) {
	connectionMux.Lock()
	defer connectionMux.Unlock()

	if cachedConnection != nil {
		log.Info().
			Str("databaseName", config.Config.Database.Database).
			Msg("Returning cached database connection")
		return cachedConnection, nil
	}

	cachedConnection = &postgres.Postgres{}

	if err := cachedConnection.Connect(); err != nil {
		log.Info().
			Err(err).
			Str("databaseName", config.Config.Database.Database).
			Msg("Cannot connect to database")
		cachedConnection = nil
		return nil, fmt.Errorf("cannot connect to database")
	}

	return cachedConnection, nil
}

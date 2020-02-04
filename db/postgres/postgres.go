package postgres

import (
	"github.com/ONSDigital/lfs-imports/db/postgres"
	"github.com/rs/zerolog/log"
	"lfs-dv/config"
	"time"
	"upper.io/db.v3/postgresql"
)

type Postgres struct {
	db postgres.PostgresConnection
}

func (s *Postgres) Connect() error {

	var settings = postgresql.ConnectionURL{
		Database: config.DatabaseConfiguration.Database.Database,
		Host:     config.DatabaseConfiguration.Database.Server,
		User:     config.DatabaseConfiguration.Database.User,
		Password: config.DatabaseConfiguration.Database.Password,
	}

	log.Debug().
		Str("databaseName", config.DatabaseConfiguration.Database.Database).
		Msg("Connecting to database")

	sess, err := postgresql.Open(settings)

	if err != nil {
		log.Error().
			Err(err).
			Str("databaseName", config.DatabaseConfiguration.Database.Database).
			Msg("Cannot connect to database")
		return err
	}

	log.Debug().
		Str("databaseName", config.DatabaseConfiguration.Database.Database).
		Msg("Connected to database")

	if config.DatabaseConfiguration.Database.Verbose {
		sess.SetLogging(true)
	}

	s.db.DB = sess

	poolSize := config.DatabaseConfiguration.Database.ConnectionPool.MaxPoolSize
	maxIdle := config.DatabaseConfiguration.Database.ConnectionPool.MaxIdleConnections
	maxLifetime := config.DatabaseConfiguration.Database.ConnectionPool.MaxLifetimeSeconds

	if maxLifetime > 0 {
		maxLifetime = maxLifetime * time.Second
		sess.SetConnMaxLifetime(maxLifetime)
	}

	log.Debug().
		Int("MaxPoolSize", poolSize).
		Int("MaxIdleConnections", maxIdle).
		Dur("MaxLifetime", maxLifetime*time.Second).
		Msg("Connection Attributes")

	sess.SetMaxOpenConns(poolSize)
	sess.SetMaxIdleConns(maxIdle)

	return nil
}

func (s Postgres) Close() {
	if s.db.DB != nil {
		_ = s.db.DB.Close()
	}
}

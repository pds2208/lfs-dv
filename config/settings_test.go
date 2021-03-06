package config

import (
	conf "github.com/ONSDigital/lfs-imports/config"
	"testing"
)

func TestConfig(t *testing.T) {
	server := DatabaseConfiguration.Database.Server
	if server != "localhost" {
		t.Errorf("server = %s; want localhost", server)
	} else {
		t.Logf("Server %s\n", server)
	}

	user := conf.Config.Database.User
	if user != "lfs" {
		t.Errorf("user = %s; want lfs", user)
	} else {
		t.Logf("user %s\n", user)
	}

	password := DatabaseConfiguration.Database.Password
	if password != "lfs" {
		t.Errorf("password = %s; want lfs", password)
	} else {
		t.Logf("password %s\n", password)
	}

	databaseName := DatabaseConfiguration.Database.Database
	if databaseName != "lfs" {
		t.Errorf("database name = %s; want lfs", databaseName)
	} else {
		t.Logf("database name %s\n", databaseName)
	}

	maxPoolsize := DatabaseConfiguration.Database.ConnectionPool.MaxPoolSize
	if maxPoolsize != 10 {
		t.Errorf("maxPoolsize = %d; want 10", maxPoolsize)
	} else {
		t.Logf("maxPoolsize %d\n", maxPoolsize)
	}

}

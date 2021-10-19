package pgdb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
)

type Config struct {
	ConnectionName string
	Hostname       string
	Username       string
	Password       string
	Database       string
	Port           string
	Debug          bool
}

// Init ...
func Init(cfg Config) (*pg.DB, error) {
	if len(cfg.Hostname) == 0 && len(cfg.ConnectionName) == 0 {
		return nil, errors.New("you need to set 'hostname' or 'connection name' for connection to work")
	}

	if len(cfg.Username) == 0 {
		return nil, errors.New("you need to user 'username' for connection to work")
	}

	if len(cfg.Database) == 0 {
		return nil, errors.New("you need to user 'database' name for connection to work")
	}

	if len(cfg.Port) == 0 {
		cfg.Port = "5432"
	}

	// Database connection options
	opts := pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Hostname, cfg.Port),
		User:     cfg.Username,
		Database: cfg.Database,
		OnConnect: func(ctx context.Context, conn *pg.Conn) error {
			log.Printf("DB %s connected", cfg.Database)

			return nil
		},
	}

	if len(cfg.Password) > 0 {
		opts.Password = cfg.Password
	}

	// Enable GCP proxy if running on GCP
	if len(cfg.ConnectionName) > 0 {
		opts.Addr = ""
		opts.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return proxy.Dial(cfg.ConnectionName)
		}
	}

	db := pg.Connect(&opts)

	if cfg.Debug {
		db.AddQueryHook(pgdebug.DebugHook{Verbose: true})
	}

	return db, nil
}

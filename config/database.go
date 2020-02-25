package config

import "github.com/koushamad/goro-core/app/helper"

var (
	// mysql | postgres | sqlite3
	Database = map[string]string{
		"driver": helper.Env("DATABASE_DRIVER", "sqlite3"),
		"host": helper.Env("DATABASE_HOST",""),
		"port": helper.Env("DATABASE_PORT",""),
		"database": helper.Env("DATABASE_NAME","goro.db"),
		"user": helper.Env("DATABASE_USER",""),
		"pass": helper.Env("DATABASE_PASS",""),
	}
)


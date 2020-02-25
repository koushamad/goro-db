package configProvider

import (
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-db/config"
)

func Load() {
	conf.Add("DB", config.Database)
}

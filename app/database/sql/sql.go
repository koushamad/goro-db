package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-db/app/database"
	"github.com/koushamad/goro-db/app/database/sql/mysql"
	"github.com/koushamad/goro-db/app/database/sql/sqlite"
)

var (
	DB database.Database
)

type Sql struct {}

func Boot() {
	switch conf.Get("db.driver") {
	case "mysql":
		DB = mysql.Boot()
		break
	default:
		DB = sqlite.Boot()
		break
	}
}

func Init() *gorm.DB {
	return DB.Init()
}

func Kill()  {
	DB.Kill()
}

func Migrate(tables []interface{})  {
	DB.Migrate(tables)
}
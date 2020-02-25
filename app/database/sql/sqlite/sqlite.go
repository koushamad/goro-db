package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-core/app/helper"
	"sync"
)

var (
	once sync.Once
	Sqlite *sqlite
)

type sqlite struct {
	DB *gorm.DB
}

func Boot() *sqlite {
	once.Do(func() {
		db, err := gorm.Open(conf.Get("db.driver"), helper.SqlitePatch())
		throw.Fatal("failed to connect database", 115, err)
		Sqlite = &sqlite{DB:db}
	})

	return Sqlite
}

func (s *sqlite) Init() *gorm.DB {
	return Sqlite.DB
}

func (s *sqlite) Kill() {
	defer s.DB.Close()
}

func (s *sqlite) Migrate(tables []interface{})  {
	for _,t := range tables{
		s.DB.AutoMigrate(t)
	}
}
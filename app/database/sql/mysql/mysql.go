package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/exception/throw"
	"sync"
)

var (
	once  sync.Once
	Mysql *mysql
)

type mysql struct {
	DB *gorm.DB
}

func Boot() *mysql {
	once.Do(func() {
		db, err := gorm.Open(conf.Get("db.driver"), getConnection)
		throw.Fatal("failed to connect database", 114, err)
		Mysql = &mysql{DB: db}
	})

	return Mysql
}

func (s *mysql) Init() *gorm.DB {
	return Mysql.DB
}

func (s *mysql) Kill() {
	defer s.DB.Close()
}

func getConnection() string {
	return conf.Get("db.user") +
		":" + conf.Get("db.pass") +
		"@" + conf.Get("db.host") + conf.Get("db.port") +
		"/" + conf.Get("db.database")
}

func (s *mysql) Migrate(tables []interface{}) {
	for _, t := range tables {
		s.DB.AutoMigrate(t)
	}
}

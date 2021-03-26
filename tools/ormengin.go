package tools

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
}

func OrmEngin() (*xorm.Engine, error) {
	dbname := ""
	engine, err := xorm.NewEngine("mysql", dbname)
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	return engine, nil
}

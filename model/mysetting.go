package model
import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
//	"time"
	"fmt"
)

type mysetting struct {
	Id int64
	KeyName string
	KeyValue string
}
var engine *xorm.Engine

func GetAllSetting() []mysetting{
	lst:=make([]mysetting,0)
	err:=engine.Find(&lst)
	if err!=nil{
		fmt.Println(err)
	}
	return lst
}
func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("xorm error:", err)
	}

	engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
	err = engine.Sync(new(mysetting))
}



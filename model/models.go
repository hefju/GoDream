package model
import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
		"time"
	"fmt"
)
var engine *xorm.Engine
func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("xorm error:", err)
	}

	engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
	err = engine.Sync(new(Mysetting),new(UpdateLog))//
	if err != nil {
		fmt.Println("xorm error333:", err)
	}
}


type Mysetting struct {
	Id int64
	KeyName string
	KeyValue string
}

func GetAllSetting() []Mysetting{
	lst:=make([]Mysetting,0)
	err:=engine.Find(&lst)
	if err!=nil{
		fmt.Println(err)
	}
	return lst
}
func (x *Mysetting)Save()int64{
	var count int64
	var err error
	if x.Id>0{
		count,err=engine.Id(x.Id).Update(x)
	}else{
		count,err=engine.Insert(x)
	}
	if err!=nil{
		fmt.Println("xorm save error:",err)
	}
	return count
}
func (x *Mysetting)Delete()int64{
	count,err:=engine.Id(x.Id).Delete(x)
	if err!=nil{
		fmt.Println("xorm save error:",err)
	}
	return count
}

type UpdateLog struct {
	Id int64
	Title string
	Content string
	Created time.Time `xorm:"created"`
}
func GetAllUpdateLog()[]UpdateLog{
	lst:=make([]UpdateLog,0)
	err:=engine.OrderBy("Title desc").Find(&lst)
	if err!=nil{
		fmt.Println(err)
	}
	return lst
}

func (x *UpdateLog)Save()int64{
	var count int64
	var err error
	if x.Id>0{
		count,err=engine.Id(x.Id).Update(x)
	}else{
		count,err=engine.Insert(x)
	}
	if err!=nil{
		fmt.Println("xorm save error:",err)
	}
	return count
}


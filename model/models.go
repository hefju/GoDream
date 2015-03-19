package model

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("xorm error:", err)
	}

	//engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
	err = engine.Sync(new(Mysetting), new(UpdateLog),new(ReportMsg),new(MyTask)) //
	if err != nil {
		fmt.Println("xorm error333:", err)
	}
}

type ReportMsg struct {
    Id int64
    Title string
    Content string
    CreatedAt time.Time `xorm:"created"`
}



//我的设置(数据库)
type Mysetting struct {
	Id       int64
	KeyName  string
	KeyValue string
}

func (x *Mysetting) GetValue(key string)(string, error){
	set:=new(Mysetting)
	has,err:=engine.Where("KeyName=?",key).Get(set)
	if has{
		return set.KeyValue, nil
	}else{
		return "",err
	}
}
func GetAllSetting() []Mysetting {
	lst := make([]Mysetting, 0)
	err := engine.Find(&lst)
	if err != nil {
		fmt.Println(err)
	}
	return lst
}
func (x *Mysetting) Save() int64 {
	var count int64
	var err error
	if x.Id > 0 {
		count, err = engine.Id(x.Id).Update(x)
	} else {
		count, err = engine.Insert(x)
	}
	if err != nil {
		fmt.Println("xorm save error:", err)
	}
	return count
}
func (x *Mysetting) Delete() int64 {
	count, err := engine.Id(x.Id).Delete(x)
	if err != nil {
		fmt.Println("xorm save error:", err)
	}
	return count
}

//每日更新的内容
type UpdateLog struct {
	Id      int64
	Title   string
	Content string
	Created time.Time `xorm:"created"`
}

func GetAllUpdateLog() []UpdateLog {
	lst := make([]UpdateLog, 0)
	err := engine.OrderBy("Title desc").Find(&lst)
	if err != nil {
		fmt.Println(err)
	}
	return lst
}

func (x *UpdateLog) Save() int64 {
	var count int64
	var err error
	if x.Id > 0 {
		count, err = engine.Id(x.Id).Update(x)
	} else {
		count, err = engine.Insert(x)
	}
	if err != nil {
		fmt.Println("xorm save error:", err)
	}
	return count
}

//我的任务
type MyTask struct {
	Id       int64
	Title    string
	Content  string
	Finish   bool
	Created  time.Time `xorm:"created"`
	FinishAt time.Time
}
func GetTask(date string) []MyTask {
    lst := make([]MyTask, 0)
    err := engine.Find(&lst)
    if err != nil {
        fmt.Println(err)
    }
    return lst
}
func GetTaskByID(id int64) *MyTask {
    t:=&MyTask{Id:id}
    _, err := engine.Get(t)
    if err != nil {
        fmt.Println(err)
    }
    return t
}

func UpdateTask(task *MyTask )(int64, error) {
    affect,err:= engine.Id(task.Id).Update(task)
    return affect,err
}
func Insert(obj interface{}) (int64, error)  {
    affect,err:=engine.InsertOne(obj)
    return affect,err
}
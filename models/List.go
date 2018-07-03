package models

import (
	"time"
	"github.com/astaxie/beedb"
	"database/sql"
	_"github.com/ziutek/mymysql/godrv"
)

type Lists struct {
	Id		int `PK`
	Title		string
	Description	string
	Created		time.Time
}

func GetLink() beedb.Model  {
	db, err := sql.Open("mymysql","belajar-golang/test/root")
	if(err != nil){
		panic(err)
	}

	orm := beedb.New(db)
	return orm;
}

func GetAll() (lists []Lists)  {
	db := GetLink();
	db.FindAll(&lists);
	return
}

func GetList(id int) (list Lists)  {
	db := GetLink()
	db.Where("id=?",id).Find(&list)
	return
}

func SaveList(list Lists) (l Lists)  {
	db := GetLink()
	db.Save(&list)
	return l
}

func DelList(id int)  {
	db := GetLink()
	db.SetTable("lists").Where("id=?",id).DeleteRow();

	return
}
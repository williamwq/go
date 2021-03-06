package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/williamwq/go/src/api/config"
)

var Eloquent *gorm.DB

func init() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	var err error
	username := viper.GetString("common.db.username")
	password := viper.GetString("common.db.password")
	addr := viper.GetString("common.db.addr")
	port := viper.GetString("common.db.port")
	dbname := viper.GetString("common.db.dbname")
	orm := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", username, password, addr, port, dbname)
	Eloquent, err = gorm.Open("mysql", orm)
	Eloquent.SingularTable(true)
	//fmt.Println(orm)
	//Eloquent, err = gorm.Open("mysql", fmt.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms",username,password,addr,port,dbname))

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}

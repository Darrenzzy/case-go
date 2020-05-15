package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"order/internal/model"
)

// MysqlCfg mysql config
type MysqlCfg struct {
	Host struct {
		Read  string `toml:"read" json:"read"`
		Write string `toml:"write" json:"write"`
	} `toml:"host" json:"Host"`

	Port    int    `toml:"port" json:"Port"`
	User    string `toml:"user" json:"User"`
	Psw     string `toml:"psw" json:"Psw"`
	DbName  string `toml:"dbname" json:"DbName"`
	LogMode bool   `toml:"Logmode" json:"Logmode"`
}

// Mysql mysql config
type Mysql struct {
	Read    string `toml:"read" json:"read"`
	Write   string `toml:"write" json:"write"`
	Port    int    `toml:"port" json:"Port"`
	User    string `toml:"user" json:"User"`
	Psw     string `toml:"psw" json:"Psw"`
	DbName  string `toml:"dbname" json:"DbName"`
	Charset string `toml:"charset" json:"charset"`
}

func NewDbOrm() (gdb *gorm.DB, cf func(), err error) {
	var (
		cfg Mysql
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	gdb, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Psw, cfg.Read, cfg.Port, cfg.DbName))
	if err != nil {
		return nil, nil, err
	}
	cf = func() {
		gdb.Close()
	}
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	err = d.db.Table("articles").Where("id=?", id).First(&art).Error

	fmt.Printf("%+v,%+v ***********", art, err)
	return
}

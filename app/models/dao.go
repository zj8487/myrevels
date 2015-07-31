package models

import (
	"gopkg.in/mgo.v2"
)

const (
	URL               = "192.168.139.132:27018"
	DbName            = "myrevelsdb"
	BlogCollection    = "blogs"
	CommentCollection = "gb_comments"
	MessageCollection = "gb_messages"
	HistoryCollection = "gb_histoty"
	EmailCollection   = "gb_emails"
	BaseYear          = 2015
)

//数据访问对象
type Dao struct {
	session *mgo.Session
}

func NewDao() (*Dao, error) {
	session, err := mgo.Dial(URL)
	if err != nil {
		return nil, err
	}
	return &Dao{session}, nil
}

//关闭
func (d *Dao) Close() {
	d.session.Close()
}

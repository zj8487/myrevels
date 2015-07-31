package models

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//留言对象
type Message struct {
	Email   string
	QQ      string
	Url     string
	CDate   time.Time
	Content string
}

//校验方法
func (message *Message) Validate(v *revel.Validation) {
	v.Check(message.Email,
		revel.Required{},
		revel.MaxSize{50})
	v.Email(message.Email)

	v.Check(message.QQ,
		revel.MaxSize{20})

	v.Check(message.Url, revel.MaxSize{200})
	v.Check(message.Content,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{1000})
}

//插入消息
func (dao *Dao) InsertMessage(message *Message) error {
	messCollection := dao.session.DB(DbName).C(MessageCollection)

	//set the time
	err := messCollection.Insert(message)
	if err != nil {
		revel.WARN.Printf("Unable to save Message: %v error %v", message, err)
	}
	return err
}

//获取所有的留言
func (dao *Dao) FindAllMessage() []Message {
	messCollection := dao.session.DB(DbName).C(MessageCollection)
	mess := []Message{}
	query := messCollection.Find(bson.M{}).Sort("-cdate")
	query.All(&mess)
	return mess
}

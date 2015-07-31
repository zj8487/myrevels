package models

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//评论
type Comment struct {
	BlogId  bson.ObjectId
	Email   string
	CDate   time.Time
	Content string
}

//评论校验
func (comment *Comment) Validate(v *revel.Validation) {
	v.Check(comment.Email,
		revel.Required{},
		revel.MaxSize{50})
	v.Email(comment.Email)

	v.Check(comment.Content,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{1000})
}

//插入评论
func (dao *Dao) InsertComment(comment *Comment) error {
	commentCollection := dao.session.DB(DbName).C(CommentCollection)
	comment.CDate = time.Now()
	err := commentCollection.Insert(comment)
	if err != nil {
		revel.WARN.Printf("Unable to save Comment: %v error %v", comment, err)
	}
	return err
}

//获取博客的所有评论
func (dao *Dao) FindCommentsByBlogId(id bson.ObjectId) []Comment {
	commentCollection := dao.session.DB(DbName).C(CommentCollection)
	comms := []Comment{}
	query := commentCollection.Find(bson.M{"blogid": id}).Sort("CDate")
	query.All(&comms)
	return comms
}

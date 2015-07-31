package models

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//博文对象
type Blog struct {
	Id           bson.ObjectId
	Email        string
	CDate        time.Time
	Title        string
	ShortTitle   string
	Subject      string
	ShortSubject string
	CommentCnt   int
	ReadCnt      int
	Year         int
}

//校验功能
func (blog *Blog) Validate(v *revel.Validation) {

	//对标题的校验
	v.Check(blog.Title,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{200})

	//对电子邮件的校验
	v.Check(blog.Email,
		revel.Required{},
		revel.MaxSize{50})
	v.Email(blog.Email)

	//对主题的校验
	v.Check(blog.Subject,
		revel.Required{},
		revel.MinSize{1})
}

//短标题
func (blog *Blog) GetShortTitle() string {
	if len(blog.Title) > 50 {
		return blog.Title[:50]
	}
	return blog.Title
}

func (blog *Blog) GetShortContent() string {
	if len(blog.Subject) > 200 {
		return blog.Subject[:200]
	}
	return blog.Subject
}

//创建博文
func (dao *Dao) CreateBlog(blog *Blog) error {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)

	//设置相关动态属性
	blog.Id = bson.NewObjectId()
	blog.CDate = time.Now()
	blog.Year = blog.CDate.Year()
	_, err := blogCollection.Upsert(bson.M{"_id": blog.Id}, blog)
	if err != nil {
		revel.WARN.Printf("Unable to save blog : %v error %v ", blog, err)
	}
	return err
}

//获取所有博文
func (dao *Dao) FindBlogs() []Blog {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	blogs := []Blog{}
	query := blogCollection.Find(bson.M{}).Sort("-cdate").Limit(50) //列出最新的50条
	query.All(&blogs)
	return blogs
}

//根号主键获取博文
func (dao *Dao) FindBlogById(id string) *Blog {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	blog := new(Blog)
	query := blogCollection.Find(bson.M{"id": bson.ObjectIdHex(id)})
	query.One(blog)
	return blog
}

//更新博文
func (dao *Dao) UpdateBlogById(id string, blog *Blog) {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	err := blogCollection.Update(bson.M{"id": bson.ObjectIdHex(id)}, blog)
	if err != nil {
		revel.WARN.Printf("Unable to update blog: %v  error %v", blog, err)
	}
}

func (dao *Dao) FindBlogsByYear(year int) []Blog {
	blogCollection := dao.session.DB(DbName).C(BlogCollection)
	blogs := []Blog{}
	query := blogCollection.Find(bson.M{"year": year}).Sort("-cdate")
	query.All(&blogs)
	return blogs
}

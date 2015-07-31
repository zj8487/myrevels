package controllers

import (
	"github.com/revel/revel"
	"myrevels/app/models"
)

type App struct {
	*revel.Controller
}

const (
	ACTIVE = "active"
)

//主页
func (c App) Index() revel.Result {
	home := ACTIVE

	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}

	defer dao.Close()
	blogs := dao.FindBlogs()

	return c.Render(home, blogs)
}

//投稿页面
func (c App) WBlog() revel.Result {
	write := ACTIVE
	return c.Render(write)
}

//博客具体信息
func (c App) BlogInfor(id string, rcnt int) revel.Result {

	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}

	defer dao.Close()
	//根据ID查找博文
	blog := dao.FindBlogById(id)
	if blog.ReadCnt == rcnt {
		blog.ReadCnt = rcnt + 1
		dao.UpdateBlogById(id, blog)
	}

	comments := dao.FindCommentsByBlogId(blog.Id)
	if len(comments) == 0 && blog.CommentCnt != 0 {
		blog.CommentCnt = 0
		dao.UpdateBlogById(id, blog)
	} else if len(comments) != blog.CommentCnt {
		blog.CommentCnt = len(comments)
		dao.UpdateBlogById(id, blog)
	}
	return c.Render(blog, rcnt, comments)
}

//留言吧

func (c App) Message() revel.Result {
	messagefun := ACTIVE

	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	messages := dao.FindAllMessage()
	return c.Render(messagefun, messages)
}

//归档功能
func (c App) History() revel.Result {
	historyfun := ACTIVE

	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	dao.CreateAllHistory()
	historys := dao.FindHistory()
	for i, _ := range historys {
		historys[i].Blogs = dao.FindBlogsByYear(historys[i].Year)
	}

	return c.Render(historyfun, historys)
}

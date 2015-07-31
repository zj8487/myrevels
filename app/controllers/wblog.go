package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myrevels/app/models"
	"strings"
)

type WBlog struct {
	App
}

//提交博文
func (c WBlog) Putup(blog *models.Blog) revel.Result {
	fmt.Println("提交博文.....")
	blog.Title = strings.TrimSpace(blog.Title)
	blog.Email = strings.TrimSpace(blog.Email)
	blog.Subject = strings.TrimSpace(blog.Subject)
	blog.Validate(c.Validation)

	//提交过程中出现异常的处理
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		fmt.Println(c.Validation)
		return c.Redirect(App.WBlog)
	}

	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}

	defer dao.Close()
	err = dao.CreateBlog(blog)
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}

	return c.Redirect(App.Index)
}

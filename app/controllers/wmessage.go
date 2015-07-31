package controllers

import (
	"github.com/revel/revel"
	"myrevels/app/models"
	"strings"
)

type WMessage struct {
	App
}

//提交留言
func (c WMessage) Putup(message *models.Message) revel.Result {
	message.Email = strings.TrimSpace(message.Email)
	message.Url = strings.TrimSpace(message.Url)
	message.Content = strings.TrimSpace(message.Content)
	message.QQ = strings.TrimSpace(message.QQ)

	message.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Errs:The email and the content should not be null,or the maxsize of email is 50.")
		return c.Redirect(App.Message)
	}

	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}

	defer dao.Close()
	err = dao.InsertMessage(message)
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	return c.Redirect(App.Message)
}

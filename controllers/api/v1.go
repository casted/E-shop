package api

import (
	"LeastMall/models"

	"github.com/astaxie/beego"
)

type V1Controller struct {
	beego.Controller
}

func (c *V1Controller) Get() {
	c.Ctx.WriteString("api v1")
}

func (c *V1Controller) Menu() {
	menu := []models.Menu{}
	models.DB.Find(&menu)
	c.Data["json"] = menu
	c.ServeJSON()
}

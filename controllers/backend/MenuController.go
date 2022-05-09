package backend

import (
	"fmt"
	"math"
	"strconv"

	"LeastMall/common"
	"LeastMall/models"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Get() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 3
	//查询数据
	menu := []models.Menu{}
	models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&menu)
	//查询menu表里面的数量
	var count int
	models.DB.Table("menu").Count(&count)
	// if len(menu) == 0 {
	// 	prvPage := page - 1
	// 	if prvPage == 0 {
	// 		prvPage = 1
	// 	}
	// 	c.Goto("/menu?page=" + strconv.Itoa(prvPage))
	// }
	c.Data["menuList"] = menu
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName = "backend/menu/index.html"
}

func (c *MenuController) Add() {
	c.TplName = "backend/menu/add.html"
}

func (c *MenuController) GoAdd() {
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")

	menu := models.Menu{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(common.GetUnix()),
	}

	err := models.DB.Create(&menu).Error
	if err != nil {
		c.Error("增加数据失败", "/menu/add")
	} else {
		c.Success("增加成功", "/menu")
	}
}

func (c *MenuController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/menu")
		return
	}
	menu := models.Menu{Id: id}
	models.DB.Find(&menu)
	c.Data["menu"] = menu
	c.Data["prevPage"] = c.Ctx.Request.Referer()
	c.TplName = "backend/menu/edit.html"
}

func (c *MenuController) GoEdit() {

	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/menu")
		return
	}
	title := c.GetString("title")
	link := c.GetString("link")
	position, _ := c.GetInt("position")
	isOpennew, _ := c.GetInt("is_opennew")
	relation := c.GetString("relation")
	sort, _ := c.GetInt("sort")
	status, _ := c.GetInt("status")
	prevPage := c.GetString("prevPage")
	fmt.Println("-----------------------", relation)
	//修改
	menu := models.Menu{Id: id}
	models.DB.Find(&menu)
	menu.Title = title
	menu.Link = link
	menu.Position = position
	menu.IsOpennew = isOpennew
	menu.Relation = relation
	menu.Sort = sort
	menu.Status = status

	err2 := models.DB.Save(&menu).Error
	if err2 != nil {
		c.Error("修改数据失败", "/menu/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", prevPage)
	}

}

func (c *MenuController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/menu")
		return
	}
	menu := models.Menu{Id: id}
	models.DB.Delete(&menu)

	c.Success("删除数据成功", c.Ctx.Request.Referer())
}

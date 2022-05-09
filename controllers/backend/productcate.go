package backend

import (
	"os"
	"strconv"

	"LeastMall/common"
	"LeastMall/models"

	"github.com/astaxie/beego"
)

type ProductCateController struct {
	BaseController
}

func (c *ProductCateController) Get() {
	productCate := []models.ProductCate{}
	models.DB.Preload("ProductCateItem").Where("pid=0").Find(&productCate)
	c.Data["productCateList"] = productCate
	c.TplName = "backend/productCate/index.html"
}

func (c *ProductCateController) Add() {
	productCate := []models.ProductCate{}
	models.DB.Where("pid=0").Find(&productCate)
	c.Data["productCateList"] = productCate
	c.TplName = "backend/productCate/add.html"
}

func (c *ProductCateController) GoAdd() {
	title := c.GetString("title")
	pid, err1 := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")

	if err1 != nil || err3 != nil {
		c.Error("传入参数类型不正确", "/productCate/add")
		return
	}
	if err2 != nil {
		c.Error("排序值必须是整数", "/productCate/add")
		return
	}
	uploadDir, _ := c.UploadImg("cate_img")
	productCate := models.ProductCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     uploadDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(common.GetUnix()),
	}
	err := models.DB.Create(&productCate).Error
	if err != nil {
		c.Error("增加失败", "/productCate/add")
		return
	}
	c.Success("增加成功", "/productCate")
}

func (c *ProductCateController) Edit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/productCate")
		return
	}
	productCate := models.ProductCate{Id: id}
	models.DB.Find(&productCate)
	c.Data["productCate"] = productCate
	productCateList := []models.ProductCate{}
	models.DB.Where("pid=0").Find(&productCateList)
	c.Data["productCateList"] = productCateList
	c.TplName = "backend/productCate/edit.html"
}

func (c *ProductCateController) GoEdit() {
	id, err := c.GetInt("id")
	title := c.GetString("title")
	pid, err1 := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err != nil || err1 != nil || err3 != nil {
		c.Error("传入参数类型不正确", "/productCate/edit")
		return
	}
	if err2 != nil {
		c.Error("排序值必须是整数", "/productCate/edit?id="+strconv.Itoa(id))
		return
	}
	uploadDir, _ := c.UploadImg("cate_img")
	productCate := models.ProductCate{Id: id}
	models.DB.Find(&productCate)
	productCate.Title = title
	productCate.Pid = pid
	productCate.Link = link
	productCate.Template = template
	productCate.SubTitle = subTitle
	productCate.Keywords = keywords
	productCate.Description = description
	productCate.Sort = sort
	productCate.Status = status
	if uploadDir != "" {
		productCate.CateImg = uploadDir
	}
	err5 := models.DB.Save(&productCate).Error
	if err5 != nil {
		c.Error("修改失败", "/productCate/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/productCate")
}

func (c *ProductCateController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/goodCate")
		return
	}
	productCate := models.ProductCate{Id: id}
	models.DB.Find(&productCate)
	address := "D:/gowork/src/gitee.com/shirdonl/LeastMall/" + productCate.CateImg
	test := os.Remove(address)
	if test != nil {
		beego.Error(test)
		c.Error("删除物理机上图片错误", "/goodCate")
		return
	}
	if productCate.Pid == 0 {
		productCate2 := []models.ProductCate{}
		models.DB.Where("pid=?", productCate.Id).Find(&productCate2)
		if len(productCate2) > 0 {
			c.Error("请删除当前顶级分类下面的商品！", "/productCate")
			return
		}
	}
	models.DB.Delete(&productCate)
	c.Success("删除成功", "/productCate")
}

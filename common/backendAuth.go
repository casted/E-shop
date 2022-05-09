package common

import (
	"net/url"
	"strings"

	"LeastMall/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//后台权限判断
func BackendAuth(ctx *context.Context) {
	pathname := ctx.Request.URL.String()
	userinfo, ok := ctx.Input.Session("userinfo").(models.Administrator)
	if !(ok && userinfo.Username != "") {
		if pathname != "/"+beego.AppConfig.String("adminPath")+"/login" &&
			pathname != "/"+beego.AppConfig.String("adminPath")+"/login/gologin" &&
			pathname != "/"+beego.AppConfig.String("adminPath")+"/login/verificode" {
			ctx.Redirect(302, "/"+beego.AppConfig.String("adminPath")+"/login")
		}
	} else {
		pathname = strings.Replace(pathname, "/"+beego.AppConfig.
			String("adminPath"), "", 1)
		urlPath, _ := url.Parse(pathname)
		if userinfo.IsSuper == 0 && !excludeAuthPath(string(urlPath.Path)) {
			roleId := userinfo.RoleId
			roleAuth := []models.RoleAuth{}
			models.DB.Where("role_id=?", roleId).Find(&roleAuth)
			roleAuthMap := make(map[int]int)
			for _, v := range roleAuth {
				roleAuthMap[v.AuthId] = v.AuthId
			}
			auth := models.Auth{}
			models.DB.Where("url=?", urlPath.Path).Find(&auth)
			if _, ok := roleAuthMap[auth.Id]; !ok {
				ctx.WriteString("没有权限")
				return
			}
		}
	}
}

//检验路径权限
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(beego.AppConfig.String("excludeAuthPath"), ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}

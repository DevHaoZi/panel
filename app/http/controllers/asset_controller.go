package controllers

import (
	"os"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"panel/app/internal"
	"panel/app/internal/services"
	"panel/pkg/tools"
)

type AssetController struct {
	setting internal.Setting
}

func NewAssetController() *AssetController {
	return &AssetController{
		setting: services.NewSettingImpl(),
	}
}

func (r *AssetController) Index(ctx http.Context) http.Response {
	entrance := facades.Config().GetString("http.entrance")
	if entrance == "/" {
		entrance = ""
	}

	// 自动纠正 URL 格式
	if ctx.Request().Path() == entrance && ctx.Request().Path() != "/" {
		return ctx.Response().Redirect(http.StatusMovedPermanently, ctx.Request().Path()+"/")
	}
	// 拒绝访问非入口文件
	if !strings.HasPrefix(ctx.Request().Path(), entrance) {
		return ctx.Response().Status(http.StatusNotFound).String(http.StatusText(http.StatusNotFound))
	}

	path := strings.TrimPrefix(ctx.Request().Path(), entrance)
	// 设置默认首页
	if path == "/" || path == "" {
		path = "/index.html"
	}

	if !tools.Exists("public" + path) {
		return ctx.Response().Status(http.StatusNotFound).String(http.StatusText(http.StatusNotFound))
	}

	file, err := os.Open("public" + path)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).String(http.StatusText(http.StatusInternalServerError))
	}

	stat, err := file.Stat()
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).String(http.StatusText(http.StatusInternalServerError))
	}

	if stat.IsDir() {
		return ctx.Response().Status(http.StatusForbidden).String(http.StatusText(http.StatusForbidden))
	}

	return ctx.Response().Header("Cache-Control", "no-cache").File("public" + path)
}

func (r *AssetController) Favicon(ctx http.Context) http.Response {
	return ctx.Response().File("public/favicon.png")
}

func (r *AssetController) Robots(ctx http.Context) http.Response {
	return ctx.Response().File("public/robots.txt")
}

func (r *AssetController) NotFound(ctx http.Context) http.Response {
	return ctx.Response().Status(http.StatusNotFound).String(http.StatusText(http.StatusNotFound))
}

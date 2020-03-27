package cate

import (
	"EK-Server/model"
	"EK-Server/util"

	"github.com/labstack/echo"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/cate"
	cate := g.Group(baseURL)
	cate.POST("/add", add)

	cate.GET("/list", list)
}

func list(c echo.Context) error {
	return util.JSONSuccess(c, model.DataBaselimit(10, 1, map[string]interface{}{"deleted_at": nil}, &[]model.GoodsCateList{}, "goods_cates", ""), "获取成功")
}

func add(c echo.Context) error {
	db := model.DB
	cate := &model.GoodsCate{}
	// 绑定json
	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	// if repeat := db.Find(&model.GoodsCate{Name: cate.Name}).RecordNotFound(); !repeat {
	// 	return util.JSONErr(c, nil, "已存在相同分类")
	// }
	db.NewRecord(cate)
	row := db.Create(cate)
	if row.RowsAffected >= 1 {
		return util.JSONSuccess(c, nil, "添加成功")
	}

	return util.JSONSuccess(c, nil, "添加失败")
}

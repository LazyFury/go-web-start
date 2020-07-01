package product

import (
	"EK-Server/model"
	"EK-Server/router/admin/product/cate"
	"EK-Server/util"
	"EK-Server/util/customtype"
	"errors"
	"math"

	"github.com/labstack/echo"
)

var (
	goods = model.Goods{}
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/product"
	product := g.Group(baseURL)

	cate.Init(product)

	product.GET("/list", goods.List)
	product.POST("/add", add)

}

func add(c echo.Context) error {

	db := model.DB
	good := &model.Goods{Title: "default 标题"}
	// 绑定json
	if err := c.Bind(good); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	if good.Price == 0 {
		return util.JSONErr(c, nil, "请填写商品价格")
	}
	// 没有分类id
	if good.Cid == 0 {
		return util.JSONErr(c, nil, "请选择商品分类")
	}
	// 查询分类是否存在
	if good.Cid > 0 {
		if empty := db.First(&model.GoodsCate{BaseControll: model.BaseControll{ID: uint(good.Cid)}}).RecordNotFound(); empty {
			return util.JSONErr(c, nil, "分类不存在")
		}
	}

	money := math.Round(float64(good.Price)*100) / 100 //保留两位小数，接口输出文本是在marshajson的时候处理的
	good.Price = customtype.Money(money)

	db.NewRecord(good) // => 主键为空返回`true`
	row := db.Create(good)

	if row.Error != nil {
		return util.JSONErr(c, row.Error, "添加失败")
	}

	if row.RowsAffected <= 0 {
		return util.JSONSuccess(c, errors.New(""), "添加失败，没有更改")
	}

	return util.JSONSuccess(c, nil, "添加成功")
}

package product

import (
	"EK-Server/model"
	"EK-Server/util"
	"errors"
	"strconv"

	"github.com/labstack/echo"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/product"
	product := g.Group(baseURL)

	product.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "admin")
	})
	product.GET("/add", add)

}

func add(c echo.Context) error {
	price := c.QueryParam("price")
	priceInt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return util.JSONErr(c, nil, "价格错误")
	}

	good := &model.Goods{Title: "default 标题", Price: priceInt}
	db := model.DB

	db.NewRecord(good) // => 主键为空返回`true`
	row := db.Create(good)

	if row.Error != nil {
		return util.JSONErr(c, nil, "添加失败")
	}

	if row.RowsAffected <= 0 {
		return util.JSONSuccess(c, errors.New(""), "添加失败，没有更改")
	}

	return util.JSONSuccess(c, nil, "添加成功")
}

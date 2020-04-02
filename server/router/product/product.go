package product

import (
	"EK-Server/model"
	"EK-Server/util"
	"fmt"
	"strings"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/product"
	product := g.Group(baseURL)

	product.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "hello")
	})

	product.GET("/list", productList)
}

// PageParams PageParams
type PageParams struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Order string `json:"order"`
}

func productList(c echo.Context) error {
	type Param struct {
		PageParams
		Cid int `json:"cid"`
	}
	page := Param{PageParams: PageParams{Page: 1, Limit: 10, Order: "id_desc"}}

	if err := c.Bind(&page); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	fmt.Printf("post json 参数：%v", page)

	if page.Order != "" {
		page.Order = strings.ReplaceAll(page.Order, "_", " ")
		// page.Order = strings.ReplaceAll(page.Order, ",", " ")
	}
	where := &model.Goods{}
	if page.Cid > 0 {
		where = &model.Goods{Cid: page.Cid}
	}

	return util.JSONSuccess(c, model.DataBaselimit(page.Limit, page.Page, where, &[]model.GoodsList{}, "goods",
		page.Order), "获取成功")
}

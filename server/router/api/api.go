package api

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"strconv"
	"suke-go-test/model"
	"suke-go-test/util"
	"time"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/api"
	api := g.Group(baseURL)
	api.GET("/addCate", addCate)
	api.GET("/apiCateAll", apiCateAll)
	api.GET("/apiCateSave", apiCateSave)
	api.GET("/addApi", addAPI)
	api.GET("/allApi", allAPI)
	api.GET("/apiSave", apiSave)
	api.GET("/delApi", delAPI)
	api.GET("/delApiCate", delAPICate)
	api.GET("/cateApi", cateAPI)
}
func delAPI(c echo.Context) (err error) {
	id := c.QueryParam("id")
	if id == "" {
		return util.JSONErr(c, nil, "ID不可空")
	}
	newid, _ := strconv.Atoi(id)

	var t time.Time = time.Now()
	api := model.API{
		Model: gorm.Model{ID: uint(newid), DeletedAt: &t},
	}

	db := model.DB
	if err = db.Model(&model.API{
		Model: gorm.Model{ID: uint(newid)},
	}).Updates(&api).Error; err != nil {
		return util.JSONErr(c, nil, "删除失败")
	}

	return util.JSONSuccess(c, nil, "删除成功")
}

func apiSave(c echo.Context) (err error) {
	id := c.QueryParam("id")
	if id == "" {
		return util.JSONErr(c, nil, "ID不可空")
	}
	newid, _ := strconv.Atoi(id)
	name := c.QueryParam("name")
	if name == "" {
		return util.JSONErr(c, nil, "api名称不可空")
	}
	desc := c.QueryParam("desc")
	if desc == "" {
		return util.JSONErr(c, nil, "api配置不可空")
	}

	db := model.DB

	api := model.API{
		Name: name,
		Data: desc,
		Cid:  "",
	}
	if err = db.Model(&model.API{Model: gorm.Model{ID: uint(newid)}}).Updates(&api).Error; err != nil {
		return util.JSONErr(c, nil, "保存失败")
	}
	return util.JSONSuccess(c, nil, "保存成功")
}

func cateAPI(c echo.Context) (err error) {
	cid := c.QueryParam("cid")
	if cid == "" {
		return util.JSONErr(c, nil, "ID不可空")
	}
	//newid,_ := strconv.Atoi(cid)
	api := model.API{Cid: string(cid)}

	db := model.DB

	var arr []model.API
	row := db.Where(&api).Find(&arr)
	if row.Error != nil {
		return util.JSONErr(c, row.Error, "查询失败")
	}
	//if row.RowsAffected<=0{
	//	return util.JSONSuccess(c,nil,"")
	//}

	return util.JSONSuccess(c, arr, "")
}
func allAPI(c echo.Context) (err error) {
	cate := model.APICate{}
	db := model.DB
	res, msg, err := cate.GetAll()
	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	type result struct {
		model.APICate
		List []model.API `json:"list"`
	}
	var arr []result

	//子查询绑定 未成功
	//Query := fmt.Sprintf("left join test_apis on test_apis.cid=test_api_cates.id")
	//res := db.Table(model.TableName("api_cates")).Select("*").Joins(Query).Scan(&arr)
	//fmt.Printf("%+v\n\n>>>>\n",res)
	//fmt.Printf("%+v\n",arr)

	//循环查询绑定 性能不好
	for i, item := range res {
		fmt.Println(i, item)
		cid := strconv.FormatUint(uint64(item.ID), 10)
		var list []model.API
		db.Where(&model.API{Cid: cid}).Find(&list)

		arr = append(arr, result{
			APICate: item,
			List:    list,
		})
	}
	return util.JSONSuccess(c, arr, "获取成功")
}

func delAPICate(c echo.Context) (err error) {
	db := model.DB
	id := c.QueryParam("id")
	if id == "" {
		return util.JSONErr(c, nil, "ID不可空")
	}
	newid, _ := strconv.Atoi(id)

	api := model.API{Cid: string(newid)}
	if db.Find(&api).RecordNotFound() {
		var t time.Time = time.Now()
		cate := model.APICate{
			Model: gorm.Model{ID: uint(newid), DeletedAt: &t},
		}

		if err = db.Model(&model.APICate{
			Model: gorm.Model{ID: uint(newid)},
		}).Updates(&cate).Error; err != nil {
			return util.JSONErr(c, nil, "删除失败")
		}

		return util.JSONSuccess(c, nil, "删除成功")
	}
	return util.JSONErr(c, nil, "该分类下还有API存在，不可删除")

}
func apiCateSave(c echo.Context) (err error) {
	id := c.QueryParam("id")
	if id == "" {
		return util.JSONErr(c, nil, "ID不可空")
	}
	newid, _ := strconv.Atoi(id)
	name := c.QueryParam("name")
	if name == "" {
		return util.JSONErr(c, nil, "Name不可空")
	}

	cate := model.APICate{
		Model: gorm.Model{ID: uint(newid)},
		Name:  name,
	}
	msg, err := cate.Save()
	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	return util.JSONSuccess(c, nil, msg)
}

//api分类列表
func apiCateAll(c echo.Context) (err error) {
	cate := model.APICate{}
	res, msg, err := cate.GetAll()
	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	return util.JSONSuccess(c, res, msg)
}

//添加分类
func addCate(c echo.Context) (err error) {
	name := c.QueryParam("name")
	if name == "" {
		return util.JSONErr(c, nil, "分类名称不可空")
	}
	desc := c.QueryParam("desc")

	cate := model.APICate{Name: name, Desc: desc}
	msg, err := cate.Add()
	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	return util.JSONSuccess(c, nil, msg)
}

// 添加API
func addAPI(c echo.Context) (err error) {
	name := c.QueryParam("name")
	if name == "" {
		return util.JSONErr(c, nil, "API名称不可空")
	}
	data := c.QueryParam("data")
	if data == "" {
		return util.JSONErr(c, nil, "配置内容不可空")
	}
	cid := c.QueryParam("cid")
	if cid == "" {
		return util.JSONErr(c, nil, "请选择分类")
	}
	api := model.API{
		Name: name,
		Data: data,
		Cid:  cid,
	}

	msg, err := api.Add()

	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	return util.JSONSuccess(c, nil, msg)
}

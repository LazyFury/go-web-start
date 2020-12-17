package model

import (
	"github.com/Treblex/go-web-template/xmodel"
	"gorm.io/gorm"
)

// DataBaselimit  mysql数据分页
func DataBaselimit(limit int, page int, where interface{}, _model Model, key string, orderBy string, userID uint, isAdmin bool) *xmodel.ObjResult {
	list := _model.PointerList()
	listModel := DB.GetObjectsOrEmpty(list, where, func(db *gorm.DB) *gorm.DB {
		return db.Order(orderBy)
	})

	if !_model.IsPublic() && !isAdmin {
		listModel.Model = listModel.Model.Where(map[string]interface{}{
			"user_id": userID,
		})
	}

	listModel.Model = _model.Search(listModel.Model, key)
	if err := listModel.Paging(page, limit, func(db *gorm.DB) *gorm.DB {
		return _model.Joins(db)
	}); err != nil {
		// pass; result empty array
	}
	return listModel.Result
}

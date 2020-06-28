package model

import (
	"fmt"
	"math"
)

//Paging 分页类型
type Paging struct {
	Href   string
	Name   string
	Active bool
}

//GeneratePaging 生成分页html数组
func GeneratePaging(l int, page int, href string) (arr []Paging) {
	// 显示的最大页码  超出8页是会显示省略号隐藏一部分
	var pagingSize int = 8
	var pagingHalf int = int(math.Ceil(float64(pagingSize / 2)))
	// 建立切片
	arr = make([]Paging, l)
	// 填充内容
	for i := range arr {
		active := (page == i+1)
		arr[i].Href = fmt.Sprintf("%s%d", href, i+1)
		if active {
			arr[i].Href = "javascript:;"
		}
		arr[i].Name = fmt.Sprintf("%d", i+1)
		arr[i].Active = active
	}
	// 处理超出隐藏
	if l > 10 {
		start := page - pagingHalf
		end := page + pagingHalf
		//默认 首....2,3,4,5.....尾
		if start < 0 { //首,1,2,3,4...尾
			start = 0
			end = pagingSize
		} else if end > l { //首...3,4,5,6,尾
			start = l - pagingSize
			end = l
		}

		arr = arr[start:end]
	}

	// 添加首尾
	arr = append([]Paging{{Href: fmt.Sprintf("%s%d", href, 1), Name: "首页", Active: page == 1}}, arr...)
	arr = append(arr, Paging{Href: fmt.Sprintf("%s%d", href, l), Name: "末页", Active: page == l})
	return
}

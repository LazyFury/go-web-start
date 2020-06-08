package util

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type (
	TemplateRenderer struct {
		Templates *template.Template
	}
	siteConfig struct {
		SiteName string `json:"siteName"`
	}
	tplFile struct {
		Name string
		Path string
	}
)

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	viewContext, isMap := data.(map[string]interface{})
	if isMap {
		viewContext["reverse"] = c.Echo().Reverse
		viewContext["config"] = &siteConfig{
			SiteName: "塞尔达",
		}
	}

	return t.Templates.ExecuteTemplate(w, name, viewContext)
}

// ParseGlob 自定义模版解析，扫描子目录
func ParseGlob(tpl *template.Template, dir string, pattern string) (t *template.Template, err error) {
	t = tpl
	fmt.Println(">>扫描模版目录：" + dir)
	files := allFiles(dir, pattern)
	for _, file := range files {
		fmt.Println(file.Path)
		b, err := ioutil.ReadFile(file.Path)
		if err != nil {
			return t, err
		}
		s := string(b)
		name := file.Name
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// 目录下的所有文件
func allFiles(dir string, suffix string) (arr []*tplFile) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(">>扫描子目录：" + file.Name())
			arr = append(arr, allFiles(path.Join(dir, file.Name()), suffix)...)
		} else {
			ok, _ := filepath.Match(suffix, file.Name())
			if ok {
				pathName := path.Join(dir, file.Name())
				list := strings.Split(filepath.ToSlash(pathName), "/")
				if len(list) > 1 {
					list = list[1:]
				}
				// fmt.Println(pathName, list)
				name := strings.Join(list, "/")
				arr = append(arr, &tplFile{Name: name, Path: pathName})
			}
		}
	}

	return
}

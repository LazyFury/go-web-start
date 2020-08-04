package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// requestInfo
func requestInfo(c echo.Context) error {
	req := c.Request()
	format := `
<pre>
	<strong>
		Request Information test auto build
	</strong>
	<code>
		Protocol: %s
		Host: %s
		Remote Address: %s
		Method: %s
		Path: %s
	</code>
</pre>
	`
	// fmt.Println(strings.Split(req.Header.Get("Accept"), ",")[0])
	// fmt.Printf("%+v", req.Header)
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}

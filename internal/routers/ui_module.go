package routers

import (
	"apiBook/internal/define"
	"fmt"
	"html/template"
)

func Input(inputType, id, class, name, text string) template.HTML {
	htmlTemplate := `<div class="form-floating %s">
		<input type="%s" class="form-control" id="%s" name="%s" placeholder="%s">
		<label for="%s">%s</label>
	</div>`
	return template.HTML(fmt.Sprintf(htmlTemplate, class, inputType, id, name, name, name, text))
}

func ApiBookInfo() template.HTML {
	htmlTemplate := `<div class="col-sm-4 apibook-info">
		<h2>API Book</h2>
		<p>接口文档管理工具</p>
		<p>https://github.com/mangenotwork/apiBook</p>
		<p>` + define.Version + `</p>
	</div>`
	return template.HTML(htmlTemplate)
}

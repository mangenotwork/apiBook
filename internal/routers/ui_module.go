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

func ProjectCard(pid, name, description string, private define.ProjectPrivateCode) template.HTML {
	privateSpan := ""
	if private == define.ProjectPrivate {
		privateSpan = `<span class="badge text-bg-primary">私有</span>`
	} else {
		privateSpan = `<span class="badge text-bg-info">公有</span>`
	}

	htmlTemplate := `<div class="card mb-3 project">
            <div class="card-body">
                <h5 class="card-title">` + name + `</h5>
                <div class="plabel">` + privateSpan + `</div>
                <p class="card-text">` + description + `</p>
                <a href="/index/` + pid + `" class="card-link">打开文档</a>
                <a onclick="addProject()" class="card-link">设置</a>
                <a href="#" class="card-link">分享</a>
            </div>
        </div>`

	return template.HTML(htmlTemplate)
}

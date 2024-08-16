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

func ProjectCard(pid, name, description string, isOperation int, private define.ProjectPrivateCode) template.HTML {
	privateSpan := ""
	edit := ""
	if isOperation == 1 {
		edit = `<a href="/project/index/` + pid + `" class="card-link">设置</a>`
	}

	if private == define.ProjectPrivate {
		privateSpan = `<span class="badge text-bg-primary">私有</span>`
	} else {
		privateSpan = `<span class="badge text-bg-info">公有</span>`
	}

	//if private == define.ProjectPrivate && isOperation == 1 {
	//	teamWorker = `<a onclick="teamWorkerProject('` + pid + `')" class="card-link">协作者</a>`
	//}

	htmlTemplate := `<div class="card mb-3 project box">
            <div class="card-body">
                <h5 class="card-title">` + name + `</h5>
                <div class="plabel">` + privateSpan + `</div>
                <p class="card-text">` + description + `</p>
                ` + edit + `
                <a href="#" class="card-link">分享</a>
				<a href="/index/` + pid + `" class="card-link">打开文档</a>
            </div>
        </div>`

	return template.HTML(htmlTemplate)
}

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
		<p>接口文档管理工具，私有化部署</p>
		<p>
			<a href="https://github.com/mangenotwork/apiBook" target="_blank" style="color: #ffffff;">
				<svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" class="bi bi-github" viewBox="0 0 16 16">
				  <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"/>
				</svg>
			</a>
		</p>
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

	htmlTemplate := `<div class="card mb-3 project box">
            <div class="card-body">
                <h5 class="card-title t1">` + name + `</h5>
                <div class="plabel">` + privateSpan + `</div>
                <p class="card-text t2">` + description + `</p>
                ` + edit + `
                <a href="#" class="card-link" onclick="openDocShare('` + pid + `')">分享</a>
				<a href="/index/` + pid + `" class="card-link">打开文档</a>
            </div>
        </div>`

	return template.HTML(htmlTemplate)
}

func MethodSelect(id string) template.HTML {
	htmlTemplate := `<select class="form-select" id="` + id + `">
                        <option value="GET">GET</option>
                        <option value="POST">POST</option>
                        <option value="PUT">PUT</option>
                        <option value="HEAD">HEAD</option>
                        <option value="OPTIONS">OPTIONS</option>
                        <option value="DELETE">DELETE</option>
                     </select>`

	return template.HTML(htmlTemplate)
}

func DocNav() template.HTML {
	htmlTemplate := `<nav aria-label="breadcrumb" id="docNav" style="display: none">
                              <ol class="breadcrumb">
                                  <li class="breadcrumb-item" id="returnLink"> <a href="javascript:void(0);" onclick="window.location.reload();">返回文档</a></li>
                                  <li class="breadcrumb-item active" aria-current="page" id="docNavSnapshot"></li>
                              </ol>
                          </nav>`

	return template.HTML(htmlTemplate)
}

func ToastTemplate() template.HTML {
	htmlTemplate := `<div class="toast-container position-fixed p-3 top-50 start-50 translate-middle">
    <div id="liveToast" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
        <div class="toast-header">
            <strong class="me-auto">提示</strong>
            <small></small>
            <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
        </div>
        <div class="toast-body">
            <span id="liveToastMsg"></span>
        </div>
    </div>
</div>`

	return template.HTML(htmlTemplate)
}

func DocMainPoint() template.HTML {
	htmlTemplate := `<div class="col-2" id="docMainML">
                        <div id="simple-list-example" class="d-flex flex-column gap-2 simple-list-example-scrollspy text-center">
                          <a class="p-1 rounded doc-point" href="#simple-list-item-1">基本信息</a>
                          <a class="p-1 rounded doc-point" href="#simple-list-item-2">接口说明</a>
                          <a class="p-1 rounded doc-point" href="#simple-list-item-3">请求Header</a>
                          <a class="p-1 rounded doc-point" href="#simple-list-item-4">请求Body</a>
                          <a class="p-1 rounded doc-point" href="#simple-list-item-5">响应</a>
                          <a class="p-1 rounded doc-point" href="#simple-list-item-7">请求代码</a>
                          <a class="p-1 rounded doc-point" href="#simple-list-item-6">日志&镜像</a>
                        </div>
                    </div>`

	return template.HTML(htmlTemplate)
}

func DocMainBaseInfo() template.HTML {
	htmlTemplate := `<li><span class="txt7">接口名:</span><span id="api-name"></span></li>
                                        <li><span class="txt7">接口Url: </span>
                                            <span style="font-size: 21px;margin-left: 16px;">
                                                <span id="api-method"></span>
                                                <span id="api-url"></span>
                                            </span>
                                        </li>`

	return template.HTML(htmlTemplate)
}

package routers

import (
	"apiBook/common/utils"
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

func ProjectCard(pid, name, description string, isOperation int, private define.ProjectPrivateCode, apiNum int) template.HTML {
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

	privateSpan += `<span class="badge text-bg-secondary">接口数量: ` + utils.AnyToString(apiNum) + `</span>`

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

func ApiBookText() template.HTML {
	htmlTemplate := `<h2>API Book</h2>
		<p>接口文档管理工具，私有化部署</p>
		<p>
			<a href="https://github.com/mangenotwork/apiBook" target="_blank" style="color: black;">
				<svg xmlns="http://www.w3.org/2000/svg" width="36" height="36" fill="currentColor" class="bi bi-github" viewBox="0 0 16 16">
				  <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"/>
				</svg>
			</a>
		</p>
		<p>` + define.Version + `</p>`
	return template.HTML(htmlTemplate)
}

func RequestCode() template.HTML {
	htmlTemplate := `<ul class="nav nav-pills mb-3" id="pills-tab" role="tablist" style="margin-top: 36px;">
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn active" id="code-jsFetch-tab" data-bs-toggle="pill" data-bs-target="#code-jsFetch" type="button" role="tab" aria-controls="code-jsFetch" aria-selected="false">
                                        js-fetch</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-jsAxios-tab" data-bs-toggle="pill" data-bs-target="#code-jsAxios" type="button" role="tab" aria-controls="code-jsAxios" aria-selected="false">
                                        js-axios</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-jsJquery-tab" data-bs-toggle="pill" data-bs-target="#code-jsJquery" type="button" role="tab" aria-controls="code-jsJquery" aria-selected="false">
                                        js-jquery</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-jsXhr-tab" data-bs-toggle="pill" data-bs-target="#code-jsXhr" type="button" role="tab" aria-controls="code-jsXhr" aria-selected="false">
                                        js-xhr</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-swift-tab" data-bs-toggle="pill" data-bs-target="#code-swift" type="button" role="tab" aria-controls="code-swift" aria-selected="false">
                                        swift</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-objectiveC-tab" data-bs-toggle="pill" data-bs-target="#code-objectiveC" type="button" role="tab" aria-controls="code-objectiveC" aria-selected="false">
                                        objective-c</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-dart-tab" data-bs-toggle="pill" data-bs-target="#code-dart" type="button" role="tab" aria-controls="code-dart" aria-selected="false">
                                        dart</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-javaUnirest-tab" data-bs-toggle="pill" data-bs-target="#code-javaUnirest" type="button" role="tab" aria-controls="code-javaUnirest" aria-selected="false">
                                        java-unirest</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-javaOkHttpClient-tab" data-bs-toggle="pill" data-bs-target="#code-javaOkHttpClient" type="button" role="tab" aria-controls="code-javaOkHttpClient" aria-selected="false">
                                        java-okHttpClient</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-curl-tab" data-bs-toggle="pill" data-bs-target="#code-curl" type="button" role="tab" aria-controls="code-curl" aria-selected="true">
                                        curl</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-wget-tab" data-bs-toggle="pill" data-bs-target="#code-wget" type="button" role="tab" aria-controls="code-wget" aria-selected="false">
                                        wget</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-powerShell-tab" data-bs-toggle="pill" data-bs-target="#code-powerShell" type="button" role="tab" aria-controls="code-powerShell" aria-selected="false">
                                        powerShell</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-phpRequest2-tab" data-bs-toggle="pill" data-bs-target="#code-phpRequest2" type="button" role="tab" aria-controls="code-phpRequest2" aria-selected="false">
                                        php-request2</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-phpHttpClient-tab" data-bs-toggle="pill" data-bs-target="#code-phpHttpClient" type="button" role="tab" aria-controls="code-phpHttpClient" aria-selected="false">
                                        php-httpClient</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-phpClient-tab" data-bs-toggle="pill" data-bs-target="#code-phpClient" type="button" role="tab" aria-controls="code-phpClient" aria-selected="false">
                                        php-client</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-pythonClient-tab" data-bs-toggle="pill" data-bs-target="#code-pythonClient" type="button" role="tab" aria-controls="code-pythonClient" aria-selected="false">
                                        python-client</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-pythonRequests-tab" data-bs-toggle="pill" data-bs-target="#code-pythonRequests" type="button" role="tab" aria-controls="code-pythonRequests" aria-selected="false">
                                        python-requests</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-c-tab" data-bs-toggle="pill" data-bs-target="#code-c" type="button" role="tab" aria-controls="code-c" aria-selected="false">
                                        c</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-CSharp-tab" data-bs-toggle="pill" data-bs-target="#code-CSharp" type="button" role="tab" aria-controls="code-CSharp" aria-selected="false">
                                        c#</a>
                                </li>

                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-ruby-tab" data-bs-toggle="pill" data-bs-target="#code-ruby" type="button" role="tab" aria-controls="code-ruby" aria-selected="false">
                                        ruby</a>
                                </li>
                                <li class="nav-item" role="presentation">
                                    <a class="nav-link codeBtn" id="code-go-tab" data-bs-toggle="pill" data-bs-target="#code-go" type="button" role="tab" aria-controls="code-go" aria-selected="false">
                                        go</a>
                                </li>
                            </ul>
                            <div class="tab-content" id="pills-tabContent">
                                <div class="tab-pane fade show active" id="code-jsFetch" role="tabpanel" aria-labelledby="code-jsFetch-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-jsFetch")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-jsAxios" role="tabpanel" aria-labelledby="code-jsAxios-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-jsAxios")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-jsJquery" role="tabpanel" aria-labelledby="code-jsJquery-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-jsJquery")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-jsXhr" role="tabpanel" aria-labelledby="code-jsXhr-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-jsXhr")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-swift" role="tabpanel" aria-labelledby="code-swift-tab" tabindex="0">
                                    <pre data-language="c"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-swift")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-objectiveC" role="tabpanel" aria-labelledby="code-objectiveC-tab" tabindex="0">
                                    <pre data-language="c"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-objectiveC")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-dart" role="tabpanel" aria-labelledby="code-dart-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-dart")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-javaUnirest" role="tabpanel" aria-labelledby="code-javaUnirest-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-javaUnirest")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-javaOkHttpClient" role="tabpanel" aria-labelledby="code-javaOkHttpClient-tab" tabindex="0">
                                    <pre data-language="javascript"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-javaOkHttpClient")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-curl" role="tabpanel" aria-labelledby="code-curl-tab" tabindex="0">
                                    <pre data-language="shell"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-curl")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-wget" role="tabpanel" aria-labelledby="code-wget-tab" tabindex="0">
                                    <pre data-language="shell"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-wget")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-powerShell" role="tabpanel" aria-labelledby="code-powerShell-tab" tabindex="0">
                                    <pre data-language="shell"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-powerShell")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-phpRequest2" role="tabpanel" aria-labelledby="code-phpRequest2-tab" tabindex="0">
                                    <pre data-language="php"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-phpRequest2")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-phpHttpClient" role="tabpanel" aria-labelledby="code-phpHttpClient-tab" tabindex="0">
                                    <pre data-language="php"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-phpHttpClient")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-phpClient" role="tabpanel" aria-labelledby="code-phpClient-tab" tabindex="0">
                                    <pre data-language="php"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-phpClient")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-pythonClient" role="tabpanel" aria-labelledby="code-pythonClient-tab" tabindex="0">
                                    <pre><code data-language="python"></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-pythonClient")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-pythonRequests" role="tabpanel" aria-labelledby="code-pythonRequests-tab" tabindex="0">
                                    <pre><code data-language="python"></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-pythonRequests")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-c" role="tabpanel" aria-labelledby="code-c-tab" tabindex="0">
                                    <pre data-language="c"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-c")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-CSharp" role="tabpanel" aria-labelledby="code-CSharp-tab" tabindex="0">
                                    <pre data-language="c"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-CSharp")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-ruby" role="tabpanel" aria-labelledby="code-ruby-tab" tabindex="0">
                                    <pre data-language="c"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-ruby")'>复制</button>
                                </div>
                                <div class="tab-pane fade" id="code-go" role="tabpanel" aria-labelledby="code-go-tab" tabindex="0">
                                    <pre data-language="c"><code></code></pre>
                                    <button type="button" class="btn btn-dark toolBtn" onclick='copyReqCodeJson("code-go")'>复制</button>
                                </div>
                            </div>`
	return template.HTML(htmlTemplate)
}

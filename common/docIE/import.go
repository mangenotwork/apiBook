package docIE

import (
	"apiBook/internal/define"
	"fmt"
)

// DocImportEr 接口文档导入
type DocImportEr interface {
	Whole(text, userAcc string, private define.ProjectPrivateCode) error // 全量导入，会创建项目再进行导入
	Increment(text, pid, userAcc, dirId string) error                    // 增量导入，导入到指定目录
}

func NewDocImport(source define.SourceCode) (DocImportEr, error) {

	switch source {

	case define.SourceOpenApi301, define.SourceOpenApi310:
		return NewOpenApi301Import(), nil

	case define.SourceSwagger:
		return NewSwaggerImport(), nil

	case define.SourceApiZZA:
		return NewApiZZAImport(), nil

	case define.SourceYApi:
		return NewYApiImport(), nil

	case define.SourceApiBook:
		return NewApiBookImport(), nil

	}

	return nil, fmt.Errorf("未知导入平台")
}

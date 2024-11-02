package docIE

import (
	"apiBook/internal/define"
	"fmt"
)

// DocExportEr 接口文档导出
type DocExportEr interface {
	ExportJson(pid string) string
	Export(pid string) interface{}
}

func NewDocExport(source define.SourceCode) (DocExportEr, error) {

	switch source {

	case define.SourceOpenApi301, define.SourceOpenApi310:
		return NewOpenApiExport(), nil
	case define.SourceApiBook:
		return NewApiBookExport(), nil

	}

	return nil, fmt.Errorf("未知导入平台")
}

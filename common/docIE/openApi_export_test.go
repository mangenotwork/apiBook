package docIE

import (
	"testing"
)

func Test_OpenApiExport(t *testing.T) {
	obj := NewOpenApiExport()
	obj.ExportJson("1ae58da57aeb11f7")
}

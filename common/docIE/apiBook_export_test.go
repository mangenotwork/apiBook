package docIE

import (
	"testing"
)

func Test_ApiBookExport(t *testing.T) {
	obj := NewApiBookExport()
	obj.ExportJson("1ae58da57aeb11f7")
}

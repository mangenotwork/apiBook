package docIE

import (
	"apiBook/common/utils"
	"testing"
	"time"
)

func Test_ApiBookExport(t *testing.T) {
	obj := NewApiBookExport()
	obj.ExportJson("1ae58da57aeb11f7")
}

func Test_FixedNumber(t *testing.T) {

	result := make(map[int64][]int64)

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Microsecond)
		num := utils.ID()
		t.Log(num)

		fixed := fixedNumber(int64(num))
		//t.Log(i, " --> ", fixed)

		if _, ok := result[fixed]; !ok {
			result[fixed] = make([]int64, 0)
		}

		result[fixed] = append(result[fixed], int64(num))

	}

	t.Log(result)

}

var fl int64 = 100

func fixedNumber(id int64) int64 {
	n1 := id / 1 % 10
	n2 := id / 10 % 10
	n3 := id / 100 % 10
	n4 := id / 1000 % 10
	var r1 int64 = 50
	if n1 != 0 {
		r1 = ((n1 + n2 + n3 + n4) * n1) % fl
	} else if n2 != 0 {
		r1 = ((n1 + n2 + n3 + n4) * n2) % fl
	} else if n3 != 0 {
		r1 = ((n1 + n2 + n3 + n4) * n3) % fl
	} else if n4 != 0 {
		r1 = ((n1 + n2 + n3 + n4) * n4) % fl
	}
	return r1
}

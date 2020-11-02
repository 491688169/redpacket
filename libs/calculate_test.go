package libs_test

import (
	"redpacket/libs"
	"testing"
)

func TestDoubleAverage2(t *testing.T) {
	doubleAverageTable := []struct {
		count int64
		mount int64
	}{
		{1, 100},
		{2, 100},
		{101, 100},
	}

	for _, tt := range doubleAverageTable {
		count, mount := tt.count, tt.mount
		sum := int64(0)
		remain := mount
		for i := int64(0); i < count; i++ {
			current := libs.DoubleAverage2(remain, count-i)
			sum += current
			remain -= current
		}
		if sum != mount {
			t.Fatal("sum != mount")
		}
	}

}

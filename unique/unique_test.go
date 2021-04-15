package unique_test

import (
	"github.com/zzhlong/tool/unique"
	"testing"
)

func BenchmarkUUID32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if len(unique.UUID32()) != 32 {
			b.Errorf("%s", "生成的uuid长度不为32")
		}
	}
}

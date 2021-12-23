package kalkan

import "testing"

func BenchmarkSignXML(b *testing.B) {
	const inputXML = "<root>GoKalkan</root>"
	for i := 0; i < b.N; i++ {
		cli.SignXML(inputXML)
	}
}

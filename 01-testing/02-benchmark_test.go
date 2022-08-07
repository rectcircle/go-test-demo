package testingdemo

import (
	"bytes"
	"html/template"
	"math/rand"
	"testing"
)

// go test -run=^$ -benchmem -bench ^BenchmarkRandInt$ ./01-testing

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// go test -run=^$ -benchmem -bench ^BenchmarkTemplateParallel$ ./01-testing

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

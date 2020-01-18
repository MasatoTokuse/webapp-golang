package trace

import (
	"fmt"
	"io"
)

// Tracer はコード内の出来事を記録できるオブジェクトを表すインターフェースです
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// New は指定のio.Writerに出力するTracerを返します
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off はTracerメソッドの呼び出しを無視するTracerを返します
func Off() Tracer {
	return &nilTracer{}
}

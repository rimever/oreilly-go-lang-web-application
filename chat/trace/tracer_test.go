package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buffer bytes.Buffer
	tracer := New(&buffer)
	if tracer == nil {
		t.Error("return value form New is nil")
	}else {
		tracer.Trace("Hello trace package!")
		if buffer.String() != "Hello trace package!\n" {
			t.Errorf("output illegal text as '%s'",buffer.String())
		}
	}
}

func TestOff(t *testing.T)  {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Data")
}
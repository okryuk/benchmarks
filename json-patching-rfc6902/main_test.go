package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func compare(got, want []byte, t *testing.T) {
	var g, w any
	if err := json.Unmarshal(got, &g); err != nil {
		t.Fatalf("error unmarshalling reply")
	}
	if err := json.Unmarshal(want, &w); err != nil {
		t.Fatalf("error unmarshalling want")
	}
	if eq := reflect.DeepEqual(g, w); eq != true {
		t.Fatalf("Response doesn't match expected")
	}

}

func TestSJSONAdd(t *testing.T) {
	got := addSJSON(DocB, AddOp)
	want := []byte(WantAddStr)
	compare(got, want, t)
}

func TestJSONPATCHAdd(t *testing.T) {
	got := opJSONPATCH(DocB, AddOp)
	want := []byte(WantAddStr)
	compare(got, want, t)
}

func TestUniJSONAdd(t *testing.T) {
	got := uniSJSONwithRawMessage(DocB, AddOp)
	want := []byte(WantAddStr)
	compare(got, want, t)
}

func TestUniJSONOptimiAdd(t *testing.T) {
	got := uniSJSONOptim(DocB, AddOp)
	want := []byte(WantAddStr)
	compare(got, want, t)
}

func TestSJSONReplace(t *testing.T) {
	got := replaceSJSON(DocB, ReplaceOp)
	want := []byte(WantReplaceStr)
	compare(got, want, t)
}

func TestJSONPATCHReplace(t *testing.T) {
	got := opJSONPATCH(DocB, ReplaceOp)
	want := []byte(WantReplaceStr)
	compare(got, want, t)
}

func TestUniJSONReplace(t *testing.T) {
	got := uniSJSONwithRawMessage(DocB, ReplaceOp)
	want := []byte(WantReplaceStr)
	compare(got, want, t)
}

func TestUniJSONOptimReplace(t *testing.T) {
	got := uniSJSONOptim(DocB, ReplaceOp)
	want := []byte(WantReplaceStr)
	compare(got, want, t)
}

func TestSJSONRemove(t *testing.T) {
	got := removeSJSON(DocB, RemoveOp)
	want := []byte(WantRemoveStr)
	compare(got, want, t)
}

func TestJSONPATCHRemove(t *testing.T) {
	got := opJSONPATCH(DocB, RemoveOp)
	want := []byte(WantRemoveStr)
	compare(got, want, t)
}

func TestUniJSONRemove(t *testing.T) {
	got := uniSJSONwithRawMessage(DocB, RemoveOp)
	want := []byte(WantRemoveStr)
	compare(got, want, t)
}

func TestUniJSONOptimRemove(t *testing.T) {
	got := uniSJSONOptim(DocB, RemoveOp)
	want := []byte(WantRemoveStr)
	compare(got, want, t)
}

// =============================================================================
// Benchmarks section

func Benchmark_Add_JSONPATCH(b *testing.B) {
	for n := 0; n < b.N; n++ {
		opJSONPATCH(DocB, AddOp)
	}
}

func Benchmark_Add_SJSON_rawOperation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		addSJSON(DocB, AddOp)
	}
}

func Benchmark_Add_UniJSONPATCH_EQ_withRawMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniSJSONwithRawMessage(DocB, AddOp)
	}
}

func Benchmark_Add_UniJSONPATCH_EQ_Optimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniSJSONOptim(DocB, AddOp)
	}
}

func Benchmark_Replace_JSONPATCH(b *testing.B) {
	for n := 0; n < b.N; n++ {
		opJSONPATCH(DocB, ReplaceOp)
	}
}

func Benchmark_Replace_SJSON_rawOperation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		replaceSJSON(DocB, ReplaceOp)
	}
}

func Benchmark_Replace_UniJSONPATCH_EQ_withRawMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniSJSONwithRawMessage(DocB, ReplaceOp)
	}
}

func Benchmark_Replace_UniJSONPATCH_EQ_Optim(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniSJSONOptim(DocB, ReplaceOp)
	}
}

func Benchmark_Remove_JSONPATCH(b *testing.B) {
	for n := 0; n < b.N; n++ {
		opJSONPATCH(DocB, RemoveOp)
	}
}

func Benchmark_Remove_SJSON_rawOperation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		removeSJSON(DocB, RemoveOp)
	}
}

func Benchmark_Remove_UniJSONPATCH_EQ_withRawMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniSJSONwithRawMessage(DocB, RemoveOp)
	}
}

func Benchmark_Remove_UniJSONPATCH_EQ_Optim(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uniSJSONOptim(DocB, RemoveOp)
	}
}

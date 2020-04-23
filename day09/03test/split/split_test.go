package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("fdfsdf","f")
	want := []string{"", "d", "sd", ""}
	if !reflect.DeepEqual(got,want){
		t.Errorf("%v，%v",got,want)
		//t.Fail()
	}
}

func TestSplit1(t *testing.T) {
	got := Split("fd","")
	want := []string{"f", "d"}
	if !reflect.DeepEqual(got,want){
		t.Errorf("%v，%v",got,want)
		//t.Fail()
	}
}

//基准测试

func BenchmarkSplit(b *testing.B) {
	for i :=0; i<b.N;i++{
		Split("a:b:c:f:f",":")
	}
}
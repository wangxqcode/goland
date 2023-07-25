package split

import (
	"reflect"
	"testing"
)



func TestMoreSplit(t *testing.T) {

	got := Split("abcd", "bc")

	want := []string{"a", "d"}

	if !reflect.DeepEqual(want,got){
		t.Errorf("excepted:%v, got:%v", want, got)
	}


}




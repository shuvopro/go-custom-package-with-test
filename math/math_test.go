package math

import "testing"

func TestAverage( t *testing.T){
	v := Average([]float64{1,2,3})

	if v != 2 {
		t.Error("Expected 2, got ", v)
	}

}
package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	str := check("012201210001021201")
	if (str != "The word W contains cube or patterns like xxx") {
		t.Error("First error")
	}
	str2 := check("0120212012102010212")

	if (str2 != "The word W does not contain cube or patterns like xxx") {
		t.Error("Second error")
	}

}

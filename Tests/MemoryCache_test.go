package Tests

import (
	"testing"
	"github.com/ruslanfedoseenko/gorutracker/Misc"
)

func TestAssignDifferentTypes(t *testing.T)  {
	c := Misc.NewMemoryCache()
	v1 := 22;
	v2 := ""
	c.Put("v", v1)
	err := c.Get("v",v2)
	if err == nil{
		t.Error("Type conversions are not allowed")
	}
}

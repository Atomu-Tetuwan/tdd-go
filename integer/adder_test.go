package integer

import (
	"fmt"
	"testing"
)

func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	// OutPut: 6
}
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expecetd := 4

	if sum != expecetd {
		t.Errorf("expecetd %d but got %d", expecetd, sum)
	}
}
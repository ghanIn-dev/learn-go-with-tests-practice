package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// Output: 을 적지 않으면 테스트에 포함되지 않는다
// 편해 보이긴 하는데 아래 주석의 형태처럼 format을 맞춰주지않으면 적용이 안된다
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

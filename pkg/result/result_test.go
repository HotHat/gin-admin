package result

import (
	"fmt"
	"testing"
)

func divide(a, b int) Result[float64] {
	if b == 0 {
		return Err[float64](fmt.Errorf("cannot divide by zero"))
	}
	return Ok(float64(a) / float64(b))
}

func TestResult(t *testing.T) {
	res1 := divide(10, 2)
	if res1.IsOk() {
		res1.Than(func(val float64) float64 {
			return val * 10
		})
		flt := res1.Unwrap().(float64)
		fmt.Println("Result:", flt)
		fmt.Println("Result:", flt == 50)

	}

	res2 := divide(10, 0)
	fmt.Println("Is res2 OK?", res2.IsOk())

}

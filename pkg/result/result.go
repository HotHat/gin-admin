package result

// Result is a generic type that can hold either a successful value or an error.
type Result[T any] struct {
	value T
	err   error
}

type BoolResult = Result[bool]

var (
	BoolOk = Result[bool]{}
)

// Ok creates a new Result with a successful value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// Err creates a new Result with an error.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// IsOk returns true if the result contains a successful value.
func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

// Unwrap returns the value or panics if an error is present.
func (r *Result[t]) Unwrap() any {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

func (r *Result[T]) Than(fun func(val T) T) {
	r.value = fun(r.value)
}

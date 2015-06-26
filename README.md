# Trial

A simple assertion library for go.

## Motivation

go-test is fine. We do not need big testing frameworks for most projects. 
One thing its lacking though is simple assertions and ability to implement simple helpers as it does not allow us to skip
test logs.

Trial gives us `th.Error(t testingT, skip int, msgs ...interface{})` allowing to skip callers to implement helpers with nice logging.

For most uses, the `trial/assert` package is enough, giving us the most basic assertions needed with nice error messages.

## trial/assert Usage

**Simple equals**

```
assert.Equal(t, 1, 2)

Output:
unit_test.go:42: Not equal:
		Expected: 1
		  Actual: 2
```


**Additional arguments** to overwrite the message

```
assert.Equal(t, 1, 2, "numbers dont match for %q", "my param")

Output:
unit_test.go:42: numbers dont match for "my param":
		Expected: 1
		  Actual: 2
```


**Type problems** are made clear

```
assert.Equal(t, 1, int64(1))

unit_test.go:42: Not equal:
		Expected: 1
		  Actual: 1
		   Types: Expected:int, Actual:int64
```

See example/example_test.go for more.


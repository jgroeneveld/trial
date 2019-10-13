# Trial [![GoDoc](https://godoc.org/github.com/jgroeneveld/trial/assert?status.svg)](https://godoc.org/github.com/jgroeneveld/trial/assert) [![GoReport](https://goreportcard.com/badge/github.com/jgroeneveld/trial)](https://goreportcard.com/report/github.com/jgroeneveld/trial)

A simple assertion library for go. Also see [schema](https://github.com/jgroeneveld/schema) for easier JSON Schema testing.

## Motivation

go-test is fine. We do not need big testing frameworks for most projects. 
One thing its lacking though is simple assertions and ability to implement simple helpers as it does not allow us to skip
test logs.

Trial gives us `th.Error(t testingT, skip int, msgs ...interface{})` allowing to skip callers to implement helpers with nice logging.

For most uses, the `trial/assert` package is enough, giving us the most basic assertions needed with nice error messages.

## trial/assert Usage

**Simple equals**

```
import "github.com/jgroeneveld/trial/assert"

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

See [example/example_test.go](example/example_test.go) for more.

### Supported Assertions

**Basic assertions**

```
Equal(expected, actual, msgf...)
MustBeEqual(expected, actual, msgf...)
NotEqual(expected, actual, msgf...)
MustNotBeEqual(expected, actual, msgf...)
DeepEqual(expected, actual, msgf...)
MustBeDeepEqual(expected, actual, msgf...)
True(expression bool, msgf...)
MustBeTrue(expression bool, msgf...)
False(expression bool, msgf...)
MustBeFalse(expression bool, msgf...)
Nil(expression, msgf...)
MustBeNil(expression, msgf...)
NotNil(expression, msgf...)
MustNotBeNil(expression, msgf...)
```

**JSON Schema assertions**

```
JSONSchema(reader, matcher, msgf...)
MustMatchJSONSchema(reader, matcher, msgf...)
```

## Writing your own assertions
`th` can be used to write simple own assertions. This for example gives you a wrapper for [schema](https://github.com/jgroeneveld/schema).MatchJSON to have simple JSON schema assertions in your tests:

```go
func AssertJSONSchema(t *testing.T, matcher schema.Matcher, r io.Reader) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
	}
}

func MustMatchJSONSchema(t *testing.T, matcher schema.Matcher, r io.Reader) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
		t.FailNow()
	}
}
```


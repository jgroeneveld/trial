package assert

import (
	"testing"
)

func TestEqual_includingOutput(t *testing.T) {
	mt := &mockT{}

	Equal(mt, "a", "a")
	Equal(mt, "a", "b")
	Equal(mt, 1, int64(1), "Test (%s)", "value")

	if len(mt.errors) != 2 {
		t.Fatalf("wrong number of errors %d", len(mt.errors))
	}

	actual := mt.errors[0]
	expected := "\r" + `assert_test.go:11: Not equal:
Expected: "a"
  Actual: "b"`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}

	actual = mt.errors[1]
	expected = "\r" + `assert_test.go:12: Test (value):
Expected: 1
  Actual: 1
   Types: Expected:int, Actual:int64`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}
}

func TestEqual(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		Equal(mt, 1, 1)
	})

	testAssertionFail(t, func(mt *mockT) {
		Equal(mt, 1, 2)
	})
}

func TestNotEqual(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		NotEqual(mt, 1, 2)
	})

	testAssertionFail(t, func(mt *mockT) {
		NotEqual(mt, 1, 1)
	})
}

func TestDeepEqual(t *testing.T) {
	type twoFields struct {
		foo string
		bar string
	}
	type oneField struct {
		foo string
	}

	testAssertionSuccess(t, func(mt *mockT) {
		DeepEqual(mt, oneField{"hello"}, oneField{"hello"})
	})

	testAssertionFail(t, func(mt *mockT) {
		DeepEqual(mt, oneField{"hello"}, oneField{"world"})
	})

	testAssertionFail(t, func(mt *mockT) {
		DeepEqual(mt,
			twoFields{"hello", "world"},
			oneField{"hello"},
		)
	})
}

func TestTrue(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		True(mt, true)
	})

	testAssertionFail(t, func(mt *mockT) {
		True(mt, false)
	})
}

func TestFalse(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		False(mt, false)
	})

	testAssertionFail(t, func(mt *mockT) {
		False(mt, true)
	})
}

func TestNil(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		Nil(mt, nil)
	})

	testAssertionFail(t, func(mt *mockT) {
		Nil(mt, 1)
	})
}

func TestNotNil(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		NotNil(mt, 1)
	})

	testAssertionFail(t, func(mt *mockT) {
		NotNil(mt, nil)
	})
}

func TestMustBeEqual(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustBeEqual(mt, 1, 1)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustBeEqual(mt, 1, 2)
	})
}

func TestMustNotBeEqual(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustNotBeEqual(mt, 1, 2)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustNotBeEqual(mt, 1, 1)
	})
}

func TestMustBeDeepEqual(t *testing.T) {
	type twoFields struct {
		foo string
		bar string
	}
	type oneField struct {
		foo string
	}

	testAssertionSuccess(t, func(mt *mockT) {
		MustBeDeepEqual(mt, oneField{"hello"}, oneField{"hello"})
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustBeDeepEqual(mt, oneField{"hello"}, oneField{"world"})
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustBeDeepEqual(mt,
			twoFields{"hello", "world"},
			oneField{"hello"},
		)
	})
}

func TestMustBeTrue(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustBeTrue(mt, true)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustBeTrue(mt, false)
	})
}

func TestMustBeFalse(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustBeFalse(mt, false)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustBeFalse(mt, true)
	})
}

func TestMustBeNil(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustBeNil(mt, nil)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustBeNil(mt, 1)
	})
}

func TestMustNotBeNil(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustNotBeNil(mt, 1)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustNotBeNil(mt, nil)
	})
}

func testAssertionSuccess(t *testing.T, successCase func(mt *mockT)) {
	mt := &mockT{}
	successCase(mt)
	if len(mt.errors) > 0 {
		t.Fatalf("expected no errors but got %#v", mt.errors)
	}
}

func testAssertionFail(t *testing.T, failureCase func(mt *mockT)) {
	mt := &mockT{}
	failureCase(mt)
	if len(mt.errors) == 0 {
		t.Fatalf("expected errors but got none")
	}
}

func testAssertionFailNow(t *testing.T, failureCase func(mt *mockT)) {
	mt := &mockT{}
	failureCase(mt)
	if len(mt.errors) == 0 {
		t.Fatalf("expected errors but got none")
	}
	if !mt.failedNow {
		t.Fatalf("expected to fail now but did not")
	}
}

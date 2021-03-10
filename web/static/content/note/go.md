---
title: "Notes on Go"
date: 2021-03-08T18:09:52Z
draft: true
url: "/go"
---

## Mocking with Interfaces

```go
package foo

type Tweeter interface {
	Tweet(s string) error
}

func SendMessage(t Tweeter, s string) error {
  // ...
  return t.Tweet(s)
}
```

```go
package foo_test

import (
  "testing"

  "github.com/revett/foo"
  "github.com/stretchr/testify/assert"
)

type mockTweeter struct{}

func (m mockTweeter) Tweet(s string) error {
  // ...
  return nil
}

func TestSendMessage(t *testing.T) {
  // ...
  err := foo.SendMessage(mockTweeter{}, "...")
  assert.NoError(t, err)
}
```

Links:

- [“Mocking Golang with Interfaces In Real Life” by Jonatas Baldin](https://dev.to/jonatasbaldin/mocking-golang-with-interfaces-in-real-life-3f1m)

## Different Testing Package

```go
package foo

// ...
```

```go
package foo_test

import (
  "testing"

  "github.com/revett/foo"
)
```

If unexported code must be tested, then create another file with
`_internal_test.go` as the suffix which imports `foo`.

## Table Driven Tests

```go
package foo

func Bar(s string) bool {
  // ...
}
```

```go
package foo_test

import (
  "testing"

  "github.com/revett/foo"
  "github.com/stretchr/testify/assert"
)

func TestBar(t *testing.T) {
  tests := map[string]struct {
    input string
    want string
  }{
    "Valid": {
      input: "a",
      want: "b",
    },
    "Invalid": {
      input: "a",
      want: "x",
    },
  }

  for n, tc := range tests {
    t.Run(n, func(t *testing.T) {
      ok := foo.Bar(tc.input)
      assert.Equal(t, tc.want, ok)
    })
  }
}
```

Links:

- [“Prefer table driven tests” by @davecheney](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## Resources

- [gophercises.com](https://gophercises.com)
- [dariubs/GoBooks](https://github.com/dariubs/GoBooks)
- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

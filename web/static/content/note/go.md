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

- ["Mocking Golang with Interfaces In Real Life" by Jonatas Baldin (dev.to)](https://dev.to/jonatasbaldin/mocking-golang-with-interfaces-in-real-life-3f1m)

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

- ["Prefer table driven tests" by Dave Cheney (dave.cheney.net)](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## Package Names

- Short and clear
- Often simple nouns (e.g. `time`, `list`, `http`)
- Lower case, with no underscores or `mixedCase`
- Avoid package and function name stutter
- Abbreviate only if the name will be unambiguous (e.g. `strconv`, `syscall`, `fmt`)
- Don't steal good names (e.g. `bufio` instead of `buf`)
- Avoid generic names (e.g. `util`, `common`, `misc`)
- Avoid exposing all API interfaces in a single package (e.g. `types`, `models`)
- Avoid unnecessary name collisions (e.g. using same name as popular `http` package)

Links:

- ["Package names" by Sameer Ajmani (blog.golang.org)](https://blog.golang.org/package-names)

## Naming Acronyms

- Consistent case
- Either all lowercase or uppercase (e.g. `url` or `URL`)
- Never use `mixedCase` (e.g. `Url`)

Links:

- ["Go Code Review Comments - Initialisms" (github.com)](https://github.com/golang/go/wiki/CodeReviewComments#initialisms)

## Resources

- [gophercises.com](https://gophercises.com)
- [dariubs/GoBooks](https://github.com/dariubs/GoBooks)
- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

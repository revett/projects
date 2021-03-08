---
title: "Notes on Go"
date: 2021-03-08T18:09:52Z
draft: true
url: "/go"
---

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

Links:

- [“Prefer table driven tests” by @davecheney](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

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

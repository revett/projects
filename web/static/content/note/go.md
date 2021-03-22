---
title: "Go"
draft: true
---

## Closing HTTP Response Body

```go
r, err := http.Get("https://example.com")
if err != nil {
	return err
}
defer r.Body.Close()
```

> The client must close the response body when finished with it.

- This must be done even if `r.Body` is not consumed
- Not closing can lead to a resource leak where the connection is not re-used

Links:

- [Package HTTP Overview (golang.org)](https://golang.org/pkg/net/http/#pkg-overview)
- [http.Client (godoc)](https://golang.org/pkg/net/http/#Client.Do)

## Examples

Go allows snippets of code to act as examples, which are:

- Displayed as package documentation via godoc
- Verified to be functional by running each of them as tests

This makes sure that a package has up-to-date documentation as it is part of the
test suite, instead of being within a `README` outside of the Go codebase.

Files which include examples must:

- Start with `example`
- End with `_test.go`

Simple file would be `example_test.go` however examples can be split across
multiple files, for example `example_interface_test.go` and
`example_search_test.go`.

Basic example:

```go
package stringutil_test

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func Example() {
	fmt.Println(stringutil.Reverse("hello"))
}
```

Note that the exported function starts with `Example` and takes no arguments.

Examples can be scoped to a specific exported identifier:

```go
func ExampleFoo()     // documents the Foo function or type
func ExampleBar_Qux() // documents the Qux method of type Bar
func Example()        // documents the package as a whole
```

Multiple examples can be provided for the same identifier:

```go
func ExampleReverse()
func ExampleReverse_second()
func ExampleReverse_third()
```

Links:

- ["Testable Examples in Go" by Andrew Gerrand](https://blog.golang.org/examples)
- [`sort` package examples](https://golang.org/src/sort/)

## Functional Options

```go
package chess

func Start(s string, opts ...func(*Engine) error) (*Engine, error) {
	e, err := startEngine(s)
	if err != nil {
		return nil, err
	}

	for _, o := range opts {
		if err := o(e); err != nil {
			return nil, err
		}
	}

	return e, nil
}

func Debug(e *Engine) error {
	e.debug = true
	return nil
}

func WithTimeout(d time.Duration) func(*Engine) error {
	return func(e *Engine) error {
		e.timeout = d
		return nil
	}
}
```

Default use case is simple:

```go
e, err := chess.Start("/path/to/engine")
```

Adding more complex initialisation through readable arguments:

```go
e, err := chess.Start(
	"/path/to/engine",
	chess.Debug,
	chess.WithTimeout(200*time.Millisecond),
)
```

Considerations:

- Default use case for an API is simple to understand
- Quick and easy to expand an API in the future with further configuration options
- Parameters are well documented
- Removes the need to export struct fields as they can now be explicitly modified
- Number and order of arguments does not matter

Links:

- ["Functional options for friendly APIs" by Dave Cheney (dave.cheney.net)](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

## Godoc

Basic example for documenting a type, variable, constant or function:

```go
// Fprint formats using the default formats for its operands and writes to w.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

Comments on a package will render the text within an `Overview` section of the
HTML godoc page:

```go
// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort
```

> See [sort](https://golang.org/pkg/sort/).

If the general documentation for the package is much longer, then it should be
within a `doc.go`:

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package gob manages streams of gobs - binary values exchanged between an
Encoder (transmitter) and a Decoder (receiver). A typical use is transporting
arguments and results of remote procedure calls (RPCs) such as those provided by
package "net/rpc".

Types and Values

The source and destination values/types need not correspond exactly. For structs,
fields (identified by name) that are in the source but absent from the receiving
variable will be ignored. Fields that are in the receiving variable but missing
from the transmitted type or value will be ignored in the destination. If a field
with the same name is present in both, their types must be compatible. Both the
receiver and transmitter will do all necessary indirection and dereferencing to
convert between gobs and actual Go values. For instance, a gob type that is
schematically,

	struct { A, B int }

can be sent from or received into any of these Go types:

	struct { A, B int }	// the same
	*struct { A, B int }	// extra indirection of the struct
	struct { *A, **B int }	// extra indirection of the fields
	struct { A, B int64 }	// different concrete value type; see below
*/
package gob
```

> Taken from [doc.go](https://golang.org/src/encoding/gob/doc.go) in
> [encoding/gob](https://golang.org/pkg/encoding/gob/).

Note that:

- **Paragraphs** are delineated by blank lines
- **Headings** are a single line without punctuation, with paragraphs before and
  after it
- Code snippets must be indented
- Godoc will generate a table of contents automatically

Known bugs can be documented (where `user` is someone who can provide more
information):

```go
// BUG(user): Description of known bug.
```

Links:

- ["Godoc: documenting Go code" by Andrew Gerrand](https://blog.golang.org/godoc)
- [natefinch/godocgo (pkg.go.dev)](https://pkg.go.dev/github.com/natefinch/godocgo)

## Higher-Order Function

- A function that operates on other functions
- It must either:
  - Recieve a (_first-class_) function as an argument
  - Return a function as output

## Map Key Exists

```go
if _, ok := m["key"]; ok {
	// ...
}
```

## Mocks

### Higher-Order Function

```go
package page

import "net/http"

type Getter func(string) (*http.Response, error)

func IsHTML(g Getter, u string) (bool, error) {
	r, err := g(u)
	// ...
}
```

Full example with tests: [github.com/revett/snippets/internal/hofunc/page](https://github.com/revett/snippets/tree/main/internal/hofunc/page)

Considerations:

- Callers of `.IsHTML` will need to import `http`, expanding the list of
  dependencies for packages that call the function
- May expand the function arguments list beyond what is reasonable to read
- Function may be more difficult to understand due to passing in logic that may
  not be clearly linked

Links:

- ["Mocking Techniques for Go" (myhatchpad.com)](https://www.myhatchpad.com/insight/mocking-techniques-for-go/)

### Interface Substitution

```go
package page

import "net/http"

type Getter interface {
	Get(s string) (r *http.Response, err error)
}

func IsHTML(g Getter, u string) (bool, error) {
	r, err := g.Get(u)
	// ...
}
```

Full example with tests: [github.com/revett/snippets/internal/isub/page](https://github.com/revett/snippets/tree/main/internal/isub/page)

### Monkey Patching

```go
package page

import "net/http"

var Getter = http.Get

func IsHTML(u string) (bool, error) {
	r, err := Getter(u)
	// ...
}
```

Full example with tests: [github.com/revett/snippets/internal/monkpatch/page](https://github.com/revett/snippets/tree/main/internal/monkpatch/page)

Considerations:

- Parallel tests will not function correctly
- Exporting the variable `Getter` will allow anyone to change it
- `Getter` being exported pollutes the package as only useful for testing

## Naming

### Interface Names

- One method interfaces are named by the method with an `-er` suffix
- This applies even if the result is not perfect English (e.g. `Execer` for `.Exec`)
- Reordering is best if it will help with readability (e.g. `ByteReader` for `.ReadByte`)

Links:

- ["Effective Go: Interface Names" (golang.org)](https://golang.org/doc/effective_go#interface-names)
- ["What's in a name?" (golang.org)](https://talks.golang.org/2014/names.slide#13)

### Naming Acronyms

- Consistent case
- Either all lowercase or uppercase (e.g. `url` or `URL`)
- Never use `mixedCase` (e.g. `Url`)

Links:

- ["Go Code Review Comments - Initialisms" (github.com)](https://github.com/golang/go/wiki/CodeReviewComments#initialisms)

### Package Names

- Short and clear
- Often simple nouns (e.g. `time`, `list`, `http`)
- Lower case, with no underscores or `mixedCase`
- Avoid package and function name stutter
- Abbreviate only if the name will be unambiguous (e.g. `strconv`, `syscall`, `fmt`)
- Don't steal good variable names (e.g. `bufio` instead of `buf`)
- Avoid generic names (e.g. `util`, `common`, `misc`)
- Avoid exposing all API interfaces in a single package (e.g. `types`, `models`)
- Avoid unnecessary name collisions (e.g. using same name as popular `http` package)

Links:

- ["Package names" by Sameer Ajmani (blog.golang.org)](https://blog.golang.org/package-names)

## Semantic Package Versioning

Follow the format: `vMAJOR.MINOR.PATCH` (note the `v...` prefix)

- **Major**: a backwards incompatible change to the public API of the module
- **Minor**: a backwards compatible change to the API, like changing
  dependencies or adding a new function, method, struct field, or type
- **Patch**: a change that does not affect the public API or dependencies, like
  fixing a bug

Pre-release versions can be specified however users must specifically request
them as normal releases are preferred by the `go` command. Examples:

- `v0.3.0-alpha`
- `v1.0.0-beta`

Releases must be continued as the `go` command will always use the greatest
semantic release version available, even if it far behind the primary branch.

## Pretty Printing Data Structures

```go
fmt.Printf("%+v\n", s)
```

```go
package main

import "https://github.com/davecgh/go-spew"

func main() {
	spew.Dump(s)
}
```

```go
d, _ := json.MarshalIndent(s, "", "\t")
fmt.Println(string(d))
```

## Testing

### Different Package

```go
package page

// ...
```

```go
package page_test

import (
	"testing"

	"github.com/revett/snippets/internal/isub/page"
)

// ...
```

If unexported code must be tested, then another file with
`_internal_test.go` as the suffix can be created.

### HTML Coverage Report

```bash
go test -v ./... -cover -coverprofile=coverage.out
```

```bash
go tool cover -html=coverage.out
```

### Table Driven

```go
package page_test

import (
	"errors"
	"testing"

	"github.com/revett/snippets/internal/isub/page"
	"github.com/stretchr/testify/assert"
)

func TestIsHTML(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		mg   mockGetter
		want bool
		err  bool
	}{
		"Simple": {
			mg: mockGetter{
				ct: "text/html",
			},
			want: true,
			err:  false,
		},
		"Error": {
			mg: mockGetter{
				err: errors.New("error"),
			},
			want: false,
			err:  true,
		},
		// ...
	}

	for n, tc := range tests {
		tc := tc

		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ok, err := page.IsHTML(tc.mg, "https://example.com")
			assert.Equal(t, tc.err, err != nil)
			assert.Equal(t, tc.want, ok)
		})
	}
}

// ...
```

Full example: [github.com/revett/snippets/internal/isub/page](https://github.com/revett/snippets/tree/main/internal/isub/page)

Note:

- That the `tc` variable is reassigned within the body of the loop to avoid
  using the wrong iterated variable, as `t.Run` is a goroutine.

Links:

- ["Prefer table driven tests" by Dave Cheney (dave.cheney.net)](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- ["Using Goroutines on Loop Iterator Variables" (go/wiki)](https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables)

## Resources

- [gophercises.com](https://gophercises.com)
- [dariubs/GoBooks](https://github.com/dariubs/GoBooks)
- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

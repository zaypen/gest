# gest
A simple wrapper of `testing.T` to provide DSL-like testing pattern in Go.

# Install

> go get github.com/zaypen/gest

Or

> dep ensure -add github.com/zaypen/gest

# Usage

```go
package xxx

import (
	"testing"
	"github.com/zaypen/gest"
)

func TestCompareVersion(t *testing.T)  {
	gest.I(t).Should("1 + 2 == 3").Expect(1 + 2).ToBe(3)
}
```
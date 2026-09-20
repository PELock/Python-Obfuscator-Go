# Python Obfuscator — Go Web API SDK

Go client for the [Python Obfuscator](https://www.pelock.com/products/python-obfuscator) Web API. Obfuscate, virtualize, and protect Python `.py` scripts.

API documentation: https://www.pelock.com/products/python-obfuscator/api

Endpoint: `https://www.pelock.com/api/python-obfuscator/v1`

## Installation

```bash
go get github.com/PELock/Python-Obfuscator-Go
```

## Usage

```go
package main

import (
	"context"
	"fmt"

	pythonobfuscator "github.com/PELock/Python-Obfuscator-Go"
)

func main() {
	client := pythonobfuscator.New("YOUR-WEB-API-KEY")
	client.CodeVirtualization = pythonobfuscator.CodeVirtualizationVM
	client.RenameStyle = pythonobfuscator.RenameStyleIL

	result, err := client.ObfuscateScriptSource(context.Background(), "print('hello')")
	if err != nil {
		panic(err)
	}
	if result.Error == pythonobfuscator.ErrorSuccess {
		fmt.Println(result.Output)
	}
}
```

`CodeVirtualization` defaults to `vm`. Set it to `""` to skip virtualization. Optional zlib compression is off by default.

See `examples/`.

## License

Apache-2.0. Copyright Bartosz Wójcik / PELock.

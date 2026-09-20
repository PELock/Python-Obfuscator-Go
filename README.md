# Python Obfuscator — Obfuscate, Virtualize & Protect Python Code

**[Python Obfuscator](https://www.pelock.com/products/python-obfuscator)** is a tool to obfuscate, virtualize & protect Python `.py` scripts against reverse engineering with a VM engine (or finite-state automata machine), self-defending integrity checks, polymorphic string encryption, anti-debugging, anti-vm, anti-emulation and 50+ total obfuscation strategies!

It's available for Windows & Linux, plus VSCode/Cursor extension:

* https://www.pelock.com/products/python-obfuscator/download

Multiple programming APIs available:

* https://www.pelock.com/products/python-obfuscator/api

An online obfuscator interface:

* https://www.pelock.com/python-obfuscator/

## Why Python scripts need obfuscation?

Scripts are typically distributed as plain `.py` files or bundled inside packages. That convenience means anyone with file access can read the full logic, hunt for credentials or API keys in strings, and steal your algorithms unless you take extra steps to hide intent.

[Python](https://www.python.org/) is one of the most popular general-purpose scripting languages. It is widely used for automation, backend services, data processing, tooling, and licensing/agent code that customers run on their own machines.

## Obfuscation strategies

Python Obfuscator comes with many advanced obfuscation, virtualization & protection strategies. You can easily tune protection versus size and performance.

![Python Obfuscation options](https://www.pelock.com/img/en/products/python-obfuscator/python-obfuscator-virtualizer-obfuscation-options.png)

### ![Obfuscation](https://www.pelock.com/img/en/icons/obfuscation-32.png) Powerful obfuscation

Polymorphic string encryption is the core of the string pipeline: each build generates a fresh decryptor so literals never share one static algorithm. Integers, floats, and decoy noise join that layer. The result conceals literals and structure while preserving tested runtime behaviour.

### ![Processor](https://www.pelock.com/img/en/icons/processor-32.png) Code virtualization

Code virtualization is the main control-flow defense. Selected statements are lifted into a randomly generated VM engine (or FSA / flattened dispatcher) with shuffled dispatch tables, decoy opcodes, and an obfuscated dispatcher loop. Analysts must interpret the virtual machine instead of reading plain Python.

### ![Node](https://www.pelock.com/img/en/icons/node-32.png) Finite-state automata (FSA)

Finite-state automata (FSA) obfuscation rewrites linear Python statement blocks into dual-state automata with opaque schedulers and shuffled dispatch handlers. Instead of reading code top to bottom, analysts must follow numeric states, transition tables, and decoy paths to reconstruct the original order.

### ![Bug](https://www.pelock.com/img/en/icons/bug-32.png) Anti-debugging

Anti-debugging protection inserts polymorphic probes that detect attached debuggers, tracing hooks, and related host signals, together with anti-VM, anti-sandbox and anti-emulator checks. When a check fires, the obfuscated script exits silently instead of revealing protected logic under interactive analysis.

### ![System monitor](https://www.pelock.com/img/en/icons/system-monitor-32.png) Self-integrity checks

Self-defending integrity checks verify that the obfuscated file has not been patched. A bootstrap probe hashes the on-disk script and sets a tamper key when it no longer matches the obfuscated build. String decryptors consume that key, so edited scripts return garbage instead of plaintext.

### ![Brick link](https://www.pelock.com/img/en/icons/brick-link-32.png) Protection linker

The protection linker adds decoy functions and fake calls so a copied fragment still looks like real program code. Hidden traps fire only if someone edits the file or runs a piece of it on its own. When the script starts normally, those extras stay silent and the program runs as usual.

## Before and after obfuscation

Look at this example — the same script becomes harder to read at a glance after obfuscation.

### Sample Python script before obfuscation

```python
label = 'SecretKey'
port = 443

def get_sum(a, b):
    return a + b

print(label, port)
r = get_sum(11, 31)
print(r)
```

### After obfuscation

![Obfuscated Python script](https://www.pelock.com/img/en/products/python-obfuscator/python-obfuscator-virtualizer-obfuscated-virtualized-python-code.png)

Would you still recognise the original intent if you only had the obfuscated text and no prior copy of the script?

### How does Python Obfuscator work?

The engine parses Python source into an AST tree, then applies selectable transforms. Code virtualization (VM, FSA, or flattening) rewrites execution so analysts cannot read the script top to bottom. Polymorphic string encryption hides literals; self-defending integrity probes detect patched files. Also available: identifier renaming, numeric encryption, noise and decoy insertion, the protection linker, and anti-debugging checks. Many techniques are specific to this product; some ideas are shared with our other protection tools.

![Python Obfuscator Pipeline](https://www.pelock.com/img/en/products/python-obfuscator/python-obfuscator-obfuscation-pipeline.png)

When all passes finish, the engine emits a new `.py` file. Edge cases in the Python grammar and hosting environments mean you should always test the output in your target runtime.

## Protect your Python scripts & algorithms

Take no chances, use **Python Obfuscator** to obfuscate, virtualize and protect your Python scripts and algorithms.

Our company has a long history in obfuscation technologies and code obfuscators (see our [PowerShell](https://www.pelock.com/products/powershell-pro-obfuscator), [Java](https://www.pelock.com/products/jobfuscator), [AutoIt](https://www.pelock.com/products/autoit-obfuscator), [x86 Assembly](https://www.pelock.com/products/obfuscator) &  obfuscators).

We actively bugfix, research and develop new obfuscation strategies for our tools.

You can count on our expertise and support in this field.

### Installation

The preferred way to install the Web API SDK is via Go modules.

```bash
go get github.com/PELock/Python-Obfuscator-Go
```

The module path is `github.com/PELock/Python-Obfuscator-Go`.

## Usage examples

### Example — `login/main.go`

```go
/******************************************************************************
 * Python Obfuscator WebApi interface usage example.
 *
 * Version        : v1.0.0
 * Language       : Go
 * Author         : Bartosz Wójcik
 * Web page       : https://www.pelock.com
 *
 *****************************************************************************/

package main

import (
	"context"
	"fmt"
	"os"

	pythonobfuscator "github.com/PELock/Python-Obfuscator-Go"
)

func main() {
	client := pythonobfuscator.New("ABCD-ABCD-ABCD-ABCD")
	result, err := client.Login(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Demo version status - %v\n", result.Demo)
	fmt.Printf("License expiration - %s\n", result.LicenseExpiration)
	fmt.Printf("Total obfuscations - %d\n", result.UsagesTotal)
	fmt.Printf("Max. source code size - %d\n", result.StringLimit)
}
```

### Example — `simple/main.go`

```go
/******************************************************************************
 * Python Obfuscator WebApi interface usage example.
 *
 * Version        : v1.0.0
 * Language       : Go
 * Author         : Bartosz Wójcik
 * Web page       : https://www.pelock.com
 *
 *****************************************************************************/

package main

import (
	"context"
	"fmt"
	"os"

	pythonobfuscator "github.com/PELock/Python-Obfuscator-Go"
)

func main() {
	client := pythonobfuscator.New("ABCD-ABCD-ABCD-ABCD")

	source := `label = 'SecretKey'
port = 443

def get_sum(a, b):
    return a + b

print(label, port)
r = get_sum(11, 31)
print(r)
`
	result, err := client.ObfuscateScriptSource(context.Background(), source)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if result.Error != pythonobfuscator.ErrorSuccess {
		fmt.Fprintf(os.Stderr, "An error occurred, error code: %d\n", result.Error)
		os.Exit(1)
	}
	fmt.Println(result.Output)
}
```

See the `examples/` directory in this repository for complete samples.

## Use Python Obfuscator Online

Online interface for Python Obfuscator is available at:

https://www.pelock.com/python-obfuscator/

## Windows GUI client and command line version

You can download it at:

https://www.pelock.com/products/python-obfuscator/download

#### Python Obfuscator comes also with full GUI version for Windows

![Python Obfuscator Windows Client](https://www.pelock.com/img/en/products/python-obfuscator/python-obfuscator-virtualizer.png)

#### Obfuscation options

![Python Obfuscation options](https://www.pelock.com/img/en/products/python-obfuscator/python-obfuscator-virtualizer-obfuscation-options.png)

#### Command line interface aka CLI

Python Obfuscator ships with a command-line interface for Windows and Linux automation. Use it to integrate obfuscation into build servers, CI jobs, or batch packaging.

![Python Obfuscator command-line interface](https://www.pelock.com/img/en/products/python-obfuscator/python-obfuscator-virtualizer-obfuscation-from-command-line.png)

## Demo mode limitations

In demo mode the obfuscator always applies `integers_to_arrays`, `mba_binops`, and `encrypt_strings` (no code virtualization). Source size is limited to 1000 characters.

Bartosz Wójcik

* Visit my site at — https://www.pelock.com
* X — https://x.com/PELock
* GitHub — https://github.com/PELock

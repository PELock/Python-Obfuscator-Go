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

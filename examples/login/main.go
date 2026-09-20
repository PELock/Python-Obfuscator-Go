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

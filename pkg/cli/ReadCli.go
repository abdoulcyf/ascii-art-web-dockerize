//====================ReadCli=========================
package cli

import (
	"errors"
	"fmt"
	"os"
)

/*
write a function that read the arguments of command-lines and return
- it must have just 1 arguments (filename)
- err : more than 1 arguments
*/
func ReadCli() (string, error) {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Error: more than 1 argument in cli")
		return "", errors.New("more than 1 input arguments")
	}

	cliStr := args[1]

	if cliStr == "" {
		return "", errors.New("empty string")
	}

	return cliStr, nil
}

//====================END===========================

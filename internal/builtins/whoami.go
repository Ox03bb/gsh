package builtins

import (
	"fmt"
	"gsh/internal/command"
	"os/user"
)

func Whoami(sh ShellContext, c command.Command) error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	fmt.Println(currentUser.Username)
	return nil
}

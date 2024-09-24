package main

import "os"

func commandExit(state *replState, args []string) error {
	os.Exit(0)
	return nil
}

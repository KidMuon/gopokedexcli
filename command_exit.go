package main

import "os"

func commandExit(state *replState) error {
	os.Exit(0)
	return nil
}

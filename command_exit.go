package main

import "os"

func commandExit(cfg *config, s string) error {
	os.Exit(0)
	return nil
}

package main

import "os"

func commandExit(cfg *config, _ []string) error {
	os.Exit(0)
	return nil
}

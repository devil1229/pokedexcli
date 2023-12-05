package main

import "os"

func commandExit(currentConfig *config) error {
	os.Exit(0)
	return nil
}
package main

import "os"

func commandExit(currentConfig *config, arg string) error {
	//os.exit(0) just exit the application
	os.Exit(0)
	return nil
}
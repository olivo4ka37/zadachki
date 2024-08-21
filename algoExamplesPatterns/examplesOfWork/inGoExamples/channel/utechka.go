package main

import "os"

func readFiles(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		// Do something with file
	}

	return nil
}

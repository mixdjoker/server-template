package config

import "github.com/joho/godotenv"

// Load reads environment variables from a file at the given path
// and loads them into the current environment.
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

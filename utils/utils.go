package utils

import "os"

func LoadEnv() map[string]string {
	env := make(map[string]string)
	env["backPort"] = os.Getenv("BE_PORT")
	return env
}

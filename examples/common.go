package examples

import (
	"flag"
	"os"
)

// ParseArguments looks for username and password in environment variables and program flags.
func ParseArguments() (username, password string) {
	flag.StringVar(&username, "username", os.Getenv("USERNAME"), "Username to authenticate with")
	flag.StringVar(&password, "password", os.Getenv("PASSWORD"), "Password to authenticate with")
	flag.Parse()

	return username, password
}

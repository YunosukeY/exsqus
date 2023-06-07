package util

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func rootPath() (string, error) {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(path)), nil
}

func envPath() (string, error) {
	root, err := rootPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, ".env"), nil
}

func loadEnv() error {
	path, err := envPath()
	if err != nil {
		return err
	}
	return godotenv.Overload(path)
}

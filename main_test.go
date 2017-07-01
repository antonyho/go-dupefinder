package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// TODO Create directories
	// TODO Create identical files seperated in different directories
	exitCode := m.Run()
	// TODO Remove created files
	// TODO Check file list

	os.Exit(exitCode)
}

func BenchmarkMain(b *testing.B) {
}

package main

import (
	"flag"
	"os"
	"testing"
)

/*
func TestMain(m *testing.M) {
	// TODO Create directories
	// TODO Create identical files seperated in different directories
	exitCode := m.Run()
	// TODO Remove created files
	// TODO Check file list

	os.Exit(exitCode)
}
*/
func TestMain(m *testing.M) {
	var dir string
	flag.StringVar(&dir, "dir", "~", "Root directory to scan")
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func BenchmarkMain(b *testing.B) {
}

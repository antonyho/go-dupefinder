package main

import (
	"github.com/antonyho/go-dupefinder/database"
	"github.com/antonyho/go-dupefinder/file"
	"github.com/gocarina/gocsv"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
	"fmt"
	"time"
)

func main() {
	var path string
	app := cli.NewApp()
	app.Name = "Dupefinder"
	app.Usage = "Find duplicated files for you and report in CSV file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dir, d",
			Value:       "/",
			Usage:       "Base `DIRECTORY` to start the scanning",
			Destination: &path,
		},
	}
	app.Action = func(c *cli.Context) error {
		// Build cache database
		spin := spinMsg("Building cache")
		cache := database.New()
		if err := filepath.Walk(path, database.Store(cache)); err != nil {
			return err
		}
		spin.Stop()
		fmt.Println()

		// Query duplicated files
		spin = spinMsg("Finding potential duplicated files")
		groups, err := cache.ListDuplicated()
		if err != nil {
			return err
		}
		spin.Stop()
		fmt.Println()

		// Create CSV report file
		reportFile, err := os.OpenFile("report.csv", os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		defer reportFile.Close()

		// Compare the files
		spin = spinMsg("Comparing files")
		for _, group := range groups {
			sort.Sort(file.BySize{group.Files})
			if group.Files[0].Size != group.Files[len(group.Files)-1].Size {
				// TODO Check this group
				message := "Non-identical group discovered with same checksum."
				separator := strings.Repeat("=", utf8.RuneCountInString(message))
				log.Println(separator)
				log.Println(message)
				for _, f := range group.Files {
					log.Printf("\n%s %d %v\n", f.Path, f.Size, f.ModificationTime)
				}
				log.Println(separator)
				continue // Skip this group
			}
			// TODO Check files are identical in byte
			// Report identical files to text file
			gocsv.MarshalFile(&(group.Files), reportFile)
		}
		spin.Stop()
		fmt.Println()

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

var (
	spinner = []string{"-", "\\", "|", "/"}
)
func spinMsg(msg string) *time.Ticker {
	n := 0
	rotateTicker := time.NewTicker(time.Second)
	go func() {
		for _ = range rotateTicker.C {
			fmt.Printf("\r%-2s%-75s", spinner[n%4], msg)
			n++
		}
	}()

	return rotateTicker
}
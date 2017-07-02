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
)

func main() {
	var path string
	app := cli.NewApp()
	app.Flags = []cli.Flag(
		cli.StringFlag{
			Name:        "dir, d",
			Value:       "/",
			Usage:       "Base `DIRECTORY` to start the scanning",
			Destination: &path,
		},
	)
	app.Action = func(c *cli.Context) error {
		// Build cache database
		cache := database.New()
		if err := filepath.Walk(path, database.Store(cache)); err != nil {
			return err
		}
		// Query duplicated files
		groups, err := cache.ListDuplicated()
		if err != nil {
			return err
		}

		// Create CSV report file
		reportFile, err := os.OpenFile("report.csv", os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		defer reportFile.Close()

		// Compare the files
		for _, group := range groups {
			sort.Sort(file.BySize{group.Files})
			if group.Files[0].Size != group.Files[len(group.Files)-1].Size {
				// TODO Check this group
				message := "Non-identical group discovered with same checksum."
				seperator := strings.Repeat("=", utf8.RuneCountInString(message))
				log.Println(seperator)
				log.Println(message)
				for _, f := range group.Files {
					log.Printf("%s %d %v", f.Path, f.Size, f.ModificationTime)
				}
				log.Println(seperator)
				continue // Skip this group
			}
			// TODO Check files are identical in byte
			// Report identical files to text file
			gocsv.MarshalFile(&(group.Files), reportFile)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

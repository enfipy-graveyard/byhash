package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	"golang.org/x/crypto/sha3"
)

func main() {
	log.SetFlags(0)

	app := cli.NewApp()
	app.Name = "byhash"
	app.Usage = "get files hash & get by hash"

	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Validate hash",
			Action: func(c *cli.Context) error {
				data := "hello world"
				hash := toKeccak([]byte(data))
				log.Printf("Data: %s, Hash: %x\n", data, hash)
				return nil
			},
		},
		{
			Name:  "ls",
			Usage: "Retrieves hash from file data in provided folder",
			Action: func(c *cli.Context) error {
				folder := "dist"

				files, err := ioutil.ReadDir(folder)
				handleError(err)
				for _, file := range files {
					if file.IsDir() {
						continue
					}

					fileName := file.Name()
					bytes, err := ioutil.ReadFile(folder + "/" + fileName)
					handleError(err)

					hash := toKeccak(bytes)
					log.Printf("%s %x", fileName, hash)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func toKeccak(data []byte) []byte {
	hash := make([]byte, 64)

	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(data)
	hash = keccak.Sum(nil)

	return hash
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Exited with error: %v", err)
	}
}

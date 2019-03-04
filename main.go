package byhash

import (
	"encoding/hex"
	"errors"
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
				folder := getFolder(c)
				iterateByFolderFiles(folder, func(fileName string) {
					bytes, err := ioutil.ReadFile(folder + "/" + fileName)
					handleError(err)

					hash := toKeccak(bytes)
					log.Printf("%s %x", fileName, hash)
				})
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Get data by hash in provided folder",
			Action: func(c *cli.Context) error {
				folder := getFolder(c)

				hash := c.Args().Get(1)
				if len(hash) != 64 {
					handleError(errors.New("Invalid hash"))
				}

				iterateByFolderFiles(folder, func(fileName string) {
					bytes, err := ioutil.ReadFile(folder + "/" + fileName)
					handleError(err)

					fileHash := hex.EncodeToString(toKeccak(bytes))
					if fileHash == hash {
						log.Printf("%s", string(bytes))
					}
				})
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	handleError(err)
}

func toKeccak(data []byte) []byte {
	hash := make([]byte, 64)

	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(data)
	hash = keccak.Sum(nil)

	return hash
}

func getFolder(c *cli.Context) string {
	folder := c.Args().First()
	if folder == "" {
		handleError(errors.New("Invalid folder"))
	}
	return folder
}

func iterateByFolderFiles(folder string, handler func(string)) {
	files, err := ioutil.ReadDir(folder)
	handleError(err)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		handler(file.Name())
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Exited with error: %v", err)
	}
}

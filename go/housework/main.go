package main

import (
	"encoding/gob"
	"encoding/json"
	"experiments/go/housework/v1/service"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

type Chore struct {
	Complete    bool
	Description string
}

var dataFile string

func init() {
	flag.StringVar(&dataFile, "file", "housework.db", "data file")
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			`Usage: %s [flags] [add chore, ...|complete #]
        add         add comma-separated chores
        complete    complete designated chore
    	Flags:
    	`,
			filepath.Base(os.Args[0]),
		)
		flag.PrintDefaults()
	}
}

func load() ([]*Chore, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return make([]*Chore, 0), nil
	}

	df, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := df.Close(); err != nil {
			fmt.Printf("closing data file: %v", err)
		}
	}()

	return GobLoad(df)
}

func flush(chores []*Chore) error {
	df, err := os.Create(dataFile)
	if err != nil {
		return err
	}

	defer func() {
		if err := df.Close(); err != nil {
			fmt.Printf("closing data file: %v", err)
		}
	}()

	return GobFlush(df, chores)
}

func list() error {
	chores, err := load()
	if err != nil {
		return err
	}

	if len(chores) == 0 {
		fmt.Println("You're all caught up!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores {
		c := " "
		if chore.Complete {
			c = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, c, chore.Description)
	}

	return nil
}

func add(s string) error {
	chores, err := load()
	if err != nil {
		return err
	}

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores = append(chores, &Chore{
				Description: desc,
			})
		}
	}

	return flush(chores)
}

func complete(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	chores, err := load()
	if err != nil {
		return err
	}

	if i < 1 || i > len(chores) {
		return fmt.Errorf("chore %d not found", i)
	}

	chores[i-1].Complete = true
	return flush(chores)
}

func main() {
	flag.Parse()

	var err error
	switch strings.ToLower(flag.Arg(0)) {
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	case "complete":
		err = complete(flag.Arg(1))
	}

	if err != nil {
		log.Fatal(err)
	}

	err = list()
	if err != nil {
		log.Fatal(err)
	}
}

// JSON (de)serialization
func StorageLoad(r io.Reader) ([]*Chore, error) {
	var chores []*Chore
	return chores, json.NewDecoder(r).Decode(&chores)
}

func StorageFlush(w io.Writer, chores []*Chore) error {
	return json.NewEncoder(w).Encode(chores)
}

// Gob (de)serialization
func GobLoad(r io.Reader) ([]*Chore, error) {
	var chores []*Chore
	return chores, gob.NewDecoder(r).Decode(&chores)
}

func GobFlush(w io.Writer, chores []*Chore) error {
	return gob.NewEncoder(w).Encode(chores)
}

func Load(r io.Reader) ([]*service.Chores, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var chores service.Chores
	return chores.Chores, proto.Unmarshal(b, &chores)
}

func Flush(w io.Writer, chores []*Chore) error {
	b, err := proto.Marshal(&Chores{Chores: chores})
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

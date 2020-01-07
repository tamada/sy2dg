package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"github.com/tamada/sy2dg"
	"github.com/tamada/sy2dg/ksu"
)

func (opts *options) toNewFileName(name, newext string) string {
	newfile := filepath.Join(opts.dest, filepath.Base(name))
	ext := filepath.Ext(newfile)
	if len(ext) > 0 && ext != newext {
		return newfile[:len(newfile)-len(ext)] + newext
	}
	return newfile
}

func (opts *options) store(data *sy2dg.SyllabusData, name string) {
	json, err := json.Marshal(data)
	if err != nil {
		return
	}
	newfile := opts.toNewFileName(name, ".json")
	file, err := os.OpenFile(newfile, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return
	}
	file.Write(json)
}

func (opts *options) performEach(name string, parser sy2dg.Parser) (*sy2dg.SyllabusData, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	syllabus, err2 := parser.Parse(file, name)
	if err2 != nil {
		return nil, err2
	}
	return syllabus, nil
}

func (opts *options) perform() int {
	parser := ksu.NewHTMLParser()
	for _, file := range opts.args {
		data, err := opts.performEach(file, parser)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		opts.store(data, file)
	}
	return 0
}

func helpMessage(prog string) string {
	return fmt.Sprintf(`%s version %s
%s [OPTIONS] <ARGUMENTS...>
OPTIONS
    -d, --dest=<DIR>     specifies destination directory.
    -h, --help           print this message.
ARGUMENTS
    html syllabus documents.`, prog, sy2dg.Version, prog)
}

type options struct {
	dest     string
	helpFlag bool
	args     []string
}

func buildFlagSet(args []string) (*flag.FlagSet, *options) {
	opts := new(options)
	flags := flag.NewFlagSet("ksu2json", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.StringVarP(&opts.dest, "dest", "d", ".", "specifies destination directory")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message")
	return flags, opts
}

func parseArgs(args []string) (*options, error) {
	flags, opts := buildFlagSet(args)
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	opts.args = flags.Args()[1:]
	return opts, nil
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if opts.helpFlag {
		fmt.Println(helpMessage("ksu2json"))
		return 0
	}
	return opts.perform()
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}

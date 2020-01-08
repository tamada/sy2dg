package main

import (
	"encoding/json"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/sy2dg"
)

type options struct {
	pattern  string
	helpFlag bool
	args     []string
}

func readSyllabuses(opts *options, parser sy2dg.Parser, dir string) []*sy2dg.SyllabusData {
	builder := sy2dg.NewSyllabusBuilder(parser, opts.pattern, "")
	return builder.ReadSyllabuses(dir)
}

func performEach(opts *options, parser sy2dg.Parser, dir string) error {

	syllabuses := readSyllabuses(opts, parser, dir)
	ds := sy2dg.NewDataSet(syllabuses)
	json, err := json.Marshal(ds)
	if err != nil {
		return err
	}
	fmt.Printf("const dataset = %s", string(json))

	return nil
}

func perform(opts *options) int {
	parser := sy2dg.NewJSONParser()
	if len(opts.args) != 1 {
		fmt.Printf("the length of args should be 1 (%d)", len(opts.args))
		return 1
	}
	if err := performEach(opts, parser, opts.args[0]); err != nil {
		fmt.Println(err.Error())
	}
	return 0
}

func helpMessage(prog string) string {
	return fmt.Sprintf(`%s version %s
%s [OPTIONS] <DIR>
OPTIONS
    -t, --target <PATTERN>    specifies the pattern of target file name in the SYLLABUSES_DIR.
    -h, --help                print this message.
DIR
    directories contain the json files.`, prog, sy2dg.Version, prog)
}

func buildFlagSet(args []string) (*flag.FlagSet, *options) {
	opts := new(options)
	flags := flag.NewFlagSet("json2dg", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.StringVarP(&opts.pattern, "target", "t", `[0-9]+\.json`, "specifies pattern.")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message")
	return flags, opts
}

func goMain(args []string) int {
	flagSet, opts := buildFlagSet(args)
	if err := flagSet.Parse(args); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if opts.helpFlag {
		fmt.Println(helpMessage("json2dg"))
		return 0
	}
	opts.args = flagSet.Args()[1:]
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}

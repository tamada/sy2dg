package main

import (
	"encoding/json"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/sy2dg"
	"github.com/tamada/sy2dg/ksu"
)

type options struct {
	pattern  string
	parser   string
	helpFlag bool
	baseURL  string
	args     []string
}

func readSyllabuses(opts *options, parser sy2dg.Parser, dir string) []*sy2dg.SyllabusData {
	builder := sy2dg.NewSyllabusBuilder(parser, opts.pattern, opts.baseURL)
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

func newParser(kind string) sy2dg.Parser {
	if kind == "json" {
		return sy2dg.NewJSONParser()
	}
	return ksu.NewHTMLParser()
}

func perform(opts *options) int {
	parser := newParser(opts.parser)
	for _, dir := range opts.args {
		if err := performEach(opts, parser, dir); err != nil {
			fmt.Println(err.Error())
		}
	}
	return 0
}

func getHelpMessage(prog string) string {
	return fmt.Sprintf(`%s version %s
%s [OPTIONS] <SYLLABUSES_DIR>
OPTIONS
    -h, --help                print this message.
    -p, --parser <PARSER>     specifies parser of a syllabus. default is "default".
                              available values are: "default", and "json".
    -u, --url <BASE_URL>      specifies the base of URL for syllabus data.
    -t, --target <PATTERN>    specifies the pattern of target file name in the SYLLABUSES_DIR.
ARGUMENTS
    SYLLABUSES_DIR            the directory containing the syllabuses data.`, prog, sy2dg.Version, prog)
}

func buildFlagSet(args []string) (*flag.FlagSet, *options) {
	opts := new(options)
	flags := flag.NewFlagSet("sy2dg", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(getHelpMessage(args[0])) }
	flags.StringVarP(&opts.pattern, "target", "t", `[0-9]+\.html`, "specifies the pattern of target file name")
	flags.StringVarP(&opts.parser, "parser", "p", "default", "specifies parser of a syllabus")
	flags.StringVarP(&opts.baseURL, "url", "u", "", "specifies the base url of the syllabuses")
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
		fmt.Println(getHelpMessage("sy2dg"))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}

[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/syllabuses2graph/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/syllabuses2graph/releases/tag/v1.0.0)

# sy2dg

This product converts syllabuses data to a directed graph for representing the relations of them.
Generally, syllabuses have the description of already taken lectures and future lectures.
These relations are quite important to plan for the diploma.

## Usage

```sh
$ syllabuses2graph -h
syllabuses2graph [OPTIONS] <SYLLABUSES_DIR>
OPTIONS
    -h, --help                print this message.
    -p, --parser=<PARSER>     specifies parser of a syllabus.
    -t, --target=<PATTERN>    specifies the pattern of target file name in the SYLLABUSES_DIR.
ARGUMENTS
    SYLLABUSES_DIR            the directory containing the syllabuses data.
```

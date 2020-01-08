[![codebeat badge](https://codebeat.co/badges/923a4d5e-961c-4eb5-99a8-4601175732b4)](https://codebeat.co/projects/github-com-tamada-sy2dg-master)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/syllabuses2graph/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/syllabuses2graph/releases/tag/v1.0.0)

# sy2dg

This product converts syllabuses data to a directed graph for representing the relations of them.
Generally, syllabuses have the description of already taken lectures and future lectures.
These relations are quite important to plan for the diploma.

## Usage

This products contains the following products.

* [`ksu2json`](#ksu2json)
    * converts syllabuses in the HTML format for the KSU to json format of `sy2dg`.
* [`json2dg`](#json2dg)
    * converts json data to directed graph for visualizing by [`dgviewer`](#dgviewer)
* [`sy2dg`](#sy2dg)
    * converts syllabuses by parsing specified parser to json format of `sy2dg`
* [`dgviewer`](#dgviewer)
    * visualizing the directed graph of syllabuses.

### `ksu2json`

```sh
ksu2json [OPTIONS] <ARGUMENTS...>
OPTIONS
    -d, --dest <DIR>     specifies destination directory.
    -h, --help           print this message.
ARGUMENTS
    html syllabus documents.
```

### `json2dg`

```sh
json2dg [OPTIONS] <DIR>
OPTIONS
    -t, --target <PATTERN>    specifies the pattern of target file name in the SYLLABUSES_DIR.
    -h, --help                print this message.
DIR
    directories contain the json files.
```

### `sy2dg`

```sh
sy2dg [OPTIONS] <SYLLABUSES_DIR>
OPTIONS
    -h, --help                print this message.
    -p, --parser <PARSER>     specifies parser of a syllabus. default is "default".
                              available values are: "default", and "json".
    -u, --url <BASE_URL>      specifies the base of URL for syllabus data.
    -t, --target <PATTERN>    specifies the pattern of target file name in the SYLLABUSES_DIR.
ARGUMENTS
    SYLLABUSES_DIR            the directory containing the syllabuses data.
```

### `dgviewer`

Graph viewer of `sy2dg` and `json2dg`.

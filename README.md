# ignore-file

Parse and test an ignore file, such as `.gitignore`

# Install

Pick an msi package [here](https://github.com/mh-cbon/ignore-file/releases)!

__chocolatey__

```sh
choco install ignore-file
```

__deb/rpm repositories__

```sh
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/ignore-file sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/ignore-file sh -xe
```

__deb/rpm packages__

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/ignore-file sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/ignore-file sh -xe
```

__go__

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon
cd $GOPATH/src/github.com/mh-cbon
git clone https://github.com/mh-cbon/ignore-file.git
cd ignore-file
glide install
go install
```


# Usage as a lib

```go
package main

import (
	"fmt"
	"github.com/mh-cbon/ignore-file/ignored"
)


func main() {

  ignore := ignored.Ignored{}
  if err := ignore.Load(".gitignore"); err != nil {
    fmt.Println(err)
    return nil
  }

  if err := ignore.Append(".git"); err != nil {
    fmt.Println(err)
    return nil
  }

  computed := ignore.ComputeDirectory(".")
  for _, l := range computed {
    fmt.Println(l)
  }
}
```

# Usage as a binary

```sh
NAME:
   ignore-file - Test a ignore file

USAGE:
   ignore-file <ignore file path> <options>

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --json, -j              return json response
   --also value, -a value  Add new rule
   --help, -h              show help
   --version, -v           print the version

EXAMPLE:
  ignore-file .gitignore
  ignore-file -a .git .gitignore
  ignore-file -a .git -a someother .gitignore
```

In other words,

```sh
$ ignore-file -a .git .gitignore
/.gitignore
/README.md
/fixtures/dira/fileb
/fixtures/dira/other
/fixtures/dira/some
/fixtures/dira/someother
/fixtures/dirb/fileb
/fixtures/dirb/other/fileb
/fixtures/dirb/other/some
/fixtures/filea
/fixtures/other
/fixtures/othera
/fixtures/otherfile
/fixtures/some
/fixtures/someother
/glide.lock
/glide.yaml
/ignored/index.go
/ignored/index_test.go
/main.go
```

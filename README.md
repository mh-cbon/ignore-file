# ignore-file

Parse an `ignore` file, such as `.gitignore`,
apply it on the directory of the file
and print non-ignored files.

# Install

# Usage

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

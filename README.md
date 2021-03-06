# enr

[![Go Report Card](https://goreportcard.com/badge/github.com/gost-c/enr)](https://goreportcard.com/report/github.com/gost-c/enr)
[![Build Status](https://travis-ci.org/gost-c/enr.svg?branch=master)](https://travis-ci.org/gost-c/enr)

> run cmd with envs

## Usage

[Download](https://github.com/gost-c/enr/releases) the package and put in any `$PATH` folder.

### before
```bash
# touch test.sh
$ echo -e "echo \$HELLO\necho \$WORLD" > test.sh
$ HELLO=test WORLD=bar bash test.sh
# output:
#   test
#   bar
```
### after

`enr` will load file `.env` in `pwd` folder as envs. You can see [.env.example](./.env.example)

```bash
# touch test.sh
$ echo -e "echo \$HELLO\necho \$WORLD" > test.sh
# touch a env folder with our envs
$ echo -e "HELLO=test\nWORLD=bar" > .env
# run cmd start with `enr`
$ enr bash test.sh
# output:
#   test
#   bar
```

#### use custom config

```bash
$ enr -c="yourCustomConfig" bash test.sh
```

## License

MIT &copy; zcong1993

# enr

[![Go Report Card](https://goreportcard.com/badge/github.com/gost-c/enr)](https://goreportcard.com/report/github.com/gost-c/enr)
[![Build Status](https://travis-ci.org/gost-c/enr.svg?branch=master)](https://travis-ci.org/gost-c/enr)

> run cmd with envs

## Usage

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

## License

MIT &copy; zcong1993

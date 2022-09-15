[![Release](https://github.com/sirateek/yaml-env-replacer/actions/workflows/go-auto-build-release.yml/badge.svg)](https://github.com/sirateek/yaml-env-replacer/actions/workflows/go-auto-build-release.yml)

# yaml-env-replacer

A Go application that read the yaml file and replace a variable in that file with a value from ENV

This application will look for a `${ENV_NAME}` syntax in your yaml file.
and it will replace it with a value from process env that have the same `ENV_NAME`

For example you have this config file:
```yaml
name: ${NAME}
surname: ${SUR_NAME}
```

This will look for 2 process envs, **NAME** and **SUR_NAME**.

(Let say you have these process envs `NAME=foo` `SUR_NAME=bar`)

The output would be like this
```yaml
name: foo
surname: bar
```

## Run from source

**Example Command**
```bash
go run main.go -config-file=/test.yaml -env-file=/test.env -out=/out.yaml
```

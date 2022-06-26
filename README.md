## BruteMQ / An exotic service bruteforce tool
[![](https://github.com/codexlynx/brutemq/workflows/CI/badge.svg)](https://github.com/codexlynx/brutemq/actions) [![AUR](https://img.shields.io/github/license/codexlynx/brutemq)](LICENSE) [![](https://goreportcard.com/badge/github.com/codexlynx/brutemq)](https://goreportcard.com/report/github.com/codexlynx/brutemq)

### Usage:

```
brutemq - An exotic service bruteforce tool.

Usage:
  brutemq [command]

Available Commands:
  amqp        Bruteforce AMQP service
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -d, --dictionary string   dictionary file path
  -h, --help                help for brutemq
  -t, --threads int         threads number (default 100)
  -u, --user string         username

Use "brutemq [command] --help" for more information about a command.
```

### Example:

> $ brutemq amqp -d passwords.txt -u admin -e localhost:5672/ -t 500


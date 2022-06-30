## BruteMQ / An exotic service bruteforce tool
[![](https://github.com/codexlynx/brutemq/workflows/CI/badge.svg)](https://github.com/codexlynx/brutemq/actions) [![AUR](https://img.shields.io/github/license/codexlynx/brutemq)](LICENSE) [![](https://goreportcard.com/badge/github.com/codexlynx/brutemq)](https://goreportcard.com/report/github.com/codexlynx/brutemq)

Services supported:

* HashiCorp Vault Userpass
* etcd v3
* AMQP PLAIN SASL

### Run:

You can compile the binary or run via OCI image.

#### Compile:
* Requirements:
    * A version of __Docker__ with __BuildKit__ support.
    * GNU __make__ utility.

* Procedure:
    * Run: `make`.
    * Check the correct creation of `dist` directory.

#### OCI Image:

> $ docker run ghcr.io/codexlynx/brutemq:latest

Add to your shell profile:
```
function brutemq {
    docker run ghcr.io/codexlynx/brutemq:latest $@
}
```
Or
```
alias brutemq='docker run ghcr.io/codexlynx/brutemq:latest'
```

For more details click [here](https://github.com/codexlynx/brutemq/pkgs/container/brutemq).

### Usage:

```
brutemq - An exotic service bruteforce tool

Usage:
  brutemq [command]

Available Commands:
  amqp        Bruteforce AMQP PLAIN SASL service endpoint
  completion  Generate the autocompletion script for the specified shell
  etcd        Bruteforce etcdv3 service endpoint
  help        Help about any command
  vault       Bruteforce HashiCorp Vault Userpass auth

Flags:
  -d, --dictionary string   dictionary file path
  -h, --help                help for brutemq
  -t, --threads int         threads number (default 100)

Use "brutemq [command] --help" for more information about a command.
```

#### Example:

> $ brutemq amqp -d passwords.txt -u admin -e localhost:5672/ -t 500
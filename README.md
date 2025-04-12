## BruteMQ / An exotic service bruteforce tool
[![](https://github.com/codexlynx/brutemq/workflows/CI/badge.svg)](https://github.com/codexlynx/brutemq/actions) [![AUR](https://img.shields.io/github/license/codexlynx/brutemq)](LICENSE) [![](https://goreportcard.com/badge/github.com/codexlynx/brutemq)](https://goreportcard.com/report/github.com/codexlynx/brutemq)

Services supported:

* HashiCorp Vault Userpass
* etcd v3
* AMQP PLAIN SASL

### History:
Brutemq was a tool initially designed to brute force rabbitmq (amqp) queue messaging services. 
After implementing the first working version, new needs arose and etcd and HashiCorp Vault Userpass services were added.

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

> $ docker run --net=host ghcr.io/codexlynx/brutemq:latest

Add to your shell profile:
```
function brutemq {
    docker run --net=host ghcr.io/codexlynx/brutemq:latest $@
}
```
Or
```
alias brutemq='docker run --net=host ghcr.io/codexlynx/brutemq:latest'
```

For more details click [here](https://github.com/codexlynx/brutemq/pkgs/container/brutemq).

### Webhook:
Set `WEBHOOK_URL` environment variable to send an http webhook request when the password is discovered.

### Kubernetes:
You can launch brutemq on a Kubernetes cluster for various reasons, either because you can't set up port-forwarding and 
want to attack an endpoint on one of the cluster's internal networks (lack of permissions in RBAC or other limitations) 
or simply because you want to manage the workload on your own cluster. Check the `deployments` directory and set the 
[manifest](deployments/kubernetes_job.yaml) to your needs.

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

### Citations and Acknowledgments

This project has been cited in the following scholarly article:

(2025) [Signature-Based Security Analysis and Detection of IoT Threats in Advanced Message Queuing Protocol](https://www.researchgate.net/publication/389061257_Signature-Based_Security_Analysis_and_Detection_of_IoT_Threats_in_Advanced_Message_Queuing_Protocol).

> Mohammad Emran Hashimyar, Mahdi Aiash, Ali Khoshkholghi and Giacomo Nalli. __Middlesex University, London__
```bibtex
@article{article,
  author = {Hashimyar, Mohammad and Aiash, Mahdi and Khoshkholghi, Ali and Nalli, Giacomo},
  year = {2025},
  month = {02},
  pages = {5},
  title = {Signature-Based Security Analysis and Detection of IoT Threats in Advanced Message Queuing Protocol},
  volume = {5},
  journal = {Network},
  doi = {10.3390/network5010005}
}
```

We appreciate the citation and use of this project.

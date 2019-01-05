## brutemq
--------
High performance RabbitMQ (amqp) Brute Force tool.

### Usage:

```
brutemq [flags]

Flags:
  -f, --file string   load several passwords from file (default "passwords.txt")
  -h, --help          help for brutemq
  -t, --threads int   number of threads (default 200)
      --url string    rabbitmq connection URL (default "localhost:5672/vhost")
  -u, --user string   username (default "guest")
```

### Example:

> $ brutemq -t 400 -u rabbit -f rockyou.txt --url mq.target.net:5672/

### About
This tool was created by: __@codexlynx__.

* Twitter: [https://twitter.com/codexlynx](https://twitter.com/codexlynx)
* GitHub: [https://github.com/codexlynx](https://github.com/codexlynx)

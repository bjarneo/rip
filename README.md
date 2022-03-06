# RIP

This is a HTTP load testing tool that run requests concurrently. Written as a Golang learning project.

![RIP](./rip.png)

## Features

-   Run requests concurrently
-   Set a timer in second for how long it should run
-   Outputs table of statistics for the end result
-   Log the requests to $HOME/rip.log
-   Supports multiple hosts
-   UDP flood attack

## Coming

-   JSON output of the result

## Usage

Install the binary from <https://github.com/bjarneo/rip/releases>, or go directly to the build the binary manually step.

```bash
# Standard by using one host
rip -c 100 -t 10 https://your.domain.com

# Multiple hosts
touch hosts.txt

# Add the content, important that each host is on a newline
http://localhost:5000
http://localhost:5000/dis-is-nice
http://localhost:5000/yas

# RIP
rip -t 10 -h hosts.txt

# Using UDP flood attack
rip -t 10 -c 10 -u -ub 4096 0.0.0.0:30000
```

### The default values

```
Usage of rip
  -t int
    How many seconds to run the test (default: 60)
  -c float
    How many concurrent users to simulate (default: 10)
  -l bool
    Log the requests to $HOME/rip.log (default: false)
  -h string
    A file of hosts. Each host should be on a new line. Will randomly choose an host.
  -u bool
    Run requests UDP flood attack and not http requests (default: false)
  -ub int
    Set the x bytes for the UDP flood attack (default: 2048)

```

## Get it up and running [DEV]

```bash
# Install dependencies
go install

# By using the go binary directly
go run main.go
```

## Build the binary manually

```bash
# Build binary
go build

# Now it will be available as "rip"
rip http://localhost:1337
```

## Troubleshooting

If you get this error message `socket: too many open files`, you might want to increase your ulimit to a higher number.

```bash
ulimit -n 12000
```

## Disclaimer

Use this tool at your own risk

## LICENSE

See [LICENSE](./LICENSE)

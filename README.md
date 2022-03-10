# RIP

This is a HTTP load testing and UDP flood attack tool that run requests concurrently.

Note: I am using this project as a Go learning project. Refactors will most likely happen often.

Looking for new features? Create an issue.

![RIP](./rip.gif)

## Features

- HTTP load testing
- UDP flood attack
- Run requests concurrently
- Set a timer in second for how long it should run
- Outputs table of statistics for the end result
- Log the requests to $HOME/rip.log
- Supports multiple hosts

## Coming

- POST/PUT/PATCH http requests
- HTTP/UDP payload attachment
- Custom HTTP headers
- JSON output of the result

## Usage

Install the binary from <https://github.com/bjarneo/rip/releases>, or go directly to the build the binary manually step.

```bash
# Standard by using one host
rip -concurrent 100 -interval 10 https://your.domain.com

# Multiple hosts
touch hosts.txt

# Add the content, important that each host is on a newline
http://localhost:5000
http://localhost:5000/dis-is-nice
http://localhost:5000/yas

# RIP
rip -interval 10 -hosts hosts.txt

# Using UDP flood attack
rip -interval 10 -concurrent 10 -udp -udp-bytes 4096 0.0.0.0:30000
```

### The default values

```bash
Usage of rip
  -interval     int
    How many seconds to run the load tests (default: 60)
  -concurrent   int
    How many concurrent users to simulate (default: 10)
  -logger       bool
    Log the requests to $HOME/rip.log (default: false)
  -hosts        string
    A file of hosts. Each host should be on a new line. Will randomly choose an host.
  -udp          bool
    Run requests UDP flood attack and not http requests (default: false)
  -udp-bytes    int
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

Use this tool at your own risk. The owner of this repository is not responsible for its usage.

## LICENSE

See [LICENSE](./LICENSE)

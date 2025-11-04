# IPex

A fast and simple string parser that returns an IP address or a list of them. Accepts CIDR notation or IP fragments. If base IP is not set, unlike in the examples below, the IP of the machine which runs the script is used.

```go mdox-exec="sh/demo.sh"
# ipex -b 127.0.0.1 22
127.0.0.22

# ipex -b 192.168.166.33 22 33 44
192.168.166.22
192.168.166.33
192.168.166.44

# ipex -b 255.255.166.33 155.22/30 99 200/30
255.255.155.21
255.255.155.22
255.255.166.99
255.255.166.201
255.255.166.202

# ipex -b 255.255.166.33 254+4 1-4 128/31
255.255.166.254
255.255.166.255
255.255.167.0
255.255.167.1
255.255.166.1
255.255.166.0
255.255.165.255
255.255.165.254
255.255.166.128
255.255.166.129

# ipex -s -u -b 255.255.166.33 99+5 99+10 90+12
255.255.166.90
255.255.166.91
255.255.166.92
255.255.166.93
255.255.166.94
255.255.166.95
255.255.166.96
255.255.166.97
255.255.166.98
255.255.166.99
255.255.166.100
255.255.166.101
255.255.166.102
255.255.166.103
255.255.166.104
255.255.166.105
255.255.166.106
255.255.166.107
255.255.166.108

```

```go mdox-exec="r -h"
Usage: ipex [<input> ...] [flags]

ip expander, transform strings into ip addresses

Arguments:
  [<input> ...]    input strings, i.e. 192.168.33.1, 192.168.33.1/29, 33.1/30,
                   1/28, 1

Flags:
  -h, --help              Show context-sensitive help.
  -b, --base-ip=STRING    base ip used for expanding input strings, default is
                          self address
  -s, --sort              sort final list before output
  -u, --uniq              uniq, remove duplicates from the output
  -V, --version-flag      display version
```

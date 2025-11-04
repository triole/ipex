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

```

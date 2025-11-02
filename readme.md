# IPex

A fast and simple string parser that returns an IP address or a list of them. Accepts CIDR notation or IP fragments. If base IP is not set, unlike in the examples below, the IP of the machine which runs the script is used.

```go mdox-exec="sh/demo.sh"
# ipex -b 255.255.166.33 22
255.255.166.22

# ipex -b 255.255.166.33 22 33 44
255.255.166.22
255.255.166.33
255.255.166.44

# ipex -b 255.255.166.33 155.22/30 99 200/30
255.255.155.21
255.255.155.22
255.255.166.99
255.255.166.201
255.255.166.202

```

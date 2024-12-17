# Swiftie DNS

Serving Taylor Swift lyrics over DNS

# Usage

On Linux (similar on other platforms), install Golang.

To start the server,

```bash
go run main.go
```

Then, to query the server, we run,

```bash
dig "Blank Space" +short -p 8053 @127.0.0.1
```

# Why?

DNS is usually used for obtaining IP addresses and other records. It contains several types of records including A, AAAA, TXT, etc.
Some tinkers have abused this protocol to send random data using TXT records. And I aspire to be one of them.

<h2 align="center">pgo - TCP ports scanner</h2>

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![GitHub release](https://img.shields.io/github/v/release/0xWerz/pgo)](https://GitHub.com/0xwerz/pgo/releases/)
[![GitHub forks](https://badgen.net/github/forks/0xwerz/pgo/)](https://GitHub.com/0xwerz/pgo/network/)
[![GitHub forks](https://badgen.net/github/stars/0xwerz/pgo)](https://GitHub.com/0xwerz/pgo/network/)

pgo is a fast tool designed to scan tcp ports more simply.


## ⬇️ Installation

```ruby
$ git clone https://github.com/0xwerz/pgo.git
$ cd pgo
$ sudo go build main.go -o pgo && cp pgo /usr/bin 
$ pgo --help
```


## 📈 Usage
```ruby
$ pgo --help
 -h string
   	Host ip/hostname
 -r1 int
   	port range start (default 22)
 -r2 int
   	port range end (default 161)
```

### Example
```ruby
$ pgo -h github.com -r1 20 -r2 161
```

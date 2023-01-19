<h2 align="center">pgo - TCP - UDP ports scanner</h2>
<p align=center>
<a href="https://go.dev/"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
<a href="https://GitHub.com/0xwerz/pgo/releases/"><img src="https://img.shields.io/github/v/release/0xWerz/pgo">
<a href="https://github.com/0xwerz/pgo/network/"><img src="https://badgen.net/github/stars/0xwerz/pgo">
<a href="https://github.com/0xwerz/pgo/network/)"><img src="https://badgen.net/github/forks/0xwerz/pgo/">
</a>
</p>


## ⬇️ Installation

```ruby
$ git clone https://github.com/0xwerz/pgo.git
$ cd pgo
$ sudo go build main.go -o pgo && cp pgo /usr/bin 
$ pgo --help
```


## 📈 Usage
```ruby
$ pgo -h
Usage of ./main:
  -h    Show help menu
  -ip string
        IP address to scan (default "127.0.0.1")
  -p string
        Range of ports to scan (e.g. 22-80) (default "1-65535")
  -type string
        Type of scan (tcp or udp) (default "tcp")
```

### Examples

TCP
```ruby
$ pgo -ip 192.168.1.1 -p 22-80 -t tcp
```

UDP
```ruby
$ pgo -ip 192.168.1.1 -p 22-80 -t udp
```
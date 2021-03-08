# registrobr-checker
> :mag_right: A simple command line tool to check domain avaiability

## Why ?
    
Easy way to validate if domain is avaiable for usage in [registro.br](https://registro.br)

## How to install

```
go get github.com/lucasmendesl/registrobr-checker/cmd
```
or build source code

```sh
$ git clone https://github.com/lucasmendesl/registrobr-checker.git
$ cd registrobr-checker
$ go build -o registrobr-checker ./cmd
```

## Usage

```sh
$ registrobr-checker -help #prints cli usage
$ registrobr -hostname example.com.br #validate domain availability
```

### Examples:

Unavaiable domain:

```sh
$ registrobr-checker -hostname pudim.com.br
Result: the domain pudim.com.br is not avaiable
Reason: domain published
Expiration: 24/02/2028
Suggestions:
        - pudim.agr.br
        - pudim.blog.br
        - pudim.dev.br
        - pudim.eco.br
        - pudim.esp.br
        - pudim.etc.br
        - pudim.far.br
        - pudim.flog.br
        - pudim.imb.br
        - pudim.ind.br
        - pudim.inf.br
        - pudim.log.br
        - pudim.net.br
        - pudim.ong.br
        - pudim.rec.br
        - pudim.seg.br
        - pudim.srv.br
        - pudim.tec.br
        - pudim.tmp.br
        - pudim.tur.br
        - pudim.tv.br
        - pudim.vlog.br
        - pudim.wiki.br
```

Avaiable domain:

```sh
$ registrobr-checker -hostname avaiabledomain.com.br
the domain avaiabledomain.com.br is avaiable
```

## Contributing
Contributions via pull requests are welcome :-).

## License

MIT © [Lucas Mendes Loureiro](http://github.com/lucasmendesl)

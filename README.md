# Whoxyrm
![Go Build](https://github.com/milindpurswani/whoxyrm/actions/workflows/go.yml/badge.svg) ![Go Release](https://github.com/milindpurswani/whoxyrm/actions/workflows/release.yml/badge.svg)


This is a tool that lets you query whoxy api to look for reverse who is domain names based on your search criteria. It was built upon Whoxy API based on [@jhaddix](https://twitter.com/Jhaddix)'s talk on [Bug Hunter's Methodology v4.02](https://www.youtube.com/watch?v=gIz_yn0Uvb8). 
## Usage

```
$ whoxyrm -company-name "Oath Inc."
```

or 

```
$ whoxyrm -name "Oath Inc."
```

or 
```
$ whoxyrm -email "test@example.com"
```

or

```
$ whoxyrm -keyword "yahoo.com"
```

## Installation

```
$ go install github.com/milindpurswani/whoxyrm@latest
```

Also, Make sure you export your whoxy api key as follows:
*without this, it won't work.*

```
$ export WHOXY_API_KEY="..."
```
*you can grab one from https://www.whoxy.com*


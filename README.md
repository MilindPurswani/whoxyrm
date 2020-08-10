# Whoxyrm
A reverse whois tool based on Whoxy API based on [@jhaddix](https://twitter.com/Jhaddix)'s talk on [Bug Hunter's Methodology v4.02](https://www.youtube.com/watch?v=gIz_yn0Uvb8). 

## Usage

```
$ whoxyrm -company-name "Oath Inc."
```

## Installation

```
go get -u github.com/milindpurswani/whoxyrm
```

Also, Make sure you export your whoxy api key as follows:
*without this, it won't work.*

```
export WHOXY_API_KEY="..."
```
*you can grab one from https://www.whoxy.com*


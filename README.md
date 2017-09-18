### INTRO

This is a simple goa project that shows the [new goa gzip middleware](https://github.com/goadesign/goa/pull/1368) is broken.

this test codes exposes two resources /good & /bad for simplicity both of them returns generated swagger code.

* /bad uses the new middleware
* /good uses the old middleware. to make it available to this test, i just copied the gzip middelgare before the pr was merged

### BUILD

```
go get
go build
```

### REPRO
#### in the good resource 

```
curl -v 'http://localhost:8080/good/swagger.json' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8' -H 'Connection: keep-alive' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,es-AR;q=0.8,es;q=0.6,en;q=0.4' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36' --compressed
```

#### in the bad resource 

```
curl -v 'http://localhost:8080/bad/swagger.json' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8' -H 'Connection: keep-alive' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,es-AR;q=0.8,es;q=0.6,en;q=0.4' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36' --compressed
```


i hope this helps.
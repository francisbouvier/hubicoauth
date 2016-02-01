# HubiC oAuth webserver

A very simple web server in Go to handle HubiC oAuth process.

## Installation

**Standalone**

Just download the executable and launch it.

```sh
user@host:~ $ ./hubicoauth
Running server on :8085 ...

# -p to define port
# hubicoauth --help for more info
```

**With Docker**

```sh
sudo docker run -dt --name hubicoauth -p $HOST_PORT:8085 francisbouvier/hubicoauth
```

## Proxy configuration

This web server is meant to be used behind a reverse proxy, like *Nginx* or *Apache*.
The reverse proxy will have to handle SSL as HubiC requires https for the redirection uri.

Here is a configuration example for *Nginx*:

```nginx
server {
    listen 80;
    server_name $HOST;

    location / {
        return 301 https://$server_name$request_uri;
    }
}

server {
    listen 443 ssl;
    server_name $HOST;

    ssl_certificate fullchain.pem;
    ssl_certificate_key privkey.pem;

    location / {
        proxy_pass http://localhost:$HOST_PORT/;
        proxy_set_header Host $host;
    }
}
```

Clipboard Syncer
=======

Sync Clipboard data between Linux and macOS


## Requirement

* docker
* docker-compose
* system: Linux or macOS

## Setup Server 

create a file named `.env`,

```text
PORT=9000
SIGN_SECRET=5S4imObpqqhol1wt
ALLOW_TOKEN=5S4imObpqqhol1wt
CHECK_SIGN=true
```

then, 

```bash
docker-compose up -d --build
```

## Setup Client

build client command line from source

```bash
git clone https://github.com/juxuny/clipboard-sync.git
cd clipboard-sync
export GOPROXY=https://goproxy.cn
go mod download && go install github.com/juxuny/clipboard-sync/cmd/syncer
```

start client daemon:
```bash
syncer run --host=http://127.0.0.1:9000 -t=5S4imObpqqhol1wt -s=5S4imObpqqhol1wt -d=2
```

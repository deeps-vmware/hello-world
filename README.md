# hello-world

[![Docker Build Status](https://img.shields.io/docker/cloud/build/deepsvmwarecom/hello-world.svg)](https://hub.docker.com/r/deepsvmwarecom/hello-world/) [![Docker Pulls](https://img.shields.io/docker/pulls/deepsvmwarecom/hello-world.svg)](https://hub.docker.com/r/deepsvmwarecom/hello-world/)

Simple http server responding "Hello World! from {NODE_ID}:{PORT}"

    > docker run --rm -p 8090:8090 deepsvmwarecom/hello-world
    Hello World! from d0bb7545993f 172.17.0.2:8090
    {"ip":"172.17.0.2","level":"info","msg":"Hello World!","nodeID":"d0bb7545993f","port":"8090","time":"2020-02-22T14:54:57Z","timeout":"5","upstream":""}

Now you can hit the sever on localhost

    > curl localhost:8090
    Hello World! from d0bb7545993f:8090 (1)

## defaults

| ENV       | default value |
|-----------|---------------|
| PORT      | 8090          |
| NODE_ID   | hostname      |
| UPSTREAM  | nil           |
| TIMEOUT   | 5 (seconds)   |

You can override these as follows

    > docker run --rm -p 8090:8888 -e PORT=8888 -e NODE_ID=my-server-01 deepsvmwarecom/hello-world
    Hello World! from my-server-01 172.17.0.2:8888
    {"ip":"172.17.0.2","level":"info","msg":"Hello World!","nodeID":"my-server-01","port":"8888","time":"2020-02-22T15:00:05Z","timeout":"5","upstream":""}

    > curl localhost:8090
    Hello World! from my-server-01:8888 (1)

## upstream

An upstream http GET call is made if env `UPSTREAM` is provided

    > docker run --rm -p 8091:8888 -e PORT=8888 -e NODE_ID=my-server-02 -e UPSTREAM=http://172.17.0.2:8888 deepsvmwarecom/hello-world
    Hello World! from my-server-02 172.17.0.3:8888
    Upstream: http://172.17.0.2:8888 Timeout: 5s
    {"ip":"172.17.0.3","level":"info","msg":"Hello World!","nodeID":"my-server-02","port":"8888","time":"2020-02-22T15:00:37Z","timeout":"5","upstream":"http://172.17.0.2:8888"}

    > curl localhost:8091
    Hello World! from my-server-02:8888 (1)
    Upstream: Hello World! from my-server-01:8888 (2)
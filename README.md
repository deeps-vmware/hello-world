# hello-world

Simple http server responding "Hello World! from {NODE_ID}:{PORT}"

    > docker run --rm -p 8090:8090 deepsvmwarecom/hello-world
    Hello World! from 172.17.0.2:8090

Now you can hit the sever on localhost

    > curl localhost:8090
    Hello World! from 172.17.0.2:8090 (1)

## defaults

| ENV       | default value |
|-----------|---------------|
| PORT      | 8090          |
| NODE_ID   | public ip     |

You can override these as follows

    > docker run --rm -p 8080:8888 -e PORT=8888 -e NODE_ID=my-server-01 deepsvmwarecom/hello-world
    Hello World! from my-server-01:8888

    > curl localhost:8080
    Hello World! from my-server-01:8888 (1)

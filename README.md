# hello-world

Simple http server responding "Hello World! from {NODE_ID}:{PORT}"

    docker run --rm -p 8090:8090 -e PORT=8090 deepsvmwarecom/hello-world

This will result in 

    Hello World! from 172.17.0.2:8090

## Defaults

| ENV       | default value |
|-----------|---------------|
| PORT      | 8090          |
| NODE_ID   | public ip     |

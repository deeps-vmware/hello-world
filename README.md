# hello-world

Simple http server responding "Hello World! from {NODE_IP}:{PORT}"

    docker run --rm -p 8090:8090 -e PORT=8090 -e NODE_IP=10.0.0.11 deepsvmwarecom/hello-world

This will result in 

    Hello World! from 10.0.0.11:8090

## Defaults

| ENV       | default value |
|-----------|---------------|
| PORT      | 8090          |
| NODE_IP   | public ip     |
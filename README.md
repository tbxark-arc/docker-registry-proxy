# docker-registry-mirror

This is a simple Docker registry mirror.

### Usage
```
Docker Registry Proxy 
  -address string
        Serve Address (default "localhost:8989")
  -help
        Show help
  -registry string
        Docker Registry Host (default "registry-1.docker.io")
```

### Example
```bash
./docker-registry-mirror -registry mirror.ccs.tencentyun.com
```

### License

**docker-registry-mirror** is released under the MIT license. See LICENSE for details.
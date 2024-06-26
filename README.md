# docker-registry-proxy


~~This is a simple Docker registry proxy. Due to the fact that Docker Registry mirror service provided by Chinese cloud service providers is only open to intranet, a proxy service is needed for access. This service is a simple proxy service that forwards requests to the specified Docker Registry.~~

### Nginx Only

```

server {
    server_name hub.example.cn;

    client_max_body_size 0;
    chunked_transfer_encoding on;

    location / {
        proxy_pass https://mirror.ccs.tencentyun.com;  
        proxy_set_header Host mirror.ccs.tencentyun.com;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;               
        proxy_buffering off;
        proxy_set_header Authorization $http_authorization;
        proxy_pass_header  Authorization;
        proxy_intercept_errors on;
        recursive_error_pages on;
        error_page 301 302 307 = @handle_redirect;
    }
    
    location @handle_redirect {
        resolver 183.60.83.19;
        set $saved_redirect_location '$upstream_http_location';
        proxy_pass $saved_redirect_location;
    }

    listen 443 ssl;
    ssl_certificate /path/to/fullchain.pem;
    ssl_certificate_key /path/to/privkey.pem;
}

```

<details>
<summary>Deprecated</summary>
<pre>
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
./docker-registry-proxy -registry mirror.ccs.tencentyun.com
```
</pre>
</details>


### License

**docker-registry-proxy** is released under the MIT license. See [LICENSE](LICENSE) for details.
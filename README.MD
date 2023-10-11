# API Files


## How to use

### Use Taskfile

> - rename `Taskfile.yaml-example` to `Taskfile.yaml`
>
> - change `vars.DOCKER_REGISTRY_URL` to your docker registry url
>
### Run help optional:
```bash
go run main.go serv -h
```
```raw
start fiber app

Usage:
  app_api serv [flags]

Flags:
  -b, --browse           Browse (default true)
  -d, --dirpath string   DirPath (default "./assets")
  -h, --help             help for serv
  -H, --host string      Host (default "0.0.0.0")
  -a, --max-age int      Cashe Max-Age (default 3600)
  -P, --port int         LitenerPort (default 80)
  -p, --prefix string    URL PathPrefix (default "/")
```

### Test run
```bash
task dev
```


### build executable
```bash
task build
```


### Docker Build
```bash
task docker-build
```

### Docker Push
```bash
task docker-push
```


## Docker Compose file
```yaml
version: "3.8"
services: 
  apifiles:
      image: attap0n/apifiles:latest
      ports:
        - 18080:8080
      environment:
        TZ: Asia/Bangkok
      # 
      command: /app/apifiles serv --host "0.0.0.0" --port 8080 --prefix "/assets" --dirpath "/data/assets"
      volumes:
        - /data/assets
volumes:
  file-data:
``` 


## Docker Stack file and Traefik proxy
```yaml
# ENV required
# --------------------------------------------
# SERVICE_NAME=filebrowser1234
# DOMAIN=localhost
# BASE_URL=/filebrowser4321
# VOLUME_MOUNT=/home/user
# PUID=1000
# PGID=1000
# --------------------------------------------
# END ENV
version: "3.8"
services:
  filebrowser:
    image: filebrowser/filebrowser:latest
    environment:
      TZ: Asia/Bangkok
      PUID: ${PUID}
      PGID: ${PGID}
      # user: admin:admin
    command: --baseurl "${BASE_URL:-/files}"
    deploy:
      mode: replicated
      replicas: 1
      labels:
        - "traefik.enable=true"
        - "traefik.http.routers.${SERVICE_NAME}.tls=true"
        - "traefik.http.routers.${SERVICE_NAME}.entrypoints=websecure"
        - "traefik.http.routers.${SERVICE_NAME}.rule=Host(`${DOMAIN}`) && PathPrefix(`${BASE_URL}`)"
        - "traefik.http.services.${SERVICE_NAME}.loadbalancer.server.port=80"
    networks:
      - proxy
    volumes:
      - database:/database/filebrowser.db:z
      - ${VOLUME_MOUNT}:/srv

networks:
  proxy:
    external: true

volumes:
  database:
```
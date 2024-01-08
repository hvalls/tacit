# Tacit

Tacit allows you to run an HTTP server using bash scripts handlers.

### Config file
Create a YAML config file with the structure below:

```yaml
---
endpoints:
  - name: # meaningful endpoint name (e.g. Get log files)
    method: # only "GET" supported so far
    path: # e.g. "/logs"
    handler: # path to script to be executed (e.g. "./get_logs.sh")
    args: [] #arguments to be passed to the script (sorted). You can use "$query." to get query params (e.g. $query.limit)
```

You can add as many endpoints as you need. 

### Run server

```
$ tacit -f {config_file} -p {server_port}
```

- Default `config_file` is `./config.yml`
- Default `server_port` is `8080`


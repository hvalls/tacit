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
$ PORT=9090 tacit -f {config_file}
```

- Default `config_file` is `./config.yml`
- Default port is 8080

### Example

```
$ cd example/
$ tacit 
Registering endpoint:  {Get log files GET /logs ./get_logs.sh [$query.limit]}
Ready. Tacit server is listening on port 8080
```

```
$ curl "http://localhost:8080/logs?limit=3"
{ "data": 
  { "logfiles": [
    "1.txt",
    "2.txt",
    "3.txt"
  ] }
}
```
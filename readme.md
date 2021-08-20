## run opensearch + dashboard + logstash
```
docker compose up -d
```


## run the container with gelf logging driver along with logstash address
```
docker run --rm --log-driver gelf --log-opt gelf-address=udp://localhost:12201 --log-opt tag=testcontainer --log-opt tag=testcontainer test:test
```
* check --log-opt labels=production_status,geo

## import tests dashboard
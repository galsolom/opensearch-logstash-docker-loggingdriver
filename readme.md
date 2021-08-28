## local opensearch stack
----
### Run opensearch + dashboard + logstash
cmd:
```
docker compose up -d
```
### import tests dashboard

* after dashboard is up @ `localhost:5601`, import the dashboard  `Tests Explorer`  using the menu in Stack Management -> Saved Objects -> import -> choose `TestsExplorer-dashboard.ndjson` from this repo dir

###

### build and  run the test container with the logging driver
```
cd testproject
docker build .  -t test:test
docker run --rm --log-driver gelf --log-opt gelf-address=udp://localhost:12201 --log-opt tag=qa  test:test
```
* check --log-opt labels=production_status,geo
pattern => "(^.+Failed.+)|(^.+Error.+)|(^.+System.+)|(^.+Stack.+)|(^.+at.+)|(^s*End of:.+)"
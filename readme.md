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
docker run --log-driver=fluentd --log-opt fluentd-address=localhost:24224 --log-opt tag=qa  test:test
```
* check --log-opt labels=production_status,geo
pattern => "(^.+Failed.+)|(^.+Error.+)|(^.+System.+)|(^.+Stack.+)|(^.+at.+)|(^s*End of:.+)"



## FLUENTD
docker run --log-driver=fluentd --log-opt fluentd-address=localhost:24224




# simple poc for testing..
requirements, python3, dotnet
DB is labels.json
to run:
cmd:
```
# will run only test case 1337
python3 labels.py identity

# will run only test case 1338
python3 labels.py submission

# will run both
python3 labels.py entities
```
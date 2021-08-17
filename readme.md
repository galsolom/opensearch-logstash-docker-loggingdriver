## run opensearch + dashboard
```
docker compose up -d
```

## run logstash through cygwin terminal (gitbash):
 * running through gitbash for multi line commands such as below
 ```
docker run -d --rm --name logstash -p 12201:12201/udp --network=opensearch-logstash-plugin_opensearch-net opensearchproject/logstash-oss-with-opensearch-output-plugin:7.13.2 -e 'input {
    gelf {
        port => 12201
      }
}
output	{
   opensearch {
     hosts => ["opensearch-node1:9200"]
     index => "opensearch-logstash-docker-%{+YYYY.MM.dd}"
     user => "admin"
     password => "admin"
     ssl_certificate_verification => false
     ssl => true
   }
 }'
 ```


## for some reason network service name does is not resolved. either host ip or container ip, (kuberentes would allow loadbalancer and ingress service)
> probably because the `logstash` service is resolved only within the network, and docker runs it outside the network.
```
docker run --network=opensearch-logstash-plugin_opensearch-net --log-driver gelf --log-opt gelf-address=udp://localhost:12201 --log-opt tag=testcontainer test:test
```
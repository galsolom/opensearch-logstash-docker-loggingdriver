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
 logging:
      driver: gelf
      options:
        gelf-address: "udp://localhost:12201"
        tag: "demo2_app"
        docker run --network=opensearch-logstash-plugin_opensearch-net --log-driver gelf --log-opt gelf-address="udp://logstash:12201" --log-opt tag=testcontainer test:test
          splunk 

 curl -XPOST http://localhost:9200/_cluster/health?pretty=true -H 'Content-Type: application/json' -d '{"enabled": true}'


   docker run --network opensearch-logstash-plugin_opensearch-net --log-driver gelf --log-opt gelf-address="tcp://localhost:12201" --log-opt tag=testcontainer test:test
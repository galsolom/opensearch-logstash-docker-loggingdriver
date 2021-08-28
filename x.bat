echo off
set arg1=%1
docker run --name %arg1% --rm --log-driver gelf --log-opt gelf-address=udp://localhost:12201 --log-opt tag=dev test:test
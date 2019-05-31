docker run --name es --network dfs --rm  -it --network-alias es  ubuntu-golang:es-run

docker run --name mq --network dfs --rm  -d --network-alias mq  ubuntu-golang:mq

docker run --name api1 --network dfs --rm -d -p 12345:12345 --network-alias api1 ubuntu-golang:api1
docker run --name api2 --network dfs --rm -d -p 12346:12346 --network-alias api2 ubuntu-golang:api2

docker run --name data1 --network dfs --rm -d --network-alias data1 ubuntu-golang:data
docker run --name data2 --network dfs --rm -d --network-alias data2 ubuntu-golang:data
docker run --name data3 --network dfs --rm -d --network-alias data3 ubuntu-golang:data
docker run --name data4 --network dfs --rm -d --network-alias data4 ubuntu-golang:data
docker run --name data5 --network dfs --rm -d --network-alias data5 ubuntu-golang:data
docker run --name data6 --network dfs --rm -d --network-alias data6 ubuntu-golang:data
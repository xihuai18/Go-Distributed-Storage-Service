# enable and start the elasticsearch service
sudo systemctl enable elasticsearch.service
sudo systemctl start elasticsearch.service

# configure the storage environments
for i in `seq 1 6`
do
    mkdir -p /tmp/$i/objects
    mkdir -p /tmp/$i/temp
    mkdir -p /tmp/$i/garbage
done


# configure the ip address, which will be reset after rebooting
# 6 dataservers
sudo ifconfig wlp3s0:1 127.0.1.1
sudo ifconfig wlp3s0:2 127.0.1.2
sudo ifconfig wlp3s0:3 127.0.1.3
sudo ifconfig wlp3s0:4 127.0.1.4
sudo ifconfig wlp3s0:5 127.0.1.5
sudo ifconfig wlp3s0:6 127.0.1.6
# 2 apiservers
sudo ifconfig wlp3s0:7 127.0.2.1
sudo ifconfig wlp3s0:8 127.0.2.2

# setup metadata server
curl localhost:9200/metadata -XPUT -H "content-type: application/JSON" -d @mapping.json


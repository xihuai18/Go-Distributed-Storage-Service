#!/bin/bash

for i in `seq 1 6`
do
    rm -rf /tmp/$i/objects/*
    rm -rf /tmp/$i/temp/*
done

curl 127.0.0.1:9200/metadata/_delete_by_query -XPOST metadata/_delete_by_query -H "content-type: application/JSON" -d'
{
  "query": { 
    "match_all": {}
  }
}
'
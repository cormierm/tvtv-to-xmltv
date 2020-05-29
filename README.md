# TvtvToXMLTV

Creates http endpoint that grabs tv guide listing from tvtv.ca and returns xml in xmltv format.

### Run in docker
Run the following commands:

```
docker build --tag tvtv-to-xmltv:1.0 .
docker run --publish 6060:8080 --detach --name TvtvToXmlTV tvtv-to-xmltv:1.0
```

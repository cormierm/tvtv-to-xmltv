# TvtvToXmlTV

Creates http endpoint that grabs tv guide listing from tvtv.ca and returns xml in xmltv format.

### Example

http://[host ip:6060]?days=7&location=3003

This will get 7 days of tv guide from tvtv.ca for location 3003 (Broadcast - Kitchener).

### Run in docker
Run the following commands:

```
docker build --tag tvtv-to-xmltv:1.0 .
docker run --publish 8080:8080 --detach --name tvtv-to-xmltv tvtv-to-xmltv:1.0
```

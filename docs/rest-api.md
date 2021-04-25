- CRUD
``` 
[GIN-debug] POST   /api/kv/:b/:k             --> github.com/yusys-cloud/ai-tools/server.(*Server).create-fm (3 handlers)
[GIN-debug] GET    /api/kv/:b/:k             --> github.com/yusys-cloud/ai-tools/server.(*Server).readAll-fm (3 handlers)
[GIN-debug] GET    /api/kv/:b/:k/:kid        --> github.com/yusys-cloud/ai-tools/server.(*Server).readOne-fm (3 handlers)
[GIN-debug] PUT    /api/kv/:b/:k/:kid        --> github.com/yusys-cloud/ai-tools/server.(*Server).update-fm (3 handlers)
[GIN-debug] DELETE /api/kv/:b/:k/:kid        --> github.com/yusys-cloud/ai-tools/server.(*Server).delete-fm (3 handlers)    

```
- Http
``` 
POST   /api/http/do  

curl --location --request POST 'http://localhost:9999/api/http/do' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "http://localhost:9999/api/kv/chaos/designer",
    "method": "post",
    "data": {
        "a": "1"
    }
}'
```
- Search
``` 
curl localhost:9999/api/search?b=snippets&k=code&key=v.name&value=linux
```
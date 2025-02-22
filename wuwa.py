import mitmproxy
from mitmproxy import ctx, http
class EndFieldModifier:
    def requestheaders(self,flow: mitmproxy.http.HTTPFlow):
        if "useTaskId" in flow.request.url:
            if flow.request.method=="CONNECT":
                return
            flow.request.scheme="http"
            flow.request.cookies.update({
                "OriginalHost":flow.request.host,
                "OriginalUrl":flow.request.url
            })
            flow.request.host="localhost"
            flow.request.port=1088
            ctx.log.info("URL:"+flow.request.url)
             
            
            
addons=[
    EndFieldModifier()
]
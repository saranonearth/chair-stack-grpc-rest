# chair-stack-grpc-rest
[![2dZHSs.md.png](https://i.ibb.co/WDZHJdp/chairstack.png)](https://i.ibb.co/WDZHJdp/chairstack.png)

This project contains implementation of REST and gRPC for a small application called ChairStack using Go Lang. 

##### To compile .proto to .go file
```

protoc services.proto --go_out=plugin=grpc .

```

##### To compile .proto to .js file
```

 protoc services.proto --js_out=import_style=commonjs:..\client\proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:..\client\proto

```

#### gRPC Web Wrapper
```
grpcServer := grpc.Server()
grpcWebServer := grpcweb.WrapServer(grpcServer)
httpServer = &http.Server{
  Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
    if r.ProtoMajor == 2 {
      grpcWebServer.ServeHTTP(w, r)
    } else { 
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
      w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
      w.Header().Set("grpc-status", "")
      w.Header().Set("grpc-message", "") 
      if grpcWebServer.IsGrpcWebRequest(r) {
        grpcWebServer.ServeHTTP(w, r) 
      }
    }
  }), &http2.Server{}),
}
```

More: https://rogchap.com/2019/07/26/in-process-grpc-web-proxy/


service=my_service
protofile=$(service).proto

all: gen gen-ruby gen-reverse-proxy gen-swagger gen-swagger-server gen-swagger-client

gen: $(service).pb.go
$(service).pb.go: $(protofile)
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		$<

gen-ruby: $(service)_pb.rb
$(service)_pb.rb: $(protofile)
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--ruby_out=. \
		$<

gen-reverse-proxy: $(service).pb.gw.go
$(service).pb.gw.go: $(protofile)
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		$<

gen-swagger: $(service).swagger.json
$(service).swagger.json: $(protofile)
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		$<

# generates restapi
gen-swagger-server: $(service).swagger.json
	swagger generate server -f $<

# generates models and client
gen-swagger-client: $(service).swagger.json
	swagger generate client -f $<

gen-swagger-js-client: $(service).swagger.json
	mkdir -p js
	swagger-codegen generate -i $< -l javascript -o js

gen-swagger-ts-fetch-client: $(service).swagger.json
	mkdir -p ts-fetch
	swagger-codegen generate -i $< -l typescript-fetch -o ts-fetch

gen-swagger-js-codegen-js-client: $(service).swagger.json
	mkdir -p swagger-js-codegen-js
	node gen.js > react/sample.js

clean:
	-rm $(service).pb.go
	-rm $(service).pb.gw.go
	-rm $(service)_pb.rb
	-rm $(service).swagger.json
	-rm -rf models client restapi js ts-fetch react

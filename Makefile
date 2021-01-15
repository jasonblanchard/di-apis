SRC=proto/notebook/v2/notebook.proto
GO_OUT=./packages/go

build_go:
	docker run -v ${shell pwd}:/defs namely/protoc-all -d notebook/v2 -l go --with-gateway

build_rest_proxy:
	docker run -v ${shell pwd}:/defs namely/gen-grpc-gateway -f notebook/v2/notebook.proto -s Notebook

build: build_go

swaggerui:
	docker run -e SWAGGER_JSON=/di-apis/gen/pb-go/notebook.swagger.json -v ${shell pwd}:/di-apis -p 9090:8080 swaggerapi/swagger-ui

clean:
	rm -rf gen

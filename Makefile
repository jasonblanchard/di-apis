SRC=proto/notebook/v2/notebook.proto
GO_OUT=./packages/go

build_go:
	docker run -v ${shell pwd}:/defs namely/protoc-all -d notebook/v2 -l go --with-gateway

build_rest_proxy:
	docker run -v ${shell pwd}:/defs namely/gen-grpc-gateway -f notebook/v2/notebook.proto -s Notebook

build: build_go

clean:
	rm -rf gen

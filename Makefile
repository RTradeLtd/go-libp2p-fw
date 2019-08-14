# uses gogofaster to get us even faster unmarshaling/marshaling speeds
.PHONY: proto-gen
proto-gen:
	# generate firewall bindings bindings
	protoc \
		-I=pb \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		pb/acl.proto \
		--gogofaster_out=plugins=grpc:pb
	# generate docs
	protoc \
		-I=pb \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--doc_out=pb \
		--doc_opt=html,index.html \
		pb/*.proto


.PHONY: deps
deps:
	go get github.com/gogo/protobuf/protoc-gen-gofast
FROM golang:1.21.3
# install unzip
RUN apt-get update
RUN apt-get install -y unzip

WORKDIR /pb

# install protoc
RUN PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
    curl -LO $PB_REL/download/v25.2/protoc-25.2-linux-x86_64.zip

RUN unzip protoc-25.2-linux-x86_64.zip -d /pb/protoc
ENV GOPROXY="https://proxy.golang.com.cn,direct"
ENV PATH="$PATH:/pb/protoc/bin"
RUN protoc --version

# install protoc-gen-go, protoc-gen-go-grpc, protoc-gen-go-http and protoc-gen-validate
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
RUN go install github.com/envoyproxy/protoc-gen-validate@latest

WORKDIR /pb/proto

# generate code
CMD find api -name *.proto | xargs -I{} sh -c 'protoc --proto_path=./api \
                                  	   					  	--proto_path=./third_party \
                                   						   	--go_out=paths=source_relative:./api \
                                   	   					   	--go-http_out=paths=source_relative:./api \
                                   	   					  	--go-grpc_out=paths=source_relative:./api \
                                   	   					  	--validate_out=paths=source_relative,lang=go:./api \
                                   	    				   	{}' && \
    find app -name *.proto | xargs -I{} sh -c 'protoc --proto_path=./app \
                                     	   					--proto_path=./third_party \
                                      						--go_out=paths=source_relative:./app \
                                      	    				{}'
FROM gitpod/workspace-full:2022-04-01-12-31-37

RUN PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
    VERSION="protoc-3.19.4-linux-x86_64.zip" && \
    curl -LO $PB_REL/download/v3.19.4/$VERSION && \
    mkdir -p $HOME/.local/protoc && \
    unzip $VERSION -d $HOME/.local/protoc && \
    rm $VERSION
    
ENV PATH=$PATH:$HOME/.local/protoc/bin
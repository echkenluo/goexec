FROM golang:1.15 as builder

# Download deps
RUN mkdir -p $GOPATH/src/fw
WORKDIR $GOPATH/src/fw

ADD go.mod go.sum ./
RUN go mod download

RUN mkdir -p /opt/fw/bin

ADD . $GOPATH/src/fw

# Build
RUN make bin

RUN cp $GOPATH/src/fw/bin/* /opt/fw/bin/
RUN chmod 755 /opt/fw/bin/*

#FROM ubuntu
FROM ubuntu:20.04

#RUN apk update && apk add openvswitch
RUN apt update && apt install -y openvswitch-switch=2.13.3-0ubuntu0.20.04.2 iproute2 iptables tcpdump && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /opt/fw/bin
COPY --from=builder /opt/fw/bin/* /opt/fw/bin/

WORKDIR /opt/fw/bin
ENV PATH=${PATH}:/opt/fw/bin

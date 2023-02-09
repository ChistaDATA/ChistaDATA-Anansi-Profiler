FROM golang:1.19 as build
RUN apt-get update -y && apt-get install -y make
WORKDIR /opt/anansi
COPY . .
RUN make build

FROM ubuntu

COPY --from=build /opt/anansi/anansi-profiler /usr/bin
WORKDIR /mnt
USER nobody
ENTRYPOINT ["/usr/bin/anansi-profiler"]

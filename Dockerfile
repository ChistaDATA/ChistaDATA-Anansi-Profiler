FROM golang:1.19 as build
RUN apt-get update -y && apt-get install -y make
WORKDIR /opt/anansi
COPY . .
RUN make build

FROM ubuntu
RUN apt update && apt install -y libssl3 && rm -rf /var/lib/apt/lists/*
COPY --from=build /opt/anansi/anansi-profiler /usr/bin
WORKDIR /mnt
USER nobody
ENTRYPOINT ["/usr/bin/anansi-profiler"]

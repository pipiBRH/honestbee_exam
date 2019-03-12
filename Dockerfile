FROM golang:1.12-stretch AS build-env
ENV GO111MODULE=on
COPY . /work
WORKDIR /work
RUN make linux
CMD ["./bin/main", "-config", "./conf/gcp.toml"]

# FROM scratch
# COPY --from=build-env /work/bin /work/bin
# COPY --from=build-env /work/conf /work/conf
# RUN mkdir /tmp
# WORKDIR /work
# EXPOSE 8080
# CMD ["./bin/main"]
FROM alpine:3.7

RUN apk update && apk upgrade && \
    apk add openrc collectd collectd-network && \
    apk add --no-cache git make musl-dev go

COPY collectd.conf /etc/collectd/collectd.conf

ENV GOROOT /usr/lib/go

COPY ./calculadora.go ${GOROOT}

RUN go get -u github.com/gorilla/mux \
&& go build ${GOROOT}/calculadora.go

RUN cd ${GOROOT}

CMD ["sh", "-c", "collectd && ./calculadora"]
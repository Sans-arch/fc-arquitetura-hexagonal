FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

#RUN go install github.com/spf13/cobra@latest && \
#    go install github.com/golang/mock/mockgen@v1.5.0 && \
#    go install github.com/spf13/cobra-cli@latest

RUN go install github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update && apt-get install sqlite3 -y

RUN mkdir -p /var/www/.cache

CMD ["tail", "-f", "/dev/null"]
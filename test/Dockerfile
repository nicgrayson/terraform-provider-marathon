FROM golang:1.7
ENV GOPATH=/go
ENV TF_LOG=TRACE
ENV TF_LOG_PATH=/go/src/github.com/Banno/terraform-provider-marathon/marathon/test-sh-tf.log
ENV TF_ACC=yes
ENV MARATHON_URL=http://marathon:8080

WORKDIR /go/src/github.com/Banno/terraform-provider-marathon/marathon/

ADD test.sh /test
CMD ["/test"]

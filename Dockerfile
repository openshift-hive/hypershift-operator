FROM openshift/origin-release:golang-1.13 as builder
RUN mkdir -p /go/src/github.com/openshift-hive/hypershift-operator
WORKDIR /go/src/github.com/openshift-hive/hypershift-operator
COPY . .
RUN go build -mod=vendor -o ./bin/hypershift-operator ./cmd/hypershift-operator/main.go

FROM registry.access.redhat.com/ubi7/ubi
RUN yum -y update && yum clean all
COPY --from=builder /go/src/github.com/openshift-hive/hypershift-operator/bin/hypershift-operator /usr/bin/hypershift-operator
ENTRYPOINT /usr/bin/hypershift-operator

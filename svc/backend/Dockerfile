FROM golang

WORKDIR /go/src/otpgo

EXPOSE 4711

COPY ./ ./

RUN go get -d -v otpgo/svc/backend
RUN go install -v otpgo/svc/backend
RUN go build otpgo/svc/backend

CMD [ "backend" ]
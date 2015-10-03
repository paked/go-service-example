FROM golang:1.4

RUN go get github.com/codegangsta/gin

RUN go get github.com/lib/pq
RUN go get gopkg.in/mgutz/dat.v1
RUN go get gopkg.in/mgutz/dat.v1/sqlx-runner
RUN go get github.com/paked/configure
RUN go get github.com/gorilla/mux
RUN go get github.com/paked/gerrycode/communicator
RUN go get github.com/paked/restrict
RUN go get github.com/dgrijalva/jwt-go

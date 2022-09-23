FROM golang:1.17.1-alpine3.14

WORKDIR /app
COPY . .

RUN apk update && \
    apk add ca-certificates wget && \
    update-ca-certificates && \
    apk add --no-cache bash make cmake

RUN apk add build-base

# Git is required for fetching the dependencies.
RUN apk add --no-cache git --update py-pip \
    && pip install setuptools \
    && pip install wheel \
    && pip install awscli --upgrade

RUN go mod download

RUN go get github.com/google/wire/cmd/wire

# Install air for Live Reload on Development environment
RUN wget -O install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh
RUN chmod +x install.sh
RUN sh install.sh
RUN cp ./bin/air /bin/air

CMD ["air", "-d"]
FROM golang:1.24-alpine as builder

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN go build -a -o /app/bin/server .

FROM ubuntu:22.04 as runner

RUN wget https://mirror.cs.uchicago.edu/google-chrome/pool/main/g/google-chrome-stable/google-chrome-stable_126.0.6478.114-1_amd64.deb \
    dpkg -i google-chrome-stable_126.0.6478.114-1_amd64 \   
    apt-get install -y -f \
    rm google-chrome-stable_126.0.6478.114-1_amd64

ENV CHROME_PATH="/usr/bin/google-chrome-stable"

WORKDIR /app

COPY --from=Builder /app/bin/ .

EXPOSE 8080

ENTRYPOINT ["/app/bin/server"]
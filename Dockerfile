FROM golang:1.16 as builder
ARG TARGETARCH
ARG TARGETOS
WORKDIR /app

COPY . .

RUN apt update && \
    apt install -y --no-install-recommends \
    xvfb libfontconfig wget fontconfig xfonts-75dpi xfonts-100dpi xfonts-scalable xfonts-base \
    && rm -rf /var/lib/apt/lists/* \
    wget  https://github.com/ca-gip/hackathon-api/releases/download/v0.1.1/hackathon-reward-${TARGETOS}-${TARGETARCH} -O hackathon-reward \
    chmod a+x hackathon-reward && \
    mv ./hackathon-reward /usr/local/bin/hackathon-reward

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hackathon-api .

FROM scratch

WORKDIR /app

COPY --from=builder /app/hackathon-api .
COPY --from=builder /usr/local/bin/hackathon-reward /usr/local/bin/hackathon-reward

EXPOSE 8080
CMD [ "./hackathon-api" ]

FROM golang:1.16 as Builder

WORKDIR /app

COPY . .

RUN apt update && \
    apt install -y --no-install-recommends \
    xvfb libfontconfig wget fontconfig xfonts-75dpi xfonts-100dpi xfonts-scalable xfonts-base \
    && rm -rf /var/lib/apt/lists/*

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hackathon-api .

FROM scratch

WORKDIR /app

COPY --from=Builder /app/hackathon-api .

EXPOSE 8080
CMD [ "./hackathon-api" ]

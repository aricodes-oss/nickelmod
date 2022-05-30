# Build phase
FROM golang:latest as build

RUN mkdir /nickelmod
WORKDIR /nickelmod

COPY . .

RUN go build
RUN strip nickelmod

# Deploy only a stripped binary
FROM scratch

COPY --from=build /nickelmod/nickelmod .

ENTRYPOINT ["/nickelmod"]

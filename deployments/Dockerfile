FROM golang:1.20.6-alpine3.18 AS build-stage
WORKDIR /devicemart ##consider this is the workingdir we copy and store all in this, and the rest of the work is in this workdir
COPY  . . ##copy the entire entire things from the current dir
RUN go mod download
RUN go build -v -o ./build/api ./cmd/api/ 

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY  --from=build-stage /devicemart/build/api /
COPY  --from=build-stage /devicemart/.env /
COPY  --from=build-stage /devicemart/templates /templates
EXPOSE 3000
CMD [ "./api" ]
#from image
FROM golang:1.20.6-alpine3.18 AS build-stage
#consider this is the workingdir we copy and store all in this, and the rest of the work is in this workdir
WORKDIR /devicemart 
#copy the entire entire things from the current dir
COPY  . . 
RUN go mod download
# first path is the out put path and the second path is the main path
RUN go build -v -o ./build/bin ./cmd/main

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY  --from=build-stage /devicemart/build/bin/ /
COPY  --from=build-stage /devicemart/.env /
COPY  --from=build-stage /devicemart/web /web
EXPOSE 3000
CMD [ "./main" ]
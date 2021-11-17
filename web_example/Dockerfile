FROM golang:alpine

# make the directory
RUN mkdir /dist

# copy everything to dist folder
ADD . /dist

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Build the application
RUN go build -o main .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/main"]
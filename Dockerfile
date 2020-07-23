##########################
###### INSTALL STAGE #####
##########################

# start from the latest golang base image
FROM golang:latest as INSTALL_STAGE

# set the working directory
WORKDIR /app

# copy dependencies cache files to container
COPY go.mod ./
COPY go.sum ./

# download all dependencies
RUN go mod download

# copy the source code to container
COPY . .

# build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


########################
###### BUILD STAGE #####
########################

# use alpine as base image
FROM alpine:latest  

# add maintainer info
LABEL Author="Caio Krauthamer <caio_k@hotmail.com>"

# add certificates
RUN apk --no-cache add ca-certificates

# set working directory
WORKDIR /root/

# copy the pre-built binary file from the install stage
COPY --from=INSTALL_STAGE /app/main .

# run the binary program
CMD ["./main"]

# Reddit Mediator

Microservice to interact with reddit.
Given an user name and password, app ID and secret it gets a session token via OAuth2.

## Running the project

### Create reddit app

In order to use this you must have a Reddit account and a Reddit app created as `script`.

To create a Reddit account: https://www.reddit.com/

To create a Reddit app: https://www.reddit.com/prefs/apps/

### Set environment variables

The project get credentials from environment variables. For that matter, set the following environment variables on your .bashrc file:
```
export REDDIT_MEDIATOR_USER_USERNAME=<YOUR REDDIT USER>
export REDDIT_MEDIATOR_USER_PASSWORD=<YOUR REDDIT USER PASSWORD>
export REDDIT_MEDIATOR_APP_ID=<YOUR REDDIT APP ID>
export REDDIT_MEDIATOR_APP_SECRET=<YOUR REDDIT APP SECRET>
```

### Executing Reddit Meadiator

```
cd <GO PROJECT FOLDER>
go get github.com/gorilla/mux
go run main.go
```

## Installing and configuring dependencies

For a detailed installation in any environment check https://golang.org/dl/

Go version used during the project is 1.13.4

All golang files should be inside golang workspace. For that matter, create a simple folder called `go` on your home directory with a `bin` folder inside it, where compiled files will be stored, and a `src` folder, where the source code (i.e. this cloned project) will be.

Check https://golang.org/doc/gopath_code.html#Workspaces for further explanation on workspaces.

### Installing golang in Ubuntu

Install it:
```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

### Golang workspace in Ubuntu

You will need to set go environment variables, which refer to `go` and `bin` folders. Add this to your .bashrc (replace `/home/krautz/go` with the path to your go workspace):
```
export GOPATH=/home/krautz/go
export PATH=$PATH:/home/krautz/go/bin
```

### Installing golang in CentOS7

Install it:
```
cd /tmp
curl -O https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
```

### Golang workspace in CentOS7

You will need to set go environment variables, which refer to `go` and `bin` folders. Add this to your .bashrc_profile:
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

### Verifying Installation

```
go version
```

It should display 1.13.4

### Golang editors plugins

For atom, install `go-plus` package.

Check https://golang.org/doc/editors.html for full editors plugins.

### Using golang

To check available commands:
```
go help
```

Initializing module (creates `go.mod` and `go.sum` files):
```
go mod init
```

To install new packages:
```
go get <package>
```
Check https://godoc.org for available packages

Running the project:
```
go run main.go
```

Compiling the project (creates an executable on `bin` folder):
```
go install
```

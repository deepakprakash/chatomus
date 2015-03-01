# Chattic
Exerimental self-hosted instant messaging with zero dependencies.

Chattic lets you deploy your own secure chat server with minimal effort and have others chat with you by visiting your Chattic URL and authenticating themselves.

## Status
Under developement, gathering facts and ongoing research. Nothing significant to try out yet - but you can follow the "Development" instructions to compile and run the current version.

## Aims
Chattic aims to be a lightweight, standalone IM server for individuals (and maybe teams) who want to own their data.

General aims (still evolving) are below. Some of these are long term and will not feature in the alpha version. Feature sets will be defined separately.

### User facing
- Single, independant executables with no dependencies available for all major platforms. Running the server is as simple as running the executable. Eg: `./chattic config.yaml`
- A bundled, canonical web interface for all actions including administrative tasks.
- Usual IM features:
  - Realtime
  - Contact list management
  - Presence
  - History
  - File transfer (Post alpha)
  - Audio/Video (Decision post alpha)
- Export your complete data to standard formats.

### Technical
- Transport over HTTP (and Websockets).
- JSON payloads.
- Well defined API endpoints instead of a specialized protocol.
- Authentication/Authorisation mechanism between Chattic servers. If A and B run separate Chattic instances, then it must be possible for B to authorize herself with A (and vice-versa) without any external dependencies. (Post alpha)

### Non-aims
- Compatibility with existing IM protocols like XMPP.


## Development

Install `golang` (preferably version 1.4) and make sure `go` is added to your PATH and always available on the command line.

### General Setup

Golang requires (or recommends) a specific [workspace arrangement](https://golang.org/doc/code.html). Lets do that setup.

```
# Create a workspace directory for our code
mkdir ~/Code/chattic-workspace
cd ~/Code/chattic-workspace

# Set the [GOPATH](https://code.google.com/p/go-wiki/wiki/GOPATH) to our workspace
export GOPATH=~/Code/chattic-workspace

# Add $GOPATH/bin to PATH so that the compile executables are available on the command line
export PATH=$PATH:$GOPATH/bin

```

We use [`godep`](https://github.com/tools/godep) for dependency management. So install it.

```
go get github.com/tools/godep
```

Once successful, the workspace directory should have a structure similar to this (with more levels):

```
.
├── bin
│   └── godep
├── pkg
│   └── darwin_amd64
└── src
    ├── github.com
    └── golang.org
```

Now we `go get` the project:

```
go get github.com/deepakprakash/chattic
```

The project will be cloned to the directory `./src/github.com/deepakprakash/chattic/`. Whenever we make any changes, it should be in this directory.

Now make sure we install our dependencies using `godep`

```
# Navigate to the chattic source directory
cd ./src/github.com/deepakprakash/chattic

# Ask godep to restore the dependencies
godep restore
```

### Compile, Install and Run

Make sure that the `GOPATH` is set correctly from the above instructions.

To build and install:

```
go install github.com/deepakprakash/chattic
```

This will compile the code and install the `chattic` executable to your `$GOPATH/bin` directory.

Since we also added `$GOPATH/bin` to the `$PATH`, the executable can be run directly from the terminal:

```
chattic
```

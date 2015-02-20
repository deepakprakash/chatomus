# Chatomus
Exerimental self-hosted instant messaging with zero dependencies.

Chatomus lets you deploy your own secure chat server with minimal effort and have others chat with you by visiting your Chatomus URL and authenticating themselves.

## Status
Gathering facts, ongoing research. Nothing to try out yet.

## Aims
Chatomus aims to be a lightweight, standalone IM server for individuals (and maybe teams) who want to own their data.

General aims (still evolving) are below. Some of these are long term and will not feature in the alpha version. Feature sets will be defined separately.

### User facing
- Single, independant executables with no dependencies available for all major platforms. Running the server is as simple as running the executable. Eg: `./chatomus config.yaml`
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
- Authentication/Authorisation mechanism between Chatomus servers. If A and B run separate Chatomus instances, then it must be possible for B to authorize herself with A (and vice-versa) without any external dependencies. (Post alpha)

### Non-aims
- Compatibility with existing IM protocols like XMPP.


## Development

Install `golang` (preferably version 1.4) and make sure `go` is added to your PATH and available on the command line.

### General Setup

Clone the repo locally - say to a directory called `~/Code/`. 

```
cd ~/Code/
git clone https://github.com/deepakprakash/chatomus.git 
```

Golang requires (or recommends) a specific [workspace arrangement](https://golang.org/doc/code.html). Lets get that setup.

```
# Create a workspace directory separate from your 
```



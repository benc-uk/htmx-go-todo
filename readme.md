# HTMX with Go

A example of a simple HTMX app built using Go for learning purposes

Tech stack:
- Go
  - [Echo](https://echo.labstack.com/) (web framework and routing)
- [HTMX](https://htmx.org/) (Dhur)
- [Bulma](https://bulma.io/) (CSS)

Yes it's a todo app, of course it is, what did you expect?!
Currently there is no persistence or state store, everything lives in memory

## Repo

```
├── server      Go code for HTTP server
├── static      Static files, images CSS etc
├── templates   HTML templates, views & fragments
└── todo        Go code for the todo handler
```

## Development

Pre-reqs:

- A recent-ish version of Go (1.20+)
- A Linux environment with bash, make etc. WSL or MacOS is perfect.

### Makefile

```text
help                 💬 This help message :)
install-tools        🔮 Install dev tools into local project tools directory
run                  🚀 Run the server
build                🔨 Build the server
```

### Quick Start

```bash
make install-tools
make run
```

Open http://localhost:4000
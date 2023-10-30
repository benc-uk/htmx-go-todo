# 🌐 HTMX with Go

An example of a simple HTMX app built using Go, created entirely for learning purposes

Tech stack:
- Go
- [Echo](https://echo.labstack.com/) - web framework and routing for Go
- [HTMX](https://htmx.org/) - hypertext magic
- [Bulma](https://bulma.io/) - CSS framework & classes

Yes it's a todo app, of course it is, what did you expect?!
Currently there is no persistence or state store, everything lives in memory

All of this and not a single line of JavaScript! 😃
![screen shot of the app](https://user-images.githubusercontent.com/14982936/279140810-efedc64c-4090-4b1b-adf6-46db4ec3c77a.jpeg)


## 📂 Repo

The structure of the repo is as follows
```
├── build       Dockerfiles and other build artifacts
├── server      Go code for HTTP server
├── static      Static files, images CSS etc
├── templates   HTML templates (see below)
└── todo        Go code for the todo handler
```

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/benc-uk/htmx-go-todo/ci.yml) ![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/m/benc-uk/htmx-go-todo)


## 🧑‍💻 Developer Notes

Pre-reqs:

- A recent-ish version of Go (1.20+)
- A Linux environment with bash, make etc. WSL or MacOS is perfect.

### Makefile

```text
help                 💬 This help message :)
install-tools        🔮 Install dev tools into local project tools directory
run                  🚀 Run the server
build                🔨 Build the server
lint                 🔍 Lint & format check only, sets exit code on error for CI
lint-fix             📝 Lint & format, attempts to fix errors & modify code
```

### Quick Start

```bash
make install-tools
make run
```

Open [http://localhost:4000](http://localhost:4000)

## 📝 Design Notes

The routes supported by the server are a little different from a traditional server side web app, or from a SPA based one.
As HTMX is based on fetching fragments of HTML the following was implemented:

```text
GET    /                       Returns the index.html
GET    /view/{name}            Views are sub-pages rendered under the nav bar
GET    /p/{name}               For URLs linking directly to a specific view
```

```text
GET    /data/todos/            Lists all todos
POST   /data/todos/            Creates a new todo
GET    /data/todos/{id}        Render a single todo row
PUT    /data/todos/{id}        Update a single todo
DELETE /data/todos/{id}        Delete a todo
GET    /data/todos/{id}/edit   Render a todo row for editing
```

The templates aligned to these routes are structured as follows

- `templates/todo/*.html` - Todo fragments mainly returned by the handers for `/data/todos/` routes amd not fetched directly.
- `templates/view/*.html` - Views show sub-sections of the app, and are exposed via the `/view/{name}` route.
- `templates/index/index.html` - Placed into a subdirectory due to the strange way that `template.ParseGlob` works.

NOTE: The names of the templates e.g. `{{ define "todo/list" }}` matches the directory structure e.g. `templates/todo/list.html` but this is just a convention I adopted to keep track of them, there is no connection between a template name and it's filename or directory.
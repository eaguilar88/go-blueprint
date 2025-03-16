# Command Application

A simple CLI tool to scaffold Go project directories and files with default templates. Generate new modules with essential files, ready for development.

### Installation

To install the CLI tool globally:

```sh
go install github.com/eaguilar88/go-blueprint@latest
```

### Options
* generate `<module-name>`: Generate a new module with default files (e.g., endpoints.go, service.go, request.go, response.go) inside a `pkg/<module-name>` directory.

```sh
go-blueprint generate <module-name>
```

This will create a directory structure like:
```
pkg/
└── foo/
    ├── endpoints.go
    ├── service.go
    ├── request.go
    └── response.go
```

Each of the files will start with 
```
package foo

//TODO: Implement
```

### License
This project is licensed under the MIT License. See [LICENSE](LICENSE)
# boostrap-library

- [boostrap-library](#boostrap-library)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Basic Options](#basic-options)
    - [GO](#go)
      - [Options](#options)
      - [Execution](#execution)
    - [STL](#stl)
      - [Options](#options-1)
      - [Execution](#execution-1)
    - [JAVA](#java)
      - [Options](#options-2)
      - [Execution](#execution-2)


## Installation

Each time something is commited (new templates, ...), execute : 

```bash
# GOPROXY=direct is needed to retrieve the latest version (if no versions are defined)
GOPROXY=direct go install github.com/SR-G/bootstrap@latest 
```

Then you can proceed with the individual commands below.

## Usage

This scaffolding project allows to create the skeleton of new : 

- `go` projects
- `stl` projects
- `java` projects (not done yet)

### Basic Options

<!-- weavedoc-bootstrap-base-options:start -->
| Long | Short | Type | Required | Default | Description |
| --- | --- | --- | --- | --- | --- |
| `--name` | - | string | No | - | defaults to the last segment of ModulePath |
| `--output` | - | string | No | - | where the project skeleton is written |
| `--force` | - | bool | No | `false` | overwrite existing files if set |
| `--wipe` | - | bool | No | `false` | overwrite existing files if set |
| `--quiet` | - | bool | No | `false` | disable all logs (for automated executions, ...) |
| `--verbose` | - | bool | No | `false` | activates debug logs |
<!-- weavedoc-bootstrap-base-options:end -->


### GO

#### Options

<!-- weavedoc-bootstrap-go:start -->
| Long | Short | Type | Required | Default | Description |
| --- | --- | --- | --- | --- | --- |
| `--module` | - | string | Yes | - | e.g. github.com/SR-G/my-new-app |
<!-- weavedoc-bootstrap-go:end -->

#### Execution

```bash
go run github.com/SR-G/bootstrap/go@latest --module "github.com/SR-G/<PROJECT_NAME>" 
``` 

### STL

#### Options

<!-- weavedoc-bootstrap-stl:start -->
| Long | Short | Type | Required | Default | Description |
| --- | --- | --- | --- | --- | --- |
| `--url` | - | string | Yes | - | e.g. printables.com/ ... |
| `--nemo` | - | bool | No | `true` | - |
<!-- weavedoc-bootstrap-stl:end -->

#### Execution

```bash
go run github.com/SR-G/bootstrap/stl@latest --url "<URL_NAME>" 
```


### JAVA

TODO

#### Options

<!-- weavedoc-bootstrap-java:start -->
<!-- weavedoc-bootstrap-java:end -->

#### Execution

```bash
```

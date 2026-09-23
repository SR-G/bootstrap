# boostrap-library

- [boostrap-library](#boostrap-library)
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


## Usage

This scaffolding project allows to create the skeleton of new : 

- `go` projects
- `stl` projects
- `java` projects (not done yet)

### Basic Options

<!-- weavedoc-bootstrap-base-options:start -->
| Long        | Short | Type   | Env                            | Required | Default | Description                                      |
| ----------- | ----- | ------ | ------------------------------ | -------- | ------- | ------------------------------------------------ |
| `--name`    | -     | string | `BOOTSTRAPPER_NAME`            | No       | -       | defaults to the last segment of ModulePath       |
| `--output`  | -     | string | `BOOTSTRAPPER_OUTPUT_DIR`      | No       | -       | where the project skeleton is written            |
| `--force`   | -     | bool   | `BOOTSTRAPPER_FORCE_OVERWRITE` | No       | `false` | overwrite existing files if set                  |
| `--wipe`    | -     | bool   | `BOOTSTRAPPER_WIPE_OUTPUT_DIR` | No       | `false` | overwrite existing files if set                  |
| `--quiet`   | -     | bool   | `BOOTSTRAPPER_LOG_QUIET`       | No       | `false` | disable all logs (for automated executions, ...) |
| `--verbose` | -     | bool   | `BOOTSTRAPPER_LOG_DEBUG`       | No       | `false` | activates debug logs                             |
<!-- weavedoc-bootstrap-base-options:end -->


### GO

#### Options

<!-- weavedoc-bootstrap-go:start -->
| Long       | Short | Type   | Env                        | Required | Default | Description                     |
| ---------- | ----- | ------ | -------------------------- | -------- | ------- | ------------------------------- |
| `--module` | -     | string | `BOOTSTRAPPER_MODULE_PATH` | Yes      | -       | e.g. github.com/SR-G/my-new-app |
<!-- weavedoc-bootstrap-go:end -->

#### Execution

```bash
go run github.com/SR-G/boostrap/stl/ --module "github.com/SR-G/<PROJECT_NAME>" 
``` 

### STL

#### Options

<!-- weavedoc-bootstrap-stl:start -->
| Long    | Short | Type   | Env                | Required | Default | Description              |
| ------- | ----- | ------ | ------------------ | -------- | ------- | ------------------------ |
| `--url` | -     | string | `BOOTSTRAPPER_URL` | Yes      | -       | e.g. printables.com/ ... |
<!-- weavedoc-bootstrap-stl:end -->

#### Execution

```bash
go run github.com/SR-G/boostrap/stl/ --url "<URL_NAME>" 
```


### JAVA

TODO

#### Options

<!-- weavedoc-bootstrap-java:start -->
<!-- weavedoc-bootstrap-java:end -->

#### Execution

```bash
```

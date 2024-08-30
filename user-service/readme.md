To ensure that `gqlgen` recognizes your `schema.graphqls` file, you need to configure `gqlgen` properly. Here is a step-by-step plan:

1. **Create the `gqlgen.yml` configuration file**: This file tells `gqlgen` where to find your schema and where to generate the code.

2. **Run `gqlgen` to generate the necessary code**.

### Step-by-Step Plan

1. **Create the `gqlgen.yml` configuration file**:

```yaml
# gqlgen.yml
schema:
  - schema.graphqls

exec:
  filename: graph/generated/generated.go
  package: generated

model:
  filename: graph/model/models_gen.go
  package: model

resolver:
  layout: follow-schema
  dir: graph
  package: graph
```

2. **Run `gqlgen` to generate the code**:

```sh
go run github.com/99designs/gqlgen generate
```
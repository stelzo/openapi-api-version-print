# OpenAPI API Version Print

Prints the version of an API saved inside a OpenAPI Specification in yaml or json.

## Usage

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Get API Version
      
      uses: stelzo/openapi-api-version-print@main
      with:
        specFile: swagger.json
```

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `specFile` _(optional)_  | Path to the OpenAPI Specification file (local or URI). If not given, it looks up openapi.yaml, openapi.yml and openapi.json in the current directory.    |

### Outputs

| Output                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `apiVersion`  | The API version in the OpenAPI Specification file.    |

## Example

### Local

```yaml
steps:
- uses: actions/checkout@master
- name: Get API Version from local file
  
  # id is used to reference the output
  id: api-version-print

  # Run the action
  uses: stelzo/openapi-api-version-print@main
  with:
    specFile: swagger.json

# Somewhere else in the workflow
- name: Check API Version
  run: |
    echo "Version - ${{ steps.api-version-print.outputs.apiVersion }}"
```

### Remote

```yaml
  uses: stelzo/openapi-api-version-print@main
  with:
    specFile: https://raw.githubusercontent.com/dofusdude/api-docs/main/openapi-3.0.yaml
```
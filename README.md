# Introduction

`polygon` provides a simple interface to interact with [Polygon](https://polygon.codeforces.com) via its API. This library contains all the custom objects mentioned in the [API Documentation](https://docs.google.com/document/d/1mb6CDWpbLQsi7F5UjAdwXdbCpyvSgWSXTJVHl52zZUQ/edit?ccid=8880487d88727f44ab2a911727d4d952) (along with convenient functions to pretty print them). It provides functions with just one input parameter for each API end point, while doing the rest of the steps internally, (such as creating the URL according to API specifications, unmarshalling into custom objects, error handling, etc)

For example, once you've created an API object, calling an endpoint is as easy as `api.ProblemsList(parameters)`, where `parameters` is a map containing all your required parameters. 

# Installation
Just running 

```
go get github.com/variety-jones/polygon
```

should add this package to your GOPATH.

# Usage
To use it to create custom applications, you just need to import it via

```
import github.com/variety-jones/polygon
```

After that, each API can be accessed via the `polygon.PolygonApi` object. 

Note that due to **Go** language specifications, all functions/structs/exposed-fields from the `polygon` library would start with a capital letter. Also, you have to skip the dots in the name and capitalize the letter following the dot.

For example, 

`problem.saveStatement` can be called via `api.ProblemSaveStatement`

# Examples
To get a better understanding of how to use this package, please go through this [example](examples/main.go)

# Documentation
For a detailed documentation of each function's description, parameters and  return types, along with the exposed fields of the custom objects, visit this [Package documentation](https://pkg.go.dev/github.com/variety-jones/polygon)

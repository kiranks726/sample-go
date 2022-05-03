# `mainmodule`


This is the directory for the Golang *main module*. 

The module root directory is the directory that contains the `go.mod` file. The `go.mod` file defines the *module path* for package references and package dependencies. The `go.mod` file functions like a node `package.json` file.

Since this is the app/service and not an importable shared library we can name the *module path* after its parent directory `mainmodule`. All the go related source files belong in this directory.

*NOTE: If this were shared library the module should be named after the git repo path for easy importablility by the Golang modules/package system. Example: `github.com/adi-ctx/kitchensink-go`*


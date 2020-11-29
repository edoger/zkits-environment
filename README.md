# ZKits Environment Library #

[![ZKits](https://img.shields.io/badge/ZKits-Library-f3c)](https://github.com/edoger/zkits-environment)
[![Build Status](https://travis-ci.org/edoger/zkits-environment.svg?branch=master)](https://travis-ci.org/edoger/zkits-environment)
[![Coverage Status](https://coveralls.io/repos/github/edoger/zkits-environment/badge.svg?branch=master)](https://coveralls.io/github/edoger/zkits-environment?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5b5f1e62d67846b3813fc77634b8dff3)](https://www.codacy.com/manual/edoger/zkits-environment?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=edoger/zkits-environment&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/edoger/zkits-environment)](https://goreportcard.com/report/github.com/edoger/zkits-environment)
[![Golang Version](https://img.shields.io/badge/golang-1.13+-orange)](https://github.com/edoger/zkits-environment)

## About ##

This package is a library of ZKits project. 
This library provides the function of managing the runtime environment for the application. 
Generally, no additional work is needed to realize the type of runtime environment shared 
between the various components of the application.

## Why? ##

Each component in an application usually needs to have different behaviors (such as more detailed log output) 
in different runtime environments. 
However, it is quite tedious for each component to maintain a set of runtime environment. 
In some large-scale applications, it is impossible to do so. 
This library is to solve this problem. 
Let's manage the runtime environment more conveniently within the application.

## Install ##

```sh
go get -u -v github.com/edoger/zkits-environment
```

## Usage ##

```go
package main
        
import (
    "github.com/edoger/zkits-environment"
)

func main() {
    // Get the current runtime environment.
    environment.Get() 

    // Set the runtime environment value.
    // If the given runtime environment is not supported, ErrInvalidEnv error is returned.
    // If the current runtime environment is locked, ErrLocked error is returned.
    // environment.Development  // "development"
    // environment.Testing      // "testing"
    // environment.Prerelease   // "prerelease"
    // environment.Production   // "production"
    environment.Set(environment.Testing)

    // Register functions can register any custom runtime environment.
    // Note: The Register method must be invoked before calling the Set method.
    environment.Register("foo")
    environment.Registered("foo") // true

    // Lock locks the current runtime environment.
    // After locking, the current runtime environment cannot be changed.
    environment.Lock()
    // SetAndLock sets and locks the current runtime environment.
    // If the runtime environment settings fail, they are not locked.
    environment.SetAndLock(environment.Testing)

    // Listen adds a given runtime environment listener.
    // When the runtime environment changes, all registered listeners will be executed.
    environment.Listen(func(after, before environment.Env) {
        // Do something! 
    })

    // New creates and returns a new instance of the runtime environment manager.
    // The default runtime environment is Development, and all built-in runtime environments
    // have been registered.
    // The runtime environment manager has all the functions and methods of the same name!
    manager := environment.New()

    // NewEmpty creates and returns an empty instance of the runtime environment manager.
    // The manager returned by this function does not register any runtime environment,
    // and the current runtime environment is empty.
    empty := environment.NewEmpty()
}
```

## License ##

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

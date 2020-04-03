# ZKits Environment Library #

[![Build Status](https://travis-ci.org/edoger/zkits-environment.svg?branch=master)](https://travis-ci.org/edoger/zkits-environment)
[![Coverage Status](https://coveralls.io/repos/github/edoger/zkits-environment/badge.svg?branch=master)](https://coveralls.io/github/edoger/zkits-environment?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5b5f1e62d67846b3813fc77634b8dff3)](https://www.codacy.com/manual/edoger/zkits-environment?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=edoger/zkits-environment&amp;utm_campaign=Badge_Grade)

## About ##

This package is a library of ZKits project. 
This library provides the function of managing the runtime environment for the application. 
Generally, no additional work is needed to realize the type of runtime environment shared 
between the various components of the application.

## Usage ##

 1. Import package.
 
    ```sh
    go get -u -v github.com/edoger/zkits-environment
    ```
    
 2. Example.
 
    ```go
    package main
    
    import (
       "flag"
    
       "github.com/edoger/zkits-environment"
    )
    
    func main() {
       var env string
       flag.StringVar(&env, "env", "", "The runtime environment")
       flag.Parse()
    
       // Set the runtime environment value.
       // If the given runtime environment is not registered, an error is returned.
       err := environment.Set(environment.Env(env))
       if err != nil {
           // Handle error.
       }
    
       // Get the current runtime environment.
       // If it has never been set, then you get the environment.Development by default.
       environment.Get()
    
       // Not enough built-in runtime environment?
       // Register functions can register any custom runtime environment.
       // Note: Registration must be before setup.
       environment.Register("foo")
    }
    ```

 3. These are the runtime environments that are already built in and registered.
    ```
    environment.Development  // "development"
    environment.Testing      // "testing"
    environment.Prerelease   // "prerelease"
    environment.Production   // "production"
    ```

## License ##

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

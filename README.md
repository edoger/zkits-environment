# ZKits Environment Library #

[![Build Status](https://travis-ci.org/edoger/zkits-environment.svg?branch=master)](https://travis-ci.org/edoger/zkits-environment)
[![Coverage Status](https://coveralls.io/repos/github/edoger/zkits-environment/badge.svg?branch=master)](https://coveralls.io/github/edoger/zkits-environment?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5b5f1e62d67846b3813fc77634b8dff3)](https://www.codacy.com/manual/edoger/zkits-environment?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=edoger/zkits-environment&amp;utm_campaign=Badge_Grade)

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
       // If the given runtime environment is not supported, ErrInvalidEnv error is returned.
       // If the current runtime environment is locked, ErrLocked error is returned.
       err := environment.Set(environment.Env(env))
       if err != nil {
           // Handle error.
       }
    
       // Get the current runtime environment.
       environment.Get()
    
       // Not enough built-in runtime environment?
       // Register functions can register any custom runtime environment.
       // Note: Registration must be before setup.
       environment.Register("foo")
       
       // Locking the current runtime environment does not allow changes.
       environment.Lock()
       // Determines whether the current runtime environment is locked.
       environment.Locked()
       
       // Sets and locks the current runtime environment.
       // If the runtime environment settings fail, they are not locked.
       err = environment.SetAndLock(environment.Env(env))
       if err != nil {
           // Handle error.
       }
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

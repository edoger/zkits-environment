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

 1. Install package.
 
    ```sh
    go get -u -v github.com/edoger/zkits-environment
    ```
 
 2. Simply get and set up the global runtime environment.
 
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
       err := environment.Set(environment.Testing)
       if err != nil {
           // Handle error.
       }
    }
    ```
 
 3. How do I customize the runtime environment?
 
    ```go
     package main
         
     import (
        "github.com/edoger/zkits-environment"
     )
     
     func main() {
        // Register functions can register any custom runtime environment.
        // Note: The Register method must be invoked before calling the Set method.
        environment.Register("foo")    
     }
     ```
 
 4. How to ensure that the runtime environment is not accidentally modified?
 
    ```go
    package main
    
    import (
       "github.com/edoger/zkits-environment"
    )
    
    func main() {
       // Lock locks the current runtime environment.
       // After locking, the current runtime environment cannot be changed.
       environment.Lock()
       // SetAndLock sets and locks the current runtime environment.
       // If the runtime environment settings fail, they are not locked.
       err := environment.SetAndLock(environment.Testing)
       if err != nil {
            // Handle error.
       }
    }
    ```
 
 5. Want to be notified when the runtime environment changes?

    ```go
    package main
    
    import (
       "github.com/edoger/zkits-environment"
    )
    
    func main() {
       // Listen adds a given runtime environment listener.
       // When the runtime environment changes, all registered listeners will be executed.
       environment.Listen(func(current, old environment.Env) {
            // Do something! 
       })
    }
    ```
 
 6. How does my system manage the runtime environment by itself?

    ```go
    package main
    
    import (
       "github.com/edoger/zkits-environment"
    )
    
    func main() {
       // New creates and returns a new instance of the runtime environment manager.
       // The default runtime environment is Development, and all built-in runtime environments
       // have been registered.
       // The runtime environment manager has all the functions and methods of the same name!
       manager := environment.New()
    }
    ```
 
 7. What are the built-in runtime environments?
    ```
    environment.Development  // "development"
    environment.Testing      // "testing"
    environment.Prerelease   // "prerelease"
    environment.Production   // "production"
    ```

## License ##

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

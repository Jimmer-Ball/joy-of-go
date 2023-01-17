# joy-of-go
Learn by doing

## Basic GO dev environment setup notes

* It looks like GO can be installed anywhere, use https://go.dev/doc/install, and you only really need to amend the environment variable GOROOT to point to a specific installed version, if you've got more than one version of go installed and need to be version specific with your projects.
* Add an ENV_VAR GOPATH that points to a directory under which all local GO projects and dev occurs, so the root of your GO code workspace. 
* For me that GOPATH is *C:\DEV\GO*, and you'll see it can be amended in IntelliJ under *Settings -> Languages And Frameworks -> Go -> GOPATH*
* Make sure your version of IntelliJ Ultimate has the GoLand plugin.  Ultimate does by default, but it looks like it doesn't tell you when a new plugin version is available.  You have to manually check via *Help -> Plugins* every once in a while for plugin updates.
* Make sure *Settings -> Languages And Frameworks -> Go -> Go Modules -> Enable Go Modules Integration* is **ON**, else you will have tests that cannot import packages. This should really be a default in IntelliJ IDEA.

## Building an EXE

Structuring a GO application properly is not easy. This uses the basic ideas outlined 
in https://www.gobeyond.dev/wtf-dial/. Accordingly, separate built artifacts or EXE get 
created according to the contents of a *cmd* package, where all the main build functions 
can go.  Within each *main* function found in the subdirectories of the *cmd* package,
the *import* statements drag in only the different dependencies the runnable end product
needs.  

The idea is to isolate the dependencies of one runnable *main* to the *import* section of the *main*
only, using the directory structure to help isolate dependencies required and winkle out the
runnable programs.

So our application *calculator* has a *cmd* subdirectory.  Within that is a *calculator* subdirectory 
holding the *main* needed to run up the calculator as a EXE. In time the *cmd* subdirectory
will hold extra subdirectories, one for calculator integration tests, and one for calculator storybook 
tests, all of which will contain runnable code.  So, we use are using the *cmd* directory structure 
to organise the runnable things we generate.

Building the application *main.go* in the subdirectory *cmd\calculator* means applying the following command

```shell
go build -o add.exe .\cmd\calculator
```

This will produce an EXE called *add.exe* in the current working directory.  This can be run using the following command

```shell
.\add
```

You can make sure the runnable things we generate, like EXE, are not checked into GIT, by adding the following lines 
to your *.gitignore* file

```shell
*.exe
*.exe~
*.dll
*.so
*.dylib
```
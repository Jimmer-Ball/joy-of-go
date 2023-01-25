# joy-of-go
Learn by doing

## Basic GO dev environment setup notes

* GO can be installed anywhere, use https://go.dev/doc/install as your guide. 
* Installing it adds the GO bin subdirectory to your local user variable *Path*
* You only really need to set the environment variable GOROOT when you've got many installed versions of GO, so you can be specific about which version of many a project should use. 
* Watch out for your *Path* settings if you've got lots of versions of GO installed
* Make sure your version of IntelliJ Ultimate has the GoLand plugin.  Ultimate does by default, but it looks like it doesn't tell you when a new plugin version is available.  You will have to manually check via *Help -> Plugins* every once in a while to discover any GO plugin updates.
* Make sure *Settings -> Languages And Frameworks -> Go -> Go Modules -> Enable Go Modules Integration* is **ON**, else you will have tests that cannot import packages. This should really be a default in IntelliJ IDEA.
* Installing GO will create an ENV_VAR GOPATH that points to a directory under your HOME directory.
* For me that default GOPATH is *C:\Users\jimmer\go*, and the directory does not need to exist
* You'll see GOPATH can be amended in a language wide manner in IntelliJ under *Settings -> Languages And Frameworks -> Go -> GOPATH*
* Make sure your GOPATH in IntelliJ under *Settings -> Languages And Frameworks -> Go - GOPATH* has empty values in the Global, Project, and Module GOPATH dropdown settings 
* Instead of specific Global, Project, and Module settings ensure that *Use GOPATH that's defined in system environment* ticked.
* This way you can have a project open with the IDE with many GO modules, each with their own *go.mod* files and IntelliJ will correctly recognise each as their own GO module
* This way the command line command *go test" and the *run* or *debug* menu options within IntelliJ will behave the same for all modules in a project

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
holding the *main* needed to run up the calculator as an EXE. In time the *cmd* subdirectory
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

## Applications within the project

There are several applications within the project

* _books_: This holds the *books* application codebase
* _calculator_: This holds the *calculator* codebase, which illustrates some good ways of writing tests
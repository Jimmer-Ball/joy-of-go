# joy-of-go
Learn by doing

## Basic GO dev environment setup notes

* It looks like GO can be installed anywhere, use https://go.dev/doc/install, and you only really need to amend the environment variable GOROOT to point to a specific installed version, if you've got more than one version of go installed and need to be version specific with your projects.
* Add an ENV_VAR GOPATH that points to a directory under which all local GO projects and dev occurs, so the root of your GO code workspace. 
* For me that GOPATH is *C:\DEV\GO*, and you'll see it can be amended in IntelliJ under *Settings -> Languages And Frameworks -> Go -> GOPATH*
* Make sure your version of IntelliJ Ultimate has the GoLand plugin.  Ultimate does by default, but it looks like it doesn't tell you when a new plugin version is available.  You have to manually check via *Help -> Plugins* every once in a while for plugin updates.
* Make sure *Settings -> Languages And Frameworks -> Go -> Go Modules -> Enable Go Modules Integration* is **ON**, else you will have tests that cannot import packages. This should really be a default in IntelliJ IDEA.


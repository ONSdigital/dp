Migrating to Go Modules
=======================

A guide on migrating an existing Go project from Vendor to Modules. The examples use the `dp-recipe-api` as reference.

Before you begin it's worth reading the [Golang Modules guide](https://blog.golang.org/using-go-modules) as it serves
 as decent introduction.

### Some title

A Go modules project requires the project to live outside of your go path. 
 - If you haven't already create a new go projects directory outside of your go path - for me this is:
    `/Users/dave/go-projects`
 - Git clone your repo into this directory 
 - Create and swicth to a feature branch for the migration.

### Creating a module
Go modules does support migrating from vendor but I ran into issues whilst trying this. For now I would recommend
 starting fresh and deleting the `/vendor` dir from your project. This will give you a clean slate and bring the
  latest verison of each dependency. If your project relies on a specific version you can change this later.

To create a module run the go mod cmd in root dir of your project or whereever you want the root of your modules to
 be.
 
 ```bash
 go mod init <MODULE_NAME>
 ```
Example:

```bash
go mod init github.com/ONSdigital/dp-recipe-api
```
This will create `go.mod` that should look similar to:
```
module github.com/ONSdigital/dp-recipe-api

go 1.12

require (
    ...
    github.com/ONSdigital/dp-api-clients-go v0.0.0-20190920133223-0b75bbb235dd
    github.com/ONSdigital/dp-mocking v0.0.0-20190905163309-fee2702ad1b9 // indirect
    github.com/ONSdigital/dp-rchttp v0.0.0-20190919143000-bb5699e6fd59
    ... 
)
```
It should also create `go.sum` file.

If you need specific version of dependencies you can edit this file with the specific versions you require. See the 
[Golang Modules guide](https://blog.golang.org/using-go-modules) for details on versioning.

If you run the tests for you module now you should see output similar too:
```
    $ go test ./...
    go: finding <SOME_DEPENDENCY> vX.X.X
    go: downloading <SOME_DEPENDENCY> vX.X.X
    go: extracting <SOME_DEPENDENCY> vX.X.X
    ...
```
Assuming all is well then the dependcies should resolved successfully and your unit tests should pass
. Congratulations you are rocking Go modules! 
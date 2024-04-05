# go-submods-example
An example of setting up Go submodules &amp; tagging them.

## What is this?

There are three `go.mod` files here:

1. The one alongside this README.md ([./go.mod](./go.mod))
1. The one in the public submodule ([./submod/go.mod](./submod/go.mod))
1. The one in the CLI being used to demo the two above ([./command/go.mod](./command/go.mod))

If you look in the command one, you'll see that despite these being in the same repo, they are installed independently of one another.

This allows you to store multiple packages with smaller dependency sets as 'sub-modules' of a main Go repository.

To test this for yourself, try running the following:

```
# Will retrieve the latest version of the root module
go get github.com/hbjydev/go-submods-example

# Will retrieve the latest version of the `submod` module
go get github.com/hbjydev/go-submods-example/submod
```

## Another Thing

They are independently viewable on pkg.go.dev too!

- [Root module](https://pkg.go.dev/github.com/hbjydev/go-submods-example)
- [Submodule](https://pkg.go.dev/github.com/hbjydev/go-submods-example/submod)

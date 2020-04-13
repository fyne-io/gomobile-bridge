# Gomobile Bridge

The Fyne toolkit provides it's own mobile packaging tools which are based on gomobile.
Because of this applications that import gomobile (directly or through a dependency)
may not compile when targetting mobile devices. Or they may not work as expected if one
of the libraries expects gomobile to be loaded.

The gomobile-bridge tool solves this problem by simulating a gomobile API for fyne mobile
apps. This enabled libraries to access the same APIs and yet call the equivalent APIs in
the Fyne ecosystem. To make use of this simply add the following line to your `go.mod` file.

```
replace golang.org/x/mobile => github.com/fyne-io/gomobile-bridge v0.0.1
```

If you have a vendor package it then you should update those files using the `fyne vendor` command:

```
$ go mod tidy && fyne vendor
```

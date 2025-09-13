package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Why '0.out' instead of '.out' or '.cache'?
	// 1. '*.out' is already in the default Go.gitignore.
	// 2. ...but '.out' is not recognized by Go tooling. We need a /^[a-zA-Z0-9]/ filename.

	err := os.MkdirAll("0.out", 0755)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("wkg", "wit", "build", "--output", "./0.out/package.wasm")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %v", cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll("0.out/wit-bindgen-go")
	if err != nil {
		log.Fatal(err)
	}

	// WARNING: Moving packages like this and exposing them as an 'import _ ".../time"' library is a hack that WILL NOT WORK if it has dependencies on
	// other WIT packages like 'wasi:io/streams@0.2.0' or something. Why? Because there would be two copies of the wit-bindgen-go generated Go-side wrapper
	// types (one here and one in the user's code). You wouldn't be able to use them as parameters or return values.

	// Need to use mvpkg instead of --package-root here because --package-root still has the ".../jcbhmr/go/..." WIT package name prefix.

	cmd = exec.Command("go", "tool", "wit-bindgen-go", "generate", "--out", "0.out/wit-bindgen-go", "./0.out/package.wasm")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %v", cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Exclude the 'world std { ... }' package. Top-level things like 'error' and 'complex128' are in the 'builtin' package instead since WIT doesn't support top-level types (it supports top-level functions).
	err = os.RemoveAll("0.out/wit-bindgen-go/jcbhmr/go/std")
	if err != nil {
		log.Fatal(err)
	}

	// golang.org/x/tools/cmd/gomvpkg doesn't support Go modules.
	// mvpkg does NOT clean the target directory. It's SAFE to put custom user-authored code in the same package/directory as long as the file name isn't in the source.
	cmd = exec.Command("go", "tool", "mvpkg", "-build-flags=-tags=tinygo,wasip2", "-recursive", "0.out/wit-bindgen-go/jcbhmr/go", ".")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %v", cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

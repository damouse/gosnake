# glusnek

Gluing Python and Go: call from go to python and back. Goroutine/GIL safe bindings with automatic type conversion. 

While either language can call the other, go has to start the show: in other words this library embeds python within go. To embed go within python check out [gopy](https://github.com/go-python/gopy).

Why is this interesting? Two reasons: because there's a lot of python code I don't want to rewrite in Go, and because go has fantastic support for parallelism and python does not. 

Tested with go1.6, Ubuntu 15.04, python 2.7. No support for python 3 yet. 

## Getting Started

"I dont care why it works, just run the demo"

```
go get https://github.com/damouse/gosnake
cd $GOPATH/src/github.com/damouse/gosnake
export PYTHONPATH=demo
go run demo/main.go
```

If you see a warning about `Python.h` then you need the python dev tools.

```
sudo apt-get install python-dev
```

### Imports and the PATH

`gosnake` imports and calls python packages the same way python does: by checking for the imported package in the `PYTHONPATH`. Lets see an example.

foo.py:

```
def hello():
    return 'Why do humans instinctively fear snakes?'
```

main.py:

```
import foo
print foo.hello()
```

Directory structure 1:

```
demo/
    main.py
    foo/
        hello.py
```

This directory structure works. `main.py` resolves the package `foo` by checking the current directory. 

```
$ cd demo
$ python main.py 
Why do humans instinctively fear snakes?
```

Directory structure 2:

```
demo/
    main.py
foo/
    hello.py
```

This one doesn't, since the package we're looking for is not present in the current directory. If `foo` is an installed package, (i.e. you can run `pip install foo`) then it works again.

An easier way to make the packages visible to python is to temporarily add the target directory to the `PYTHONPATH`.

```
$ cd demo
# export PYTHONPATH="${PYTHONPATH}:.."
$ python main.py 
Why do humans instinctively fear snakes?
```

# TODO

- Dictionary support
- Multi-level imports (from a.b.c import d)
- Exception handling
- Benchmark tests
- Formal stability tests
- Concurrency tests
- Cleanup and deinit 
- Exporting: exportable methods, panic on non-functions, unexport

#### Bugs

- go -> py -> go always returns results as a list. This is likely due to the conversion process in binding.h. See `binding_test.go` line 80.

#### Advanced Features:

- "Permanent", goroutine-safe imports
- Performance Cleanup
    - Dynamically create threads as needed to handle requests
    - Pre-create goroutines for outbound
- Getting object references from python
- Better python interface, preferably screwing with  __attr__
- Exporting entire go package using the reflection tools from github.com/damouse/pinmo

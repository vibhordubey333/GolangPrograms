##### What is memory leak ?

Memory leaks are a class of bugs where memory is not released even after it is no longer needed. They are often explicit, and highly visible, which makes them a great candidate to begin learning debugging. Go is a language particularly well suited to identifying memory leaks because of its powerful toolchain, which ships with amazingly capable tools (pprof) which make pinpointing memory usage easy.
##### Things `pprof` can do:


The way pprof works is using profiles.
A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event, such as allocation.
The file runtime/pprof/pprof.go contains the detailed information and implementation of the profiles.<br/>
Go has several built in profiles for us to use in common cases:

   - goroutine — stack traces of all current goroutines
   - heap — a sampling of memory allocations of live objects
   - allocs — a sampling of all past memory allocations
   - threadcreate — stack traces that led to the creation of new OS threads
   - block — stack traces that led to blocking on synchronization primitives
   - mutex — stack traces of holders of contended mutexes

##### Getting started with pprof

1. Install dependencies: apt-get install graphviz gv (debian) or brew install graphviz (mac)
2. Install pprof: go get -u github.com/google/pprof
3. Add an import: import _ "net/http/pprof"
4. Add a server for pprof:
   ```
   go func() {
   fmt.Println(http.ListenAndServe("localhost:6060", nil))
   }()
   ```

##### Commands related to analyzing pprof graph.

1. `go tool pprof -web http://localhost:7777/debug/pprof/goroutine`
2. To fire up web-based `go tool pprof -web pprof.goroutine.001.pb.gz`

####
   
---
#### References:

1. https://dev.to/agamm/how-to-profile-go-with-pprof-in-30-seconds-592a
2. 
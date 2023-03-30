##### Debugging Go Programs Using Delve
1. In terminal fire: `go get github.com/go-delve/delve/cmd/dlv`
2. To open delve
3. To set break point. Mention your file name with line number: `break your-program-file-name:19`
4. Then to proceed further fire in delve prompt itself `continue`
5. To list function arguments. Fire up `args`
6. To print local variables `locals`.
7. To jump to next line `next`.

##### FAQ

1. Difference between `next` and `continue` ?
    


---
##### References For Delve:

    - https://vtimothy.com/posts/debugging-goroutines/
    
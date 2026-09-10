# md-to-html

A small Go utility that scans Markdown files, finds inline HTTP links, checks each one with a concurrent worker pool, and prints whether each link is working or broken.

## What the project does

The program does the following:

1. Walks the `md-to-html` directory recursively.
2. Finds `.md` files.
3. Extracts Markdown links that look like `[text](https://...)`.
4. Sends each discovered URL to a worker pool.
5. Uses an HTTP client to check the response status.
6. Prints `[OK]` for healthy links and `[BROKEN]` for invalid or failed links.

From the `md-to-html` folder:

```bash
go run ./cmd/main.go
```

## Sample output

When run, prints like this:

```text
Searching directory:  /Users/amish130437/Desktop/learning_Go/md-to-html
Found 12 links...
[BROKEN] Error: Head "https://docs.go.dev/search?q=interfaces": dial tcp: lookup docs.go.dev: no such host https://docs.go.dev/search?q=interfaces
[OK] http://example.com
[OK] https://gobyexample.com/
[OK] https://www.google.com
[BROKEN] Error: 403 Forbidden https://exercism.org/tracks/go
[OK] https://go.dev/doc/effective_go
[OK] https://go.dev/tour/
[OK] https://go.dev/play/
[OK] https://go.dev/doc/
[OK] https://gophercises.com/
[OK] https://github.com/quii/learn-go-with-tests
```

## Notes

- The parser currently matches only Markdown inline links in the form `[text](http(s)://...)`.
- Relative links like `./docs/install.md` and anchors like `#conclusion` are intentionally ignored.

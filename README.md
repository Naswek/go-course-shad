# Go Course — SHAD

My solutions to the Go programming course assignments from the Yandex School of Data Analysis (SHAD).

This repository is used to track my progress in learning Go, practicing problem-solving, and preparing for backend development.

**Original course:** [manytask/itmo-go](https://gitlab.com/manytask/itmo-go/public)  
**Lectures:** [Go course slides](https://p.go.manytask.org/)  
**Syllabus:** [Course syllabus](docs/syllabus.md)

## Progress

**Completed: 1 / 57 assignments**

Tasks are grouped according to the original course configuration.

### 01. Hello World

- [ ] [sum](sum/)
- [ ] [tour0](tour0/)
- [ ] [wordcount](wordcount/)
- [ ] [urlfetch](urlfetch/)
- [ ] [fetchall](fetchall/)

### 02. Basics

- [x] [hotelbusiness](hotelbusiness/)
- [ ] [hogwarts](hogwarts/)
- [ ] [utf8](utf8/)
- [ ] [varfmt](varfmt/)
- [ ] [speller](speller/)
- [ ] [forth](forth/)

### 03. Interfaces

- [ ] [otp](otp/)
- [ ] [lrucache](lrucache/)
- [ ] [externalsort](externalsort/)
- [ ] [retryupdate](retryupdate/)
- [ ] [ciletters](ciletters/)

### 04. Goroutines

- [ ] [tour1](tour1/)
- [ ] [once](once/)
- [ ] [rwmutex](rwmutex/)
- [ ] [waitgroup](waitgroup/)
- [ ] [cond](cond/)
- [ ] [ratelimit](ratelimit/)

### 05. Testing

- [ ] [testequal](testequal/)
- [ ] [fileleak](fileleak/)
- [ ] [tabletest](tabletest/)
- [ ] [tparallel](tparallel/)
- [ ] [iprange](iprange/)

### 06. Concurrency with Shared Memory

- [ ] [dupcall](dupcall/)
- [ ] [keylock](keylock/)
- [ ] [batcher](batcher/)
- [ ] [pubsub](pubsub/)

### 07. HTTP

- [ ] [urlshortener](urlshortener/)
- [ ] [digitalclock](digitalclock/)
- [ ] [middleware](middleware/)
- [ ] [olympics](olympics/)
- [ ] [firewall](firewall/)

### 08. SQL

- [ ] [dao](dao/)
- [ ] [ledger](ledger/)
- [ ] [shopfront](shopfront/)
- [ ] [wscat](wscat/)

### 09. Generics

- [ ] [genericsum](genericsum/)
- [ ] [treeiter](treeiter/)
- [ ] [coverme](coverme/)

### 10. Reflection

- [ ] [reversemap](reversemap/)
- [ ] [jsonlist](jsonlist/)
- [ ] [jsonrpc](jsonrpc/)
- [ ] [structtags](structtags/)

### 11. Low Level

- [ ] [illegal](illegal/)
- [ ] [blowfish](blowfish/)

### 12. Analysis

- [ ] [testifycheck](testifycheck/)
- [ ] [gzep](gzep/)

### Additional Projects

- [ ] [gitfame](gitfame/)
- [ ] [distbuild](distbuild/)

### Bonus

- [ ] [consistenthash](consistenthash/)
- [ ] [gossip](gossip/)
- [ ] [smartsched](smartsched/)
- [ ] [wasm](wasm/)

---

## Running Tests

Run tests for a specific assignment:

```bash
go test -v ./hotelbusiness/...
```

Run all tests:

```bash
go test ./...
```

## Notes

- Each completed assignment is tracked with a separate Git commit.
- Some assignments may be revisited for refactoring and optimization.
- The original assignments, tests, and course materials belong to their respective authors.

## Credits

Based on the [Go course repository](https://gitlab.com/manytask/itmo-go/public) from SHAD.

See [LICENSE](LICENSE) for the original license.

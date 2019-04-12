echo 1
go install -i -a -v cmd/compile
echo 2
go install -i -a -v cmd/link
echo 3
go install -i -a -v cmd/go
echo 4
go install -i -a -v cmd/addr2line 
go install -i -a -v cmd/api 
go install -i -a -v cmd/asm 
go install -i -a -v cmd/buildid
go install -i -a -v cmd/cgo
go install -i -a -v cmd/cover
go install -i -a -v cmd/dist
go install -i -a -v cmd/doc
go install -i -a -v cmd/fix
go install -i -a -v cmd/gofmt
go install -i -a -v cmd/nm
go install -i -a -v cmd/objdump
go install -i -a -v cmd/pack
go install -i -a -v cmd/pprof
go install -i -a -v cmd/test2json
go install -i -a -v cmd/trace
go install -i -a -v cmd/vet


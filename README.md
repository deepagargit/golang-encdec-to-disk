# golang-encdec-to-disk
This golang covers struct (with fields as map, struct etc) encode decode to disk

Pull the repository to disk

1. Install Golang on Windows
https://golang.org/doc/install?download=go1.5.3.windows-amd64.msi


2.  Install Git bash on Windows
https://git-scm.com/download/win

3. Get the Golang gods package
a. Set the GOPATH
$ export GOPATH=/c/New/exo/exoRedis

b. Get the package
$ go get github.com/emirpasic/gods/sets/treeset

4. Build
go build

$./golang-encdec-to-disk.exe
Data structure saved to disk
 save :  {3 4 5 Pythagoras {10 11 Trad} {{1 0} 0 0 -1073741824 0} map[1:{6 7 Qua
d} 2:{8 9 R}] 0xc082026060}
 save :  [A B]

Data structure loaded from disk 
 load :  {3 4 5 Pythagoras {10 11 Trad} {{0 0} 0 0 0 0} map[1:{6 7 Quad} 2:{8 9
R}] 0xc082026230}
 load :  [A B]

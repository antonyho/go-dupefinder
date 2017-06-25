# go-dupefinder
A lightweight application to lookup duplicated files written in Go

Libraries:
----------
* go dep
* sqlite
* cmd


**TODO**
* Initialise SQLite on-memory database
* Build file table with: file full path, file size, hash, file type?, create date, last modify date
* Compare on file hash
* Compare on file size
* Compare on file type?
* Check byte to verify whether they are identical
* Keep on zero size file

# go-dupefinder
A lightweight application to lookup duplicated files written in Go

Libraries:
----------
* go dep
* sqlite
* cmd


**TODO**
- [x] Initialise SQLite on-memory database
- [x] Build file table with: file full path, file size, hash, file type?, create date, last modify date
- [x] Compare on file hash
- [x] Compare on file size
- [ ] Compare on file type?
- [ ] Check byte to verify whether they are identical
- [ ] Skip zero size file

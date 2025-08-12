Simple variable-length code archiver:

Downloading dependencies and building project
```bash
go mod tidy
go build
```
Packing .txt file
```bash
.\vlc-archiver.exe pack -m vlc path/to/file.txt
```
Unpacking .vlc file
```bash
.\vlc-archiver.exe pack -m vlc path/to/file.vlc
```

cara run unit test golang :
1. untuk semua file golang dari root path
   go test -v ./...
2. untuk semua file di dalam folder
   go test -v 
3. untuk spesifik file di dalam folder
   go test -v -run namaFunc
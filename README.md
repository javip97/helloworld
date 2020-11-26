curl -L -O https://github.com/javip97/helloworld/archive/main.zip
unzip main.zip
cd helloworld-main
./helloworld



# CGO_ENABLED=0 GOOS=aix GOARCH=ppc64 go build -o binary
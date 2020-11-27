# Start
cd usr
cd go
cd helloworld-main


curl -L -O https://github.com/javip97/helloworld/archive/main.zip
unzip main.zip
cd helloworld-main
./binary -fileMode -agentHost="thales-cert" -fileBasePath=""


# Clean
rm -rf helloworld-main
rm main.zip


# Compile
# CGO_ENABLED=0 GOOS=aix GOARCH=ppc64 go build -o binary
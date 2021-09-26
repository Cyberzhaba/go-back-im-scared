# Go-Backend (for Mercuryo)
## Build and run
```
git clone https://github.com/Cyberzhaba/go-back-im-scared.git
cd go-back-im-scared
cp configs/apiserver.toml.bak configs/apiserver.toml
make build
./apiserver
```

## From Dockerfile
```
git clone https://github.com/Cyberzhaba/go-back-im-scared.git
cd go-back-im-scared
sudo docker build -t gobackend . 
sudo docker run -d -p 7777:7777 gobackend
```
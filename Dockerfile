FROM golang:latest
RUN git clone https://github.com/Cyberzhaba/go-back-im-scared.git
WORKDIR go-back-im-scared
RUN cp configs/apiserver.toml.bak configs/apiserver.toml
RUN make build
CMD ./apiserver
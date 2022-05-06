From debian:stretch-slim

WORKDIR /

COPY my-scheduler /usr/local/bin

CMD ["my-scheduler"]

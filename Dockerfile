FROM ubuntu

ADD simple-server /simple-server
ENTRYPOINT ["/simple-server"]

FROM reg.lechange.com/library/debian:7.8
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8  

ADD simple-server /
ENTRYPOINT ["/simple-server"]

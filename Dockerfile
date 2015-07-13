FROM whatwedo/golang:latest

#Maintainer
MAINTAINER Felix Imobersteg <felix.imobersteg@me.com>

#Update package lists
RUN apt-get update -y

#Install build tools
RUN apt-get install -y make git

#Add and compile source code
ADD . /usr/src/ddesktop
WORKDIR /usr/src/ddesktop
RUN rm -rf src/github.com src/gopkg.in pkg
RUN make backend
WORKDIR /root

#Move binary and config
RUN cp -R /usr/src/ddesktop/bin/* /root
RUN mkdir -p /etc/ddesktop && cp /usr/src/ddesktop/config.yml /etc/ddesktop

#Slimming down Docker container
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

#Expose Ports
EXPOSE 9000

#Alter upstart script
RUN echo -n "cd /root && ./ddesktop" >> /bin/upstart

#Set upstart script
CMD /bin/upstart

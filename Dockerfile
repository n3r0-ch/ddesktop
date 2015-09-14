FROM whatwedo/golang:latest

#Maintainer
MAINTAINER Felix Imobersteg <felix.imobersteg@me.com>

#Update package lists
RUN apt-get update -y

#Install build tools
RUN curl -sL https://deb.nodesource.com/setup | bash -
RUN apt-get install -y make git apache2-utils nodejs
RUN npm install -g grunt-cli bower

#Add and compile source code
ADD . /usr/src/ddesktop
WORKDIR /usr/src/ddesktop
RUN rm -rf src/github.com src/gopkg.in src/golang.org pkg bin
RUN make all
WORKDIR /root

#Move binary and config
RUN cp -R /usr/src/ddesktop/bin/* /root
RUN mkdir -p /etc/ddesktop && cp /usr/src/ddesktop/config.yml /etc/ddesktop

#Create sample SSL certificate
RUN openssl req -new -newkey rsa:4096 -days 3652 -nodes -x509 -subj "/C=CH/ST=ddesktop/L=ddesktop/O=ddesktop/CN=ddesktop.io" -keyout /etc/ddesktop/key.pem  -out /etc/ddesktop/cert.pem

#Create passwordfile
RUN htpasswd -cb /etc/ddesktop/.htpasswd ddesktop ddesktop

#Slimming down Docker container
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

#Expose Ports
EXPOSE 80
EXPOSE 443

#Create volume
VOLUME /etc/ddesktop

#Alter upstart script
RUN echo -n "cd /root && ./ddesktop" >> /bin/upstart

#Set upstart script
CMD /bin/upstart

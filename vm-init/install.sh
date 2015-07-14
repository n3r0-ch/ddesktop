#!/bin/bash
##################################################################
#                                                                #
# @Version:				0.1                                      #
# @Author:				Felix Imobersteg  			  			 #
# Description:                                                   #
# This script install required VM dependecies                    #
#													             #
##################################################################

#Change directory
cd /tmp

#Update sytem
apt-get update -y
apt-get upgrade -y

#Install base packages
apt-get install make nano vim less zip unzip telnet git curl mercurial ca-certificates -y

#Install go
curl -s https://storage.googleapis.com/golang/go1.3.linux-amd64.tar.gz | tar -v -C /usr/local -xz

#Install Frontend deps
apt-get install -y npm nodejs
ln -s /usr/bin/nodejs /usr/bin/node
npm install -g grunt-cli bower

#Install docker
curl -sSL https://get.docker.com/ubuntu/ | sh

#Enable aufs
apt-get -y install linux-image-extra-$(uname -r) aufs-tools
apt-get -y install lxc-docker
service docker restart

#Use docker with vagrant user without sudo
gpasswd -a vagrant docker
service docker restart

#Create sample SSL certificate
mkdir -p /etc/ddesktop
openssl req -new -newkey rsa:4096 -days 3652 -nodes -x509 -subj "/C=CH/ST=ddesktop/L=ddesktop/O=ddesktop/CN=ddesktop.io" -keyout /etc/ddesktop/key.pem  -out /etc/ddesktop/cert.pem 

#Install docker-compose
curl -L https://github.com/docker/compose/releases/download/1.1.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

#Add go to path
echo 'export PATH=$PATH:/usr/local/go/bin:/go/bin' >> /home/vagrant/.bashrc

#Pull docker images
docker pull n3r0ch/ddesktop-client:latest
docker pull n3r0ch/ddesktop-server:latest

#Cleanup
apt-get autoremove -y
apt-get autoclean -y


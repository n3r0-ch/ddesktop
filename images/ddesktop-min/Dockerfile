FROM ubuntu:14.04

#Maintainer
MAINTAINER Felix Imobersteg <felix.imobersteg@me.com>

#Base settings
ENV DEBIAN_FRONTEND noninteractive
ENV HOME /root

#Prevent initramfs updates from trying to run grub and lilo.
#https://journal.paul.querna.org/articles/2013/10/15/docker-ubuntu-on-rackspace/
#http://bugs.debian.org/cgi-bin/bugreport.cgi?bug=594189
RUN export INITRD=no
RUN mkdir -p /etc/container_environment
RUN echo -n no > /etc/container_environment/INITRD

## Enable Ubuntu Universe and Multiverse.
RUN sed -i 's/^#\s*\(deb.*universe\)$/\1/g' /etc/apt/sources.list
RUN sed -i 's/^#\s*\(deb.*multiverse\)$/\1/g' /etc/apt/sources.list

#Update system
RUN apt-get update -y
RUN apt-get upgrade -y

#Fix locales
RUN apt-get install -y language-pack-en
RUN locale-gen en_US
RUN update-locale LANG=en_US.UTF-8 LC_CTYPE=en_US.UTF-8
RUN echo -n en_US.UTF-8 > /etc/container_environment/LANG
RUN echo -n en_US.UTF-8 > /etc/container_environment/LC_CTYPE
RUN apt-get install localepurge

#Fix: TERM environment variable not set.
ENV TERM dumb

#Create upstart script
RUN echo "#\0041/bin/bash" > /bin/upstart
RUN chmod 755 /bin/upstart

#Add firstbootscript
ADD files/firstboot /bin
RUN echo "/bin/firstboot" >> /bin/upstart
RUN chmod 755 /bin/firstboot

#Install software
RUN apt-get install -y sudo net-tools lxde gtk2-engines-murrine \ 
	ttf-ubuntu-font-family fonts-wqy-microhei lxterminal \ 
	x11vnc Xvfb supervisor pwgen websockify curl less nano wget \ 
	unzip

#Create ubuntu user
RUN useradd --create-home --shell /bin/bash --user-group --groups adm,sudo ubuntu

#Add supervisor config
ADD files/supervisor /etc/supervisor/conf.d

#Edit firstboot script
ADD files/pw-reset /bin
RUN chmod 755 /bin/pw-reset
RUN echo "/bin/pw-reset" >> /bin/firstboot

#Slimming down Docker container
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

#Expose Ports
EXPOSE 6080

#Alter upstart script
RUN echo -n "exec /usr/bin/supervisord -n" >> /bin/upstart

#Set upstart script
CMD /bin/upstart

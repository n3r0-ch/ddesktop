FROM whatwedo/base:latest

#Maintainer
MAINTAINER Felix Imobersteg <felix.imobersteg@me.com>

#Update system
RUN apt-get update -y
RUN apt-get upgrade -y

#Install software
RUN apt-get install -y sudo net-tools lxde gtk2-engines-murrine \ 
	ttf-ubuntu-font-family fonts-wqy-microhei lxterminal \ 
	x11vnc Xvfb supervisor pwgen websockify curl less nano wget \ 
	unzip firefox libreoffice retext evince gimp gimp-data \ 
	gimp-plugin-registry gimp-data-extras filezilla chromium-browser

#Create ubuntu user
RUN useradd --create-home --shell /bin/bash --user-group ubuntu

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

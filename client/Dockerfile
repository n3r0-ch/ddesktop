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
	gimp-plugin-registry gimp-data-extras filezilla cifs-utils

#Create ubuntu user
RUN useradd --create-home --shell /bin/bash --user-group ubuntu

#Edit startmenu and taskbar
RUN cd /usr/share/applications && rm -f debian-uxterm.desktop \
	debian-xterm.desktop gcr-prompter.desktop gcr-viewer.desktop \
	gksu.desktop lxde-logout.desktop lxde-screenlock.desktop \
	lxde-x-terminal-emulator.desktop openjdk-7-policytool.desktop \
	python2.7.desktop python3.4.desktop x11vnc.desktop \
	xscreensaver-properties.desktop
RUN rm -f /etc/xdg/menus/lxde-applications.menu /usr/share/lxpanel/profile/LXDE/panels/panel
ADD files/menu /etc/xdg/menus
ADD files/taskbar /usr/share/lxpanel/profile/LXDE/panels

#Set wallpaper
USER ubuntu
RUN mkdir -p /home/ubuntu/.config/pcmanfm/LXDE
ADD files/pcmanfm /home/ubuntu/.config/pcmanfm/LXDE
USER root

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

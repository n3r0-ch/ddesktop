#.ddesktop

####Virtual desktops powered by Docker

![Screenshot](https://raw.githubusercontent.com/n3r0-ch/ddesktop/master/screenshot.png)

##Features
* Access virtual desktops powered by [Docker](https://www.docker.com/)
* Instant creation of new containers on access
* Open, free and easy virtual desktop solution
* Fast and performant
* Less operation effort required compared to other virtual desktop environments


##Usage
* Run the following command:

```docker run -p 80:80 -p 443:443 -v /var/run/docker.sock:/var/run/docker.sock n3r0ch/ddesktop-server```

* Access the webinterface under `https://[your-dockerhost]` and the following credentials

| | |
|---|---|
| User: | ddesktop |
| Password | ddesktop  |


##Configuration
Connect to the running [Docker](https://www.docker.com/) container with `docker exec -it [container-id] bash` and edit the well commented configuration file under `/etc/ddesktop/config.yml`


##Usermanagement
Connect to the running [Docker](https://www.docker.com/) container with `docker exec -it [container-id] bash` and edit the htpasswd file. The path can be set under `/etc/ddesktop/config.yml` (Default: `/etc/ddesktop/.htpasswd`). 


##Client Image
The default client image is set to [n3r0ch/ddesktop-client](https://registry.hub.docker.com/u/n3r0ch/ddesktop-client/). This images suffices for the most ddesktop environments, but it's also possible to build a new image. The easiest way is to use the [n3r0ch/ddesktop-client](https://registry.hub.docker.com/u/n3r0ch/ddesktop-client/) image as base image.

A .ddesktop image needs a running [websockify](https://github.com/kanaka/websockify) on port 6080 redirecting forwarding to the container local VNC server.


##Current limitations
* No integrated persistent file storage
* No userprofiles
* Just a single image for all users
* No usermanagement over the webinterface
* Only local [Docker](https://www.docker.com/) environments supported (no remote docker socket)


##Bugs and Issues
If you have any problems with this image, feel free to open a new issue in our issue tracker [https://github.com/n3r0-ch/ddesktop/issues](https://github.com/n3r0-ch/ddesktop/issues)


##License
The MIT License (MIT)

Copyright (c) 2014 Felix Imobersteg

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

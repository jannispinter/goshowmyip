# goshowmyip
A simple IP address and User Agent displaying web app written in Go. As an additional gimmick, the web app fetches random kitten pictures from flickr.

# Installation
The idea is to run this behind a reverse proxy, preferably nginx. You need to add an additional Header to preserve the original IP address.

Clone the source code and build the program:
```sh
git clone https://github.com/jannispinter/goshowmyip
cd goshowmyip
go build
```

Here is an example nginx configuration:
```
...
        location / {
		proxy_set_header        X-Real-IP       $remote_addr;
		proxy_pass http://localhost:8080;	
	}
...
```

You also need a Flickr API key which has to be added to the source code, recompile it afterwards. Launch the binary on your server and have fun.

# Demo
https://ip.pinterjann.is

# License
 ![gpl3](https://www.gnu.org/graphics/gplv3-127x51.png)

  goshowmyip is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  goshowmyip is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

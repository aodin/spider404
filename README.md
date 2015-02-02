# spider404
Check for 404 on a website


### Install

#### gokogiri

1. Install libxml2

    brew install libxml2

2. Symbolic link to pkgconfig - location and version may vary

    sudo ln -s /usr/local/Cellar/libxml2/2.8.0/lib/pkgconfig/libxml-2.0.pc /usr/local/lib/pkgconfig/libxml-2.0.pc

3. Check with pkg-config:

    pkg-config --cflags libxml-2.0 libxml-2.0

4. Install gokogiri

    go get github.com/moovweb/gokogiri

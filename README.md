SynComm
=======

Aggregates and distributes messages.

Code written in go (golang).

Installation
============
it looks like you need libzmq-dev and libzmq3-dev but they conflict so...
```
sudo apt-get install golang libtool autoconf automake uuid-dev build-essential libzmq3 libzmq3-dev libzmq3-dbg pkg-config
sudo apt-get install libzmq-dev
```
set up paths with export or add to /etc/environment
```
mkdir $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
packages
```
go get github.com/alecthomas/gozmq
```
to compile a newer zmq; good luck
```
wget http://download.zeromq.org/zeromq-4.0.4.tar.gz
tar zxvf zeromq-4.0.4.tar.gz && cd zeromq-4.0.4
./configure
make && make install
sudo ldconfig
```
Now, to get the lua zmq binding, you can use luarocks or cmake it
```
git clone git@github.com:Neopallium/lua-zmq.git
sudo /usr/local/openresty/luajit/bin/luarocks install lua-zmq/rockspecs/lua-zmq-scm-1.rockspec ZEROMQ_LIBDIR=/usr/lib/x86_64-linux-gnu/
```
cmake:
```
git clone git@github.com:Neopallium/lua-zmq.git
cd lua-zmq ; mkdir build ; cd build
cmake ..
# now set LUA_INCLUDE_DIR=/usr/local/openresty/luajit/include/luajit-2.1 in CMakeCache.txt
cmake ..
make
sudo make install
```

If you would like to install SynComm from git without asking for your credentials, then change your ~/.gitconfig to have:
[url "git@github.com:"]
    insteadOf = https://github.com/

Finally:
```    
go get github.com/Lumate/SynComm
```

# ircproxy

filtering i2p irc client proxy server supporting anonymous dcc

## building

to build use `go get`

    $ go get -u github.com/majestrate/ircproxy
    
make sure you have `$GOPATH/bin` (usually `~/go/bin` ) added to `$PATH`

then run the server with

     $ ircproxy 6mk5za2izxm5ubu7bhzw3io7x5h6yjnlc7iccmn2ilbwptceaiwq.b32.i2p

then connect to `irc://127.0.0.1/`

replace `6mk5za2izxm5ubu7bhzw3io7x5h6yjnlc7iccmn2ilbwptceaiwq.b32.i2p` with another irc network that supports dcc if it's down after trying many times.

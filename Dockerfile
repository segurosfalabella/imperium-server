FROM scratch

ADD /imperium-server /imperium-server

ENTRYPOINT ["/imperium-server"]

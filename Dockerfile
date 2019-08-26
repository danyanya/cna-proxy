FROM centurylink/ca-certs
MAINTAINER dsliusar@wimark.com

ADD cna-proxy /

ENTRYPOINT ["/cna-proxy"]
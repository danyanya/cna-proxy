FROM centurylink/ca-certs
LABEL author="danyanya <danya.brain@gmail.com>"

ADD cna-proxy /

ENTRYPOINT ["/cna-proxy"]
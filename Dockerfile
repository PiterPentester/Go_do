### Dockerfile for my own GO_DO app  (by Bogdan Pakhomov) ###

FROM centos:centos7

COPY main /root/main

RUN yum install -y epel-release git \
    && useradd -mr git \
    && chmod +x /root/main \
    && touch /root/start.sh \ 
    && echo "#!/bin/sh" >> /root/start.sh \
    && echo "runuser -l git -c '/root/main'" >> /root/start.sh \
    && chmod ugo+x /root/start.sh

EXPOSE 9876

ENTRYPOINT ["/root/start.sh"]

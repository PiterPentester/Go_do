### Dockerfile for my own GO_DO app  (by Bogdan Pakhomov) ###

FROM centos:centos7

COPY main /root/main

RUN useradd -mr git \
    && chmod +x /root/main \
    && mkdir -p /home/git/go_do \
    && mv /root/main /home/git/go_do \
    && chown -R git:git /home/git/go_do \
    && touch /root/start.sh \ 
    && echo "#!/bin/sh" >> /root/start.sh \
    && echo "runuser -l git -c '/home/git/go_do/main'" >> /root/start.sh \
    && chmod ugo+x /root/start.sh

EXPOSE 9876

ENTRYPOINT ["/root/start.sh"]

### Dockerfile for my own GO_DO app  (by Bogdan Pakhomov) 16.11.19 ###

FROM centos

# Copy the local package files to the container’s workspace.

COPY godo-app /

RUN groupadd -r godo -g 433 \ 
    && useradd -u 431 -r -g godo -d /app -s /sbin/nologin -c "Docker image user" godo \
    && mkdir -p /home/godo \
    && mv /godo-app /home/godo/godo-app \
    && chown -R godo:godo /home/godo

EXPOSE 9876

VOLUME ["/home/godo"]

USER godo

WORKDIR ["/home/godo"]

CMD ["./godo-app"]

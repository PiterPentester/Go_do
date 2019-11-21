### Dockerfile for my own GO_DO app  (by Bogdan Pakhomov) 16.11.19 ###

FROM scratch

# Copy the local package files to the container’s workspace.
ADD app.ini godo-app /

EXPOSE 9876

CMD ["/godo-app"]

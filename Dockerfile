FROM scratch
COPY dist/accounts /
EXPOSE 80
VOLUME ["/config"]
ENTRYPOINT ["/accounts"]
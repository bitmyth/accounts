FROM scratch
COPY dist/accounts /
EXPOSE 80
VOLUME ["/tmp"]
ENTRYPOINT ["/traefik"]
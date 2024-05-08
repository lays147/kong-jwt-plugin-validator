FROM golang:1.21-bullseye AS plugin-builder

WORKDIR /
COPY . .
RUN chmod +x scripts/go_build.sh && bash scripts/go_build.sh

FROM kong:3.6.1-ubuntu

USER root

COPY --chown=kong:kong config/kong.conf /etc/kong
COPY --chown=kong:kong --from=plugin-builder plugins /usr/local/bin

USER kong

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 8000 8443 8001 8002 8444

STOPSIGNAL SIGQUIT
HEALTHCHECK --interval=60s --timeout=10s --retries=10 CMD kong health

CMD ["kong", "docker-start"]

# Certificate dependencies
FROM anzx-docker.artifactory.gcp.anz/library/certificates:latest AS certificates

# Build
FROM hub.artifactory.gcp.anz/golang:1.20.2-bullseye as builder
ARG GOPROXY=https://platform-gomodproxy.services-platdev.x.gcpnp.anz/,https://artifactory.gcp.anz/artifactory/api/go/go,direct
ENV GOPROXY=${GOPROXY}

# Get certificates for building binaries
COPY --from=certificates /global/*.crt /usr/local/share/ca-certificates/
COPY --from=certificates /globaltest/*.crt /usr/local/share/ca-certificates/
COPY --from=certificates /external/DigiCert_Global_Root_CA.crt /usr/local/share/ca-certificates/
RUN /usr/sbin/update-ca-certificates

RUN mkdir /app
COPY ./bin/linux_todoInGo.test /app/bin/todoInGo.test
COPY todoInGoService.sh /app
WORKDIR /app
CMD ["./todoInGoService.sh"]
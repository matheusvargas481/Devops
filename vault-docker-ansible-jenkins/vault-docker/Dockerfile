FROM alpine:3.7

ENV VAULT_VERSION 1.3.0

RUN mkdir /vault

RUN apk --no-cache add openssl wget

RUN wget --quiet --output-document=/tmp/vault.zip https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_linux_amd64.zip && \
    unzip /tmp/vault.zip -d /vault && \
    rm -f /tmp/vault.zip && \
    chmod +x /vault

ENV PATH=/vault
ENV VAULT_ADDR 'http://127.0.0.1:8200'

EXPOSE 8200

ENTRYPOINT [ "/bin/ash", "-c", "vault server -dev" ]
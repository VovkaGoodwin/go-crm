FROM golang:1.23-bookworm AS go

ARG UID=22
ARG GID=501

RUN addgroup --gid $GID app \
    && adduser --uid $UID --gid $GID --gecos 'app' app \
    && echo "web all=(ALL) NOPASSWD: ALL" >> /etc/sudoers \
    && mkdir /app \
    && chown app:app -R /app

WORKDIR /app
USER app

COPY --chown=app:app go.mod go.sum ./

RUN go mod download

COPY --chown=app:app . .

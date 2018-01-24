# Build Stage
FROM lacion/docker-alpine:gobuildimage AS build-stage

LABEL app="build-ScrumPoker"
LABEL REPO="https://github.com/jata84/ScrumPoker"

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PROJPATH=/gopath/src/github.com/jata84/ScrumPoker

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /gopath/src/github.com/jata84/ScrumPoker
WORKDIR /gopath/src/github.com/jata84/ScrumPoker

RUN make build-alpine

# Final Stage
FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/jata84/ScrumPoker"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/ScrumPoker/bin

WORKDIR /opt/ScrumPoker/bin

COPY --from=build-stage /gopath/src/github.com/jata84/ScrumPoker/bin/ScrumPoker /opt/ScrumPoker/bin/
RUN chmod +x /opt/ScrumPoker/bin/ScrumPoker

CMD /opt/ScrumPoker/bin/ScrumPoker
FROM golang

ARG ACCESS_KEY
ARG BRANCH

WORKDIR /github/workspace

ADD ./sonar-scanner /github/workspace/sonar-scanner

LABEL "com.github.actions.name"="Scan Test"
LABEL "com.github.actions.description"="Running scan test"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="http://github.com/unbounce/paw"
LABEL "homepage"="http://github.com/unbounce/paw"
LABEL "maintainer"="UnbounceRabbit <unbouncerabbit@github.com>"

RUN export PATH=$PATH:/github/workspace/scanner/bin

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

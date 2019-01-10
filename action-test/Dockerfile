FROM golang

ADD ./sonar /github/workspace/sonar
WORKDIR /github/workspace/


LABEL "com.github.actions.name"="Scan Test"
LABEL "com.github.actions.description"="Running scan test"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="http://github.com/unbounce/paw"
LABEL "homepage"="http://github.com/unbounce/paw"
LABEL "maintainer"="UnbounceRabbit <unbouncerabbit@github.com>"

EXPOSE 8080
RUN pwd
RUN ./sonar/sonar-scanner --version

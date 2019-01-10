FROM java:openjdk-8-alpine

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
RUN echo $(pwd)
RUN cd sonar/bin/ && chmod +x ./sonar-scanner && ./sonar-scanner --version

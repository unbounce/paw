FROM golang

WORKDIR /github/workspace/sonar-scanner

ADD ./sonar /github/workspace/sonar-scanner

LABEL "com.github.actions.name"="Scan Test"
LABEL "com.github.actions.description"="Running scan test"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="http://github.com/unbounce/paw"
LABEL "homepage"="http://github.com/unbounce/paw"
LABEL "maintainer"="UnbounceRabbit <unbouncerabbit@github.com>"

EXPOSE 8080
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

FROM debian:9.5-slim

LABEL "com.github.actions.name"="Scan Test"
LABEL "com.github.actions.description"="Running scan test"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="http://github.com/unbounce/paw"
LABEL "homepage"="http://github.com/unbounce/paw"
LABEL "maintainer"="UnbounceRabbit <unbouncerabbit@github.com>"
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

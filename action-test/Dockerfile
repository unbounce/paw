FROM java:openjdk-8-jre-alpine

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
RUN cd sonar/bin/ && chmod +x ./sonar-scanner
RUN ./sonar-scanner \
    -Dsonar.projectKey=sast-test \
    -Dsonar.organization=unbouncerabbit-github \
    -Dsonar.sources=. \
    -Dsonar.host.url=https://sonarcloud.io \
    -Dsonar.login=$SONAR_LOGIN  \
    -Dsonar.branch.name=develop


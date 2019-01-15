FROM bitnami/minideb:jessie

WORKDIR /github/workspace/

#ADD $SONAR_DOWNLOAD_URL /tmp/scanner.zip
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /tmp/scanner.zip

RUN apt-get update -qq && apt-get install unzip -y

RUN mkdir -p /tmp/sonar && unzip /tmp/scanner.zip -d /tmp/sonar

RUN rm /tmp/scanner.zip

CMD /tmp/sonar/sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner \
    -Dsonar.projectKey=unbounce-paw \
    -Dsonar.organization=unbouncerabbit-github \
    -Dsonar.sources=/github/workspace/ \
    -Dsonar.host.url=https://sonarcloud.io \
    -Dsonar.login=$SONAR_LOGIN \
    -Dsonar.branch.name=github \
    -X

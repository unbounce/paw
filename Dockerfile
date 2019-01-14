FROM openjdk

WORKDIR /github/workspace/

#download zip
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /tmp/scanner.zip

#extract all files
RUN mkdir /tmp/sonar && unzip /tmp/scanner.zip -d /tmp/sonar

CMD /tmp/sonar/sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner \
    -Dsonar.projectKey=unbounce-paw \
    -Dsonar.organization=unbouncerabbit-github \
    -Dsonar.sources=/github/workspace/ \
    -Dsonar.host.url=https://sonarcloud.io \
    -Dsonar.login=$SONAR_LOGIN \
    -Dsonar.branch.name=$GITHUB_REF \
    -X

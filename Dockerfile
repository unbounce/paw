FROM openjdk

WORKDIR /github/workspace/

#download zip
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /github/workspace/

#extract all files
RUN unzip ./sonar-scanner-cli-3.2.0.1227-linux.zip

CMD ./sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner -Dsonar.projectKey=unbounce-paw -Dsonar.organization=unbouncerabbit-github -Dsonar.sources=. -Dsonar.host.url=https://sonarcloud.io -Dsonar.login=${SONAR_LOGIN} -Dsonar.branch.name=github -X

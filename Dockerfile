FROM openjdk

WORKDIR /github/workspace/

#download zip
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /github/workspace/

#copy all files into workdir
COPY ./ .

#extract all files
RUN unzip ./sonar-scanner-cli-3.2.0.1227-linux.zip
RUN ./sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner --version

RUN env

RUN ./sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner -Dsonar.projectKey=unbounce-paw -Dsonar.organization=unbouncerabbit-github -Dsonar.sources=. -Dsonar.host.url=https://sonarcloud.io -Dsonar.login=${SONAR_LOGIN} -Dsonar.branch.name=github -X

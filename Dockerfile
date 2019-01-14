FROM openjdk

WORKDIR /github/workspace/

# #download zip
# ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /tmp/scanner.zip

# #extract all files
# RUN mkdir /tmp/sonar && unzip /tmp/scanner.zip -d /tmp/sonar
# ENV SONAR_DOWNLOAD_USER=https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip
# ENV SONAR_PROJECT_KEY=unbounce-paw
# ENV SONAR_ORG=unbouncerabbit-github

#RUN apt-get update -qq && apt-get install unzip -y

#ADD $SONAR_DOWNLOAD_URL /tmp/scanner.zip
ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /tmp/scanner.zip

RUN mkdir -p /tmp/sonar && unzip /tmp/scanner.zip -d /tmp/sonar

CMD /tmp/sonar/sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner \
    -Dsonar.projectKey=unbounce-paw \
    -Dsonar.organization=unbouncerabbit-github \
    -Dsonar.sources=/github/workspace/ \
    -Dsonar.host.url=https://sonarcloud.io \
    -Dsonar.login=$SONAR_LOGIN \
    -Dsonar.branch.name=github \
    -X


#CMD /tmp/sonar/bin/sonar-scanner \
# CMD /tmp/sonar/sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner \
#   -Dsonar.projectKey=$SONAR_PROJECT_KEY \
#   -Dsonar.organization=$SONAR_ORG \
#   -Dsonar.sources=/github/workspace \
#   -Dsonar.host.url=https://sonarcloud.io \
#   -Dsonar.login=$SONAR_LOGIN \
#   -Dsonar.branch.name=github

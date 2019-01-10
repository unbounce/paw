FROM java:openjdk-8-jre-alpine

WORKDIR /github/workspace/

ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /github/workspace/sonar
RUN unzip sonar-scanner-cli-3.2.0.1227-linux.zip /github/workspace/sonar

RUN mkdir -p /github/workspace/code
COPY ./ /github/workspace/code

RUN chmod +x /github/workspace/sonar/bin/sonar-scanner
RUN ls /github/workspace/
RUN ls /github/workspace/sonar
ENTRYPOINT [ "entrypoint.sh" ]

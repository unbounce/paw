FROM openjdk

WORKDIR /github/workspace/

ADD https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip /github/workspace/

COPY ./ .

RUN ls

RUN unzip ./sonar-scanner-cli-3.2.0.1227-linux.zip
RUN ls
RUN ./sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner --version

ENTRYPOINT [ "entrypoint.sh" ]

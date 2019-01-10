FROM java:openjdk-8-jre-alpine

ADD ./sonar /github/workspace/sonar
WORKDIR /github/workspace/


RUN chmod +x .sonar/bin/sonar-scanner
RUN echo $(ls)


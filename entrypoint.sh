#!/bin/sh

sh -c "sonar-scanner \
  -Dsonar.projectKey=sast-test \
  -Dsonar.organization=unbouncerabbit-github \
  -Dsonar.sources=. \
  -Dsonar.host.url=https://sonarcloud.io \
  -Dsonar.login=$(SONAR_LOGIN)  \
  -Dsonar.branch.name=develop "

workflow "Sonar Scanner" {
    on = "push"
    resolves = ["sonar scan"]
}

action "sonar scan" {
    uses = "docker://unbounce/sonarcloud-github"
    secrets = ["SONAR_LOGIN"]
    env = {
        "SONAR_PROJECT_KEY" = "unbounce-paw"
        "SONAR_ORG" = "unbouncerabbit-github"
    }
}

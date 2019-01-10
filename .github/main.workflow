workflow "Scan Test" {
    on = "push"
    resolves = ["action test"]
}

action "action test" {
    uses = "./"
    secrets = ["SONAR_LOGIN"]
    args = "\"SECRET: $SONAR_LOGIN!\""
}

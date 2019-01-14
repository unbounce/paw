workflow "Scan Test" {
    on = "push"
    resolves = ["sonar test"]
}

action "sonar test" {
    uses = "./"
    secrets = ["SONAR_LOGIN"]
}

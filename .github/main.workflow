workflow "Scan Test" {
    on = "push"
    resolves = ["action test"]
}

action "action test" {
    uses = "./action-test/"
}

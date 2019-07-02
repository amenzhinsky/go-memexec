workflow "Main" {
  on = "push"
  resolves = ["go test"]
}

action "go test" {
  uses = "docker://golang:1.12"
  args = "go test"
}

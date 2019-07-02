workflow "Main" {
  on = "push"
  resolves = ["go test"]
}

action "go test" {
  uses = "docker://golang:1.12-alpine"
  args = "go test"
}

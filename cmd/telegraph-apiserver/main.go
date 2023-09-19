package main

import "github.com/kingofzihua/telegraph/internal/apiserver"

func main() {
	apiserver.NewApp("apiserver").Run()
}

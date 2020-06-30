package main

import (
	_ "github.com/lib/pq"
	"github.com/ngavinsir/clickbait/cmd"
)

//go:generate sqlboiler --wipe psql

func main() {
	cmd.Execute()
}

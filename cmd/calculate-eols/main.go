package main

import "flag"

func main() {
	metadata := flag.String("metadata", "metadata.sqlite3", "Path to metadata database")

	flag.Parse()

	query := "SELECT * FROM sdk_releases
}

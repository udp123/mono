package main
import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

func main () {
    db, err := sql.Open("postgres0", "user=udp123 dbname=mono sslmode=disable")
    if err != nill {
        log.Fatal("Error: data source args not valid")
    }
    err = db.Ping()
    if err != nill {
        log.Fatal("Error: could not establish db connection")
    }
    queryStmt, err := db.Prepare("SELECT name FROM users WHERE id=$1")
    if err != nill {
        log.Fatal(err)
    }
    var name string
    err = queryStmt.QueryRow(15).Scan(&name)
    fmt.Printf(name)
}

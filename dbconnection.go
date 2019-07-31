package dbconnection
import(
	"log"
	_"github.com/go-sql-driver/mysql"
	
	"database/sql"
)
var db *sql.DB
func Connection(dbpath string)*sql.DB{

	Db,err := sql.Open(dbpath)
	
	if err!=nil{
		log.Fatalln(err)
	}
	return Db
}

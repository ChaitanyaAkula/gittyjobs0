package dbconnection
import(
	"log"
	_"github.com/go-sql-driver/mysql"
	//_"mysqlserver"
	"database/sql"
)
var db *sql.DB
func Connection()*sql.DB{

	Db,err := sql.Open("mysql",Dbpath)
	
	if err!=nil{
		log.Fatalln(err)
	}
	return Db
}
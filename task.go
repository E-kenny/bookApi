package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"

)

//Book type
type Book struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Publish_at string `json:"publish_at"`
}

var db *sql.DB
var err error
func loadBooks(w http.ResponseWriter, r *http.Request){
	var allbooks = []Book{}
	rows, err := db.Query("SELECT * FROM bless")
	if err != nil {
		panic(err)
	}else{
		fmt.Fprintln(w, "successfully selected")
	}
	//defer rows.Close()
    //defer db.Close()


	for rows.Next() {
		book:=Book{}
		err := rows.Scan(&book.ID,&book.Name,&book.Author,&book.Publish_at)
		if err != nil {
			panic(err)
		}
		allbooks = append(allbooks,book)

	}

	w.Header().Set("Content-Type","application/json")
	
	bk, err := json.Marshal(allbooks)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	
	w.Write(bk)
	
}


func createBook(w http.ResponseWriter, r *http.Request){
	queries := mux.Vars(r)["id"] 
		fmt.Fprintf(w, "Category: %v\n", queries)
	    
	
	insert, err := db.Prepare("INSERT INTO bless VALUES (?,?,?,?)")
	// insert.Exec(queries,"second","Alchemist","01-11-2020")
	//  if err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	// 	return
	// 	}
	_,err = insert.Exec(queries,"second","Alchemist","01-11-2020")
if err != nil {
panic(err)
} else{
fmt.Println("The value was successfully inserted!")
}
	//defer insert.Close()
		fmt.Fprintf(w, "Record Updated!")
	}
	

func updateBook(w http.ResponseWriter, r *http.Request){
	queries := mux.Vars(r)["id"] 
		fmt.Fprintf(w, "Category: %v\n", queries)
	    

		UpdateStatement :="UPDATE  bless SET  name=?,author=?,publish_at=? WHERE id=?"	
	
		UpdateResult, UpdateResultErr := db.Prepare(UpdateStatement)
		if UpdateResultErr != nil {
			panic(UpdateResultErr)
		}
		//UpdateResult.Exec("ekenny","dragon","02-23-2020",queries)
		_,err = UpdateResult.Exec("fifty","dragon","02-11-2020",queries)
if err != nil {
panic(err)
} else{
fmt.Println("The value was successfully inserted!")
}
		fmt.Fprintf(w,"query %v updated successfully",queries)		
	  // defer db.Close()

			   
}


func deleteBook(w http.ResponseWriter, r *http.Request){	
		queries := mux.Vars(r)["id"]
		fmt.Fprintf(w, "Category: %v\n", queries)
	    
	
	DeleteStatement :="DELETE FROM bless WHERE id =?"	
	
		DeleteResult, DeleteResultErr := db.Prepare(DeleteStatement)
		if DeleteResultErr != nil {
			panic(DeleteResultErr)
		}
		DeleteResult.Exec(queries)
		
		fmt.Fprintf(w,"query %v deleted successfully",queries)
		//defer db.Close()
			
	
	}

	func logger(next http.HandlerFunc)http.HandlerFunc  {
		return func (w http.ResponseWriter,r *http.Request)  {
			fmt.Printf("This request was received from %v",r.URL)
			next(w,r)
		}
	}

func main()  {
	db, err = sql.Open("mysql", "root:Ekenny2468@tcp(localhost:3306)/ekenny")
if err != nil {
	panic(err)
}else{
	fmt.Println("The connection to the DB was successfully initialized ")
}


err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

// DBCreate :=` CREATE  TABLE  bless
//  (
	  
// 	id INT,
// 	name TEXT,
// 	author TEXT,
// 	publish_at TEXT UNIQUE NOT NULL
		  
//  )

//`
//  _, err = db.Exec(DBCreate)
// if err != nil {
// 	panic(err)
// }else{
// 	fmt.Println("The table was successfully created")
// }


router :=mux.NewRouter()
api := router.PathPrefix("/api/v1").Subrouter()

api.HandleFunc("/Books",logger(loadBooks)).Methods(http.MethodGet)
api.HandleFunc("/Books/id/{id}",logger(createBook)).Methods(http.MethodPost)
api.HandleFunc("/Books/id/{id}",logger(updateBook)).Methods(http.MethodPatch)
api.HandleFunc("/Books/id/{id}",logger(deleteBook)).Methods(http.MethodDelete)
fmt.Println("server started successfully")
 log.Fatalln(http.ListenAndServe(":8080",router))
}
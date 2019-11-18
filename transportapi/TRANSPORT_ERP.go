package main

// imports
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// structs
type Stop_Master struct {
	Stop_Id     *int    `json:"StopId"`
	Stop_Name   *string `json:"StopName"`
	Priority_No *int    `json:"PriorityNo"`
	Created_By  *string `json:"CreatedBy"`
	Created_On  *string `json:"CreatedOn"`
	Modified_By *string `json:"ModifiedBy"`
	Modified_On *string `json:"ModifiedOn"`
	Is_Active   *int    `json:"IsActive"`
}
type Route_Master struct {
	Route_Id    *int    `json:"RouteId"`
	Route_No    *int    `json:"RouteNo"`
	Route_Name  *string `json:"RouteName"`
	Priority_No *int    `json:"PriorityNo"`
	Created_By  *string `json:"CreatedBy"`
	Created_On  *string `json:"CreatedOn"`
	Modified_By *string `json:"ModifiedBy"`
	Modified_On *string `json:"ModifiedOn"`
	Is_Active   *int    `json:"IsActive"`
}
type Vehicle_Master struct {
	Vehicle_No    *int    `json:"RouteId"`
	Vehicle_Name  *int    `json:"RouteNo"`
	Seats         *int    `json:"RouteName"`
	Vehicle_Model *string `json:"PriorityNo"`
	Driver_Id     *int    `json:"DriverId "`
	Created_By    *string `json:"CreatedBy"`
	Created_On    *string `json:"CreatedOn"`
	Modified_By   *string `json:"ModifiedBy"`
	Modified_On   *string `json:"ModifiedOn"`
	Is_Active     *int    `json:"IsActive"`
}
type Vehicle_Routes struct {
	Routes_No   *int    `json:"RoutesNo"`
	Vehicle_No  *int    `json:"VehicleNo"`
	Created_By  *string `json:"CreatedBy"`
	Created_On  *string `json:"CreatedOn"`
	Modified_By *string `json:"ModifiedBy"`
	Modified_On *string `json:"ModifiedOn"`
	Is_Active   *int    `json:"IsActive"`
}
type Stop_Group struct {
	Stop_Group_Id *int    `json:"StopGroupId"`
	Stop_Id       *int    `json:"StopId"`
	Created_By    *string `json:"CreatedBy"`
	Created_On    *string `json:"CreatedOn"`
	Modified_By   *string `json:"ModifiedBy"`
	Modified_On   *string `json:"ModifiedOn"`
	Is_Active     *int    `json:"IsActive"`
}
type StandardResponse struct {
	Data  string `json:"Data"`
	Value int    `json:"Value"`
}

var mySigningKey = []byte("manavrachnakey")

// jwt validation
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {

		//fmt.Println(r.Header["Token"])
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error") // nill here
				}
				return mySigningKey, nil
			})

			if err != nil {
				//panic(err.Error())
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
				//fmt.Println("Authorized")
			}
			fmt.Println("test", token)
			//fmt.Fprintf(w, "Authorized")
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testing function")

}
// create stop api
func createstop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createstop called")
	var srs []StandardResponse
	var sr StandardResponse
	defer func() {
		srs = append([]StandardResponse{sr}, srs...)
		json.NewEncoder(w).Encode(srs)
	}()
	sr.Value = 0
	//Set header type and initialize connection string
	w.Header().Set("Content-Type", "application/json")
	connString := "server=localhost;user id=sa;password=ravidev2018;database=transport"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		sr.Value = -1
		return
	}

	//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.

	defer r.Body.Close()

	var inputArray []Stop_Master

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sr.Data = "Error Reading Input"
		sr.Value = -201
		return
	}

	//Decode input json and insert it into s
	err = json.Unmarshal(body, &inputArray)

	if err != nil {
		sr.Data = "Error Parsing Input"
		sr.Value = -202
		return
	}

	sr.Value = 1

	for _, s := range inputArray {
		fmt.Println("Data begin")
		fmt.Println(*s.Stop_Id)
		fmt.Println(*s.Stop_Name)
		fmt.Println(*s.Priority_No)
		fmt.Println(*s.Created_By)
		fmt.Println(*s.Created_On)
		fmt.Println(*s.Modified_By)
		fmt.Println(*s.Modified_On)
		fmt.Println(*s.Is_Active)
		fmt.Println("Data end")
		var tempsr StandardResponse
		tempsr.Value = 0
		tempsr.Data = strconv.Itoa(*s.Stop_Id)
		// tempsr.Data = ""
		//Create and initialize transaction
		fmt.Println("Initializing transaction")
		txn, err := db.Begin()
		if err != nil {
			//sr.Data = err.Error()
			fmt.Println(err.Error())
			tempsr.Data = "The database could not be connected."
			tempsr.Value = -101
			srs = append(srs, tempsr)
			txn.Rollback()
			continue

		}
		fmt.Println("Transaction initialized")
		//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.
		defer txn.Rollback()
		//Query to insert student information data
		query := "insert into stopmaster(StopName,PriorityNo,CreatedBy,CreatedOn,IsActive) values"
		query = query + "(@sname,@pno,@createdby,current_timestamp,1);"

		//Initialize variable vals to contain parameters
		vals := []interface{}{} // empty interface holds any type of value of vals

		//Start appending parameters to vals
		vals = append(vals, sql.Named("sname", s.Stop_Name), sql.Named("pno", s.Priority_No))
		vals = append(vals, sql.Named("createdby", s.Created_By))

		// Execution of vals
		fmt.Println(query)
		res, err := txn.Exec(query, vals...)
		if err != nil {
			fmt.Println(err.Error())
			tempsr.Data = "SQL Query Error"
			fmt.Println("2")
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra := res.RowsAffected
		fmt.Println("The result of first query is ", &ra)

		fmt.Println("Rows affected: ", &ra)
        // logs
		fmt.Println("log query started")
			//Query to insert student information data
			query_log := "insert into stopmaster_log(StopName,PriorityNo,CreatedBy,CreatedOn,IsActive,LogType) values"
			query_log = query_log+ "(@sname,@pno,@createdbylog,current_timestamp,1,1);"
	
			//Initialize variable vals to contain parameters
			vals_log := []interface{}{} // empty interface holds any type of value of vals
	
			//Start appending parameters to vals
			vals_log = append(vals, sql.Named("sname", s.Stop_Name), sql.Named("pno", s.Priority_No))
			vals_log = append(vals, sql.Named("createdbylog", s.Created_By))
	
			// Execution of vals
			fmt.Println(query_log)
			_, err = txn.Exec(query_log, vals_log...)
			if err != nil {
				fmt.Println(err.Error())
				tempsr.Data = "SQL Query_log Error"
				fmt.Println("2")
				tempsr.Value = -301
				srs = append(srs, tempsr)
				txn.Rollback()
				continue
			}
			ra_log := res.RowsAffected
			fmt.Println("The result of first query is ", &ra_log)
	
			fmt.Println("Rows affected: ", &ra_log)

		txn.Commit()
		tempsr.Value = 1
		srs = append(srs, tempsr)
		fmt.Println("Executed successfully")
	}
}


// create routes api
func createroutes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createroute called")
	var srs []StandardResponse
	var sr StandardResponse
	defer func() {
		srs = append([]StandardResponse{sr}, srs...)
		json.NewEncoder(w).Encode(srs)
	}()
	sr.Value = 0
	//Set header type and initialize connection string
	w.Header().Set("Content-Type", "application/json")
	connString := "server=localhost;user id=sa;password=ravidev2018;database=transport"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		sr.Value = -1
		return
	}

	//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.

	defer r.Body.Close()

	var inputArray []Route_Master

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sr.Data = "Error Reading Input"
		sr.Value = -201
		return
	}

	//Decode input json and insert it into s
	err = json.Unmarshal(body, &inputArray)

	if err != nil {
		sr.Data = "Error Parsing Input"
		sr.Value = -202
		return
	}

	sr.Value = 1

	for _, s := range inputArray {
		fmt.Println("Data begin")
		fmt.Println(*s.Route_Id)
		fmt.Println(*s.Route_Name)
		fmt.Println(*s.Route_No)
		fmt.Println(*s.Priority_No)
		fmt.Println(*s.Created_By)
		fmt.Println(*s.Created_On)
		fmt.Println(*s.Modified_By)
		fmt.Println(*s.Modified_On)
		fmt.Println(*s.Is_Active)
		fmt.Println("Data end")
		var tempsr StandardResponse
		tempsr.Value = 0
		tempsr.Data = strconv.Itoa(*s.Route_Id)
		// tempsr.Data = ""
		//Create and initialize transaction
		fmt.Println("Initializing transaction")
		txn, err := db.Begin()
		if err != nil {
			//sr.Data = err.Error()
			fmt.Println(err.Error())
			tempsr.Data = "The database could not be connected."
			tempsr.Value = -101
			srs = append(srs, tempsr)
			txn.Rollback()
			continue

		}
		fmt.Println("Transaction initialized")
		//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.
		defer txn.Rollback()
		//Query to insert student information data
		query := "insert into routemaster(RouteNo,RouteName,PriorityNo,CreatedBy,CreatedOn,IsActive) values"
		query = query + "(@rno,@rname,@pno,@createdby,current_timestamp,1);"

		//Initialize variable vals to contain parameters
		vals := []interface{}{} // empty interface holds any type of value of vals

		//Start appending parameters to vals
		vals = append(vals, sql.Named("rno", s.Route_No),sql.Named("rname", s.Route_Name), sql.Named("pno", s.Priority_No))
		vals = append(vals, sql.Named("createdby", s.Created_By))

		// Execution of vals
		fmt.Println(query)
		res, err := txn.Exec(query, vals...)
		if err != nil {
			fmt.Println(err.Error())
			tempsr.Data = "SQL Query Error"
			fmt.Println("2")
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra := res.RowsAffected
		fmt.Println("The result of first query is ", &ra)

		fmt.Println("Rows affected: ", &ra)
        // logs
		fmt.Println("log query started")
			//Query to insert student information data
			query_log := "insert into routemaster_log(RouteNo,RouteName,PriorityNo,CreatedBy,CreatedOn,IsActive,LogType) values"
			query_log = query_log+ "(@rno,@rname,@pno,@createdby,current_timestamp,1,1);"
	
			//Initialize variable vals to contain parameters
			vals_log := []interface{}{} // empty interface holds any type of value of vals
	
			//Start appending parameters to vals
			vals_log = append(vals, sql.Named("rno", s.Route_No),sql.Named("rname", s.Route_Name), sql.Named("pno", s.Priority_No))
			vals_log = append(vals, sql.Named("createdbylog", s.Created_By))
	
			// Execution of vals
			fmt.Println(query_log)
			_, err = txn.Exec(query_log, vals_log...)
			if err != nil {
				fmt.Println(err.Error())
				tempsr.Data = "SQL Query_log Error"
				fmt.Println("2")
				tempsr.Value = -301
				srs = append(srs, tempsr)
				txn.Rollback()
				continue
			}
			ra_log := res.RowsAffected
			fmt.Println("The result of first query is ", &ra_log)
	
			fmt.Println("Rows affected: ", &ra_log)

		txn.Commit()
		tempsr.Value = 1
		srs = append(srs, tempsr)
		fmt.Println("Executed successfully")
	}

}

// create vehicle api
func createvehicle(w http.ResponseWriter, r *http.Request) {
}


// handler
func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test).Methods("POST")
	router.HandleFunc("/createstop", createstop).Methods("POST")
	router.HandleFunc("/createroutes", createroutes).Methods("POST")
	router.HandleFunc("/createvehicle", createvehicle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))
}

// main function
func main() {
	handleRequests()
}

// jwt token generation
func GenerateJWT() {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		// fmt.Errorf("Something Went Wrong: %s" + err.Error())
		// return
	}

	fmt.Println("Token is ", tokenString)
}

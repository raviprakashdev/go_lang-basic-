package main

// imports
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

var mySigningKey = []byte("captainjacksparrowsayshi")

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


// api createstop
func createstop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createstop called")
	fmt.Fprintf(w,"api createstop called")

	// var srs []StandardResponse
	var sr StandardResponse
	// defer func() {
	// 	srs = append([]StandardResponse{sr}, srs...)
	// 	json.NewEncoder(w).Encode(srs)
	// }()
	// sr.Value = 0
	// //Set header type and initialize connection string
	// w.Header().Set("Content-Type", "application/json")
	// connString := "server=localhost;userid=sa;password=ravidev2018;port=1433;database=stopmaster"
	// db, err := sql.Open("sqlserver", connString)
	// if err != nil {
	// 	sr.Value = -1
	// 	return
	// }

	// //This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.

	defer r.Body.Close()

	var inputArray []Stop_Master

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sr.Data = "Error Reading Input"
		sr.Value = -201
		return
	}

	// //Decode input json and insert it into s
	err = json.Unmarshal(body, &inputArray)

	if err != nil {
		sr.Data = "Error Parsing Input"
		sr.Value = -202
		return
	}

	sr.Value = 1
	fmt.Println(len(inputArray))
	for _, s := range inputArray {
		fmt.Println(*s.Stop_Name)
		// var tempsr StandardResponse
		// tempsr.Value = 0
		// tempsr.Data = strconv.Itoa(*s.StudentInformation.Application_Number)
		//Create and initialize transaction
		// txn, err := db.Begin()
		// if err != nil {
			//sr.Data = err.Error()
			// tempsr.Data = "The database could not be connected."
			// tempsr.Value = -101
			// srs = append(srs, tempsr)
			// txn.Rollback()
			continue

		}
	

	// 	//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.
	// 	defer txn.Rollback()
	// 	//Query to insert student information data
	// 	query := "insert into reg_student_information(ApplicationNumber,ApplicationStatus,FirstName,MiddleName,LastName,Class,DOB,AppliedBatch,"
	// 	query = query + "StaffWard,ResidentialCategory,AdmissionType,Medical,MedicalPrecaution,Nationality,Gender,BloddGroup,EmergencyNumber,"
	// 	query = query + "MotherTongue,AadharNumber,Religion,SocialCategory,Locality,BirthPlace,Transport,PreviousSchool,Allergy,"
	// 	query = query + "AllergyPrecaution,StudentStatus,PaymentStatus,ProfilePicture,AppliedBy,InstituteId,CreatedBy,CreatedOn) values"
	// 	query = query + "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp());"

	// 	//Query to insert student information data into log
	// 	queryLog := "insert into reg_student_information_log(ApplicationNumber,ApplicationStatus,FirstName,MiddleName,LastName,Class,DOB,AppliedBatch,"
	// 	queryLog = queryLog + "StaffWard,ResidentialCategory,AdmissionType,Medical,MedicalPrecaution,Nationality,Gender,BloddGroup,EmergencyNumber,"
	// 	queryLog = queryLog + "MotherTongue,AadharNumber,Religion,SocialCategory,Locality,BirthPlace,Transport,PreviousSchool,Allergy,"
	// 	queryLog = queryLog + "AllergyPrecaution,StudentStatus,PaymentStatus,ProfilePicture,AppliedBy,InstituteId,CreatedBy,CreatedOn,Timestamp,OperationType) values"
	// 	queryLog = queryLog + "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp(),1);"

	// 	//Initialize variable vals to contain parameters
	// 	vals := []interface{}{} // empty interface holds any type of value of vals
	// 	valslog := []interface{}{}

	// 	//Start appending parameters to vals
	// 	vals = append(vals, s.StudentInformation.Application_Number, s.StudentInformation.Application_Status, s.StudentInformation.First_Name, s.StudentInformation.Middle_Name)
	// 	vals = append(vals, s.StudentInformation.Last_Name, s.StudentInformation.Class, s.StudentInformation.DOB, s.StudentInformation.Batch)
	// 	vals = append(vals, s.StudentInformation.Staff_Ward, s.StudentInformation.Residential_Category, s.StudentInformation.Admission_Type)
	// 	vals = append(vals, s.StudentInformation.Medical, s.StudentInformation.Medical_Precaution, s.StudentInformation.Nationality)
	// 	vals = append(vals, s.StudentInformation.Gender, s.StudentInformation.Blood_Group, s.StudentInformation.Emergency_Number)
	// 	vals = append(vals, s.StudentInformation.Mother_Tongue, s.StudentInformation.Aadhar_Number, s.StudentInformation.Religion)
	// 	vals = append(vals, s.StudentInformation.Social_Category, s.StudentInformation.Locality, s.StudentInformation.Birth_Place)
	// 	vals = append(vals, s.StudentInformation.Transport_Facilities, s.StudentInformation.Previous_School)
	// 	vals = append(vals, s.StudentInformation.Student_Allergy, s.StudentInformation.Allergy_Precaution)
	// 	vals = append(vals, s.StudentInformation.Student_Status, s.StudentInformation.Payment_Status, s.StudentInformation.Profile_Picture)
	// 	vals = append(vals, s.StudentInformation.Applied_By, s.StudentInformation.Institute_Id, s.StudentInformation.Created_By)

	// 	//Start appending parameters to valslog

	// 	valslog = append(valslog, s.StudentInformation.Application_Number, s.StudentInformation.Application_Status, s.StudentInformation.First_Name, s.StudentInformation.Middle_Name)
	// 	valslog = append(valslog, s.StudentInformation.Last_Name, s.StudentInformation.Class, s.StudentInformation.DOB, s.StudentInformation.Batch)
	// 	valslog = append(valslog, s.StudentInformation.Staff_Ward, s.StudentInformation.Residential_Category, s.StudentInformation.Admission_Type)
	// 	valslog = append(valslog, s.StudentInformation.Medical, s.StudentInformation.Medical_Precaution, s.StudentInformation.Nationality)
	// 	valslog = append(valslog, s.StudentInformation.Gender, s.StudentInformation.Blood_Group, s.StudentInformation.Emergency_Number)
	// 	valslog = append(valslog, s.StudentInformation.Mother_Tongue, s.StudentInformation.Aadhar_Number, s.StudentInformation.Religion)
	// 	valslog = append(valslog, s.StudentInformation.Social_Category, s.StudentInformation.Locality, s.StudentInformation.Birth_Place)
	// 	valslog = append(valslog, s.StudentInformation.Transport_Facilities, s.StudentInformation.Previous_School)
	// 	valslog = append(valslog, s.StudentInformation.Student_Allergy, s.StudentInformation.Allergy_Precaution)
	// 	valslog = append(valslog, s.StudentInformation.Student_Status, s.StudentInformation.Payment_Status, s.StudentInformation.Profile_Picture)
	// 	valslog = append(valslog, s.StudentInformation.Applied_By, s.StudentInformation.Institute_Id, s.StudentInformation.Created_By)

	// 	// Execution of vals

	// 	res, err := txn.Exec(query, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra := res.RowsAffected
	// 	fmt.Println("The result of first query is ", &ra)

	// 	fmt.Println("Rows affected: ", &ra)

	// 	// Execution of valslog

	// 	res, err = txn.Exec(queryLog, valslog...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		//fmt.Println(err.Error())
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("The result of first query is ", &ra)

	// 	fmt.Println("Rows affected: ", &ra)

	// 	//Guardian details. The process is same as above
	// 	query2 := "insert into reg_guardian_details(ApplicationNumber,LocalGuardianName,LocalGuardiaAddress,GuardiaMobileNumber,"
	// 	query2 = query2 + "EmergencyContactPerson,GuardianPhotograph,CreatedBy,CreatedOn) values(?,?,?,?,?,?,?,current_timestamp());"

	// 	//Log Guardian details. The process is same as above
	// 	logquery2 := "insert into reg_guardian_details_log(ApplicationNumber,LocalGuardianName,LocalGuardiaAddress,GuardiaMobileNumber,"
	// 	logquery2 = logquery2 + "EmergencyContactPerson,GuardianPhotograph,CreatedBy,CreatedOn,Timestamp,OperationType) values(?,?,?,?,?,?,?,current_timestamp(),"
	// 	logquery2 = logquery2 + "current_timestamp(),1);"

	// 	fmt.Println(query2)
	// 	fmt.Println(logquery2)

	// 	vals = []interface{}{}

	// 	valslog = []interface{}{}

	// 	vals = append(vals, s.StudentInformation.Application_Number, s.GuardianDetails.Guardian_Name, s.GuardianDetails.Guardian_Address, s.GuardianDetails.Mobile_Number)
	// 	vals = append(vals, s.GuardianDetails.Emergency_Contact_Person, s.GuardianDetails.Photograph, s.GuardianDetails.Created_By)

	// 	valslog = append(valslog, s.StudentInformation.Application_Number, s.GuardianDetails.Guardian_Name, s.GuardianDetails.Guardian_Address, s.GuardianDetails.Mobile_Number)
	// 	valslog = append(valslog, s.GuardianDetails.Emergency_Contact_Person, s.GuardianDetails.Photograph, s.GuardianDetails.Created_By)

	// 	//fmt.Println(query2)
	// 	res, err = txn.Exec(query2, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	//fmt.Println(logquery2)

	// 	res, err = txn.Exec(logquery2, valslog...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	//Guardian Languages.

	// 	vals = []interface{}{}

	// 	valslog = []interface{}{}

	// 	glquery := "insert into reg_guardian_languages(ApplicationNumber,Language,CreatedBy,CreatedOn) values"
	// 	//(last_insert_id(),?,?,current_timestamp());
	// 	for _, gl := range s.GuardianDetails.Languages {
	// 		glquery += "(?,?,?,current_timestamp()),"
	// 		vals = append(vals, s.StudentInformation.Application_Number, gl.Language, gl.CreatedBy)
	// 	}
	// 	glquery = glquery[0 : len(glquery)-1]
	// 	glquery += ";"
	// 	fmt.Println(glquery)

	// 	//log

	// 	glquerylog := "insert into reg_guardian_languages_log(ApplicationNumber,Language,CreatedBy,CreatedOn,Timestamp,OperationType) values"
	// 	//(last_insert_id(),?,?,current_timestamp());
	// 	for _, gl := range s.GuardianDetails.Languages {
	// 		glquerylog += "(?,?,?,current_timestamp(),current_timestamp(),1),"
	// 		valslog = append(valslog, s.StudentInformation.Application_Number, gl.Language, gl.CreatedBy)
	// 	}
	// 	glquerylog = glquerylog[0 : len(glquerylog)-1]
	// 	glquerylog += ";"
	// 	fmt.Println(glquerylog)

	// 	res, err = txn.Exec(glquery, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	// for log

	// 	res, err = txn.Exec(glquerylog, valslog...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	//Guardian activities.

	// 	vals = []interface{}{}

	// 	valslog = []interface{}{}

	// 	gaquery := "insert into reg_guardian_activities(ApplicationNumber,Activity,CreatedBy,CreatedOn) values"
	// 	//(last_insert_id(),?,?,current_timestamp());
	// 	for _, ga := range s.GuardianDetails.Activities {
	// 		gaquery += "(?,?,?,current_timestamp()),"
	// 		vals = append(vals, s.StudentInformation.Application_Number, ga.Activity, ga.CreatedBy)
	// 	}
	// 	gaquery = gaquery[0 : len(gaquery)-1]
	// 	gaquery += ";"
	// 	fmt.Println(gaquery)

	// 	// for log

	// 	gaquerylog := "insert into reg_guardian_activities_log(ApplicationNumber,Activity,CreatedBy,CreatedOn,Timestamp,OperationType) values"
	// 	//(last_insert_id(),?,?,current_timestamp());
	// 	for _, ga := range s.GuardianDetails.Activities {
	// 		gaquerylog += "(?,?,?,current_timestamp(),current_timestamp(),1),"
	// 		valslog = append(valslog, s.StudentInformation.Application_Number, ga.Activity, ga.CreatedBy)
	// 	}
	// 	gaquerylog = gaquerylog[0 : len(gaquerylog)-1]
	// 	gaquerylog += ";"
	// 	fmt.Println(gaquerylog)

	// 	res, err = txn.Exec(gaquery, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	// for log

	// 	res, err = txn.Exec(gaquerylog, valslog...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	//Secondary Exam. The process is same
	// 	query3 := "insert into reg_secondary_exam_details(ApplicationNumber,PassingYear,SchoolName,SchoolAddress,Schoolstatus,YearAttended,"
	// 	query3 = query3 + "BoardRollNumber,ReasonNotTakingAdmissionInSameSchool,Board,ResultInPercentage,Location,Preference1,Preference2,"
	// 	query3 = query3 + "Preference3,CreatedBy,CreatedOn) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,Current_Timestamp());"

	// 	//Secondary Exam Log. The process is same
	// 	query3log := "insert into reg_secondary_exam_details_log(ApplicationNumber,PassingYear,SchoolName,SchoolAddress,Schoolstatus,YearAttended,"
	// 	query3log = query3log + "BoardRollNumber,ReasonNotTakingAdmissionInSameSchool,Board,ResultInPercentage,Location,Preference1,Preference2,"
	// 	query3log = query3log + "Preference3,CreatedBy,CreatedOn,Timestamp,OperationType) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,Current_Timestamp(),Current_Timestamp(),1);"

	// 	vals = []interface{}{}

	// 	valslog = []interface{}{}

	// 	vals = append(vals, s.StudentInformation.Application_Number, s.SecondaryExam.Passing_Year, s.SecondaryExam.School_Name, s.SecondaryExam.School_Address)
	// 	vals = append(vals, s.SecondaryExam.School_status, s.SecondaryExam.Year_Attended, s.SecondaryExam.Board_Roll_Number)
	// 	vals = append(vals, s.SecondaryExam.Reason_not_taking_Admission_in_same_School, s.SecondaryExam.Board)
	// 	vals = append(vals, s.SecondaryExam.Result_in_percentage, s.SecondaryExam.Location, s.SecondaryExam.Preference_1)
	// 	vals = append(vals, s.SecondaryExam.Preference_2, s.SecondaryExam.Preference_3, s.SecondaryExam.Created_By)

	// 	valslog = append(valslog, s.StudentInformation.Application_Number, s.SecondaryExam.Passing_Year, s.SecondaryExam.School_Name, s.SecondaryExam.School_Address)
	// 	valslog = append(valslog, s.SecondaryExam.School_status, s.SecondaryExam.Year_Attended, s.SecondaryExam.Board_Roll_Number)
	// 	valslog = append(valslog, s.SecondaryExam.Reason_not_taking_Admission_in_same_School, s.SecondaryExam.Board)
	// 	valslog = append(valslog, s.SecondaryExam.Result_in_percentage, s.SecondaryExam.Location, s.SecondaryExam.Preference_1)
	// 	valslog = append(valslog, s.SecondaryExam.Preference_2, s.SecondaryExam.Preference_3, s.SecondaryExam.Created_By)

	// 	fmt.Println(query3)

	// 	fmt.Println(query3log)

	// 	res, err = txn.Exec(query3, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	// for log

	// 	res, err = txn.Exec(query3log, valslog...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	//Address Info. The process is same.
	// 	vals = []interface{}{}

	// 	valslog = []interface{}{}

	// 	addressQuery := "insert into reg_address_information(ApplicationNumber,Address,Pincode,City,State,Country,PhoneNumber,AddressType,"
	// 	addressQuery += "CreatedBy,CreatedOn) values"
	// 	for _, ai := range s.AddressInformation {
	// 		addressQuery += "(?,?,?,?,?,?,?,?,?,current_timestamp()),"
	// 		vals = append(vals, s.StudentInformation.Application_Number, ai.Address, ai.Pincode, ai.City, ai.State, ai.Country, ai.Phone_Number, ai.Address_Type, ai.Created_By)
	// 	}
	// 	// for log
	// 	addressQuerylog := "insert into reg_address_information_log(ApplicationNumber,Address,Pincode,City,State,Country,PhoneNumber,AddressType,"
	// 	addressQuerylog += "CreatedBy,CreatedOn,Timestamp,OperationType) values"
	// 	for _, ai := range s.AddressInformation {
	// 		addressQuerylog += "(?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp,1),"
	// 		valslog = append(valslog, s.StudentInformation.Application_Number, ai.Address, ai.Pincode, ai.City, ai.State, ai.Country, ai.Phone_Number, ai.Address_Type, ai.Created_By)
	// 	}
	// 	addressQuery = addressQuery[0 : len(addressQuery)-1]
	// 	addressQuery += ";"
	// 	res, err = txn.Exec(addressQuery, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}

	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	addressQuerylog = addressQuerylog[0 : len(addressQuerylog)-1]
	// 	addressQuerylog += ";"
	// 	res, err = txn.Exec(addressQuerylog, valslog...)

	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	//Parents info
	// 	vals = []interface{}{}

	// 	valslog = []interface{}{}

	// 	parentsquery := "insert into reg_parent_information(ApplicationNumber,Name,Qualification,College,CourseName,CollegePassingYear,WorkingStatus,"
	// 	parentsquery += "EmploymentType,Designation,Occupation,MobileNumber,EmailID,Organization,AnnualIncome,OrganizationAddress,"
	// 	parentsquery += "SchoolAlumini,SchoolPassingYear,CreatedBy,ProfilePicture,ParentType,Salutation,CreatedOn) values"

	// 	for _, pi := range s.ParentInformation {
	// 		parentsquery += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp()),"
	// 		vals = append(vals, s.StudentInformation.Application_Number, pi.Name, pi.Qualification, pi.College, pi.Course_Name, pi.College_Passing_Year, pi.Working_Status, pi.Employment_Type)
	// 		vals = append(vals, pi.Designation, pi.Occupation, pi.Mobile_Number, pi.Email_ID, pi.Organization, pi.Annual_Income, pi.Organization_Address)
	// 		vals = append(vals, pi.School_Alumini, pi.School_Passing_year, pi.Created_By, pi.Profile_Picture, pi.Parent_Type, pi.Salutation)
	// 	}

	// 	// for log

	// 	parentsquerylog := "insert into reg_parent_information_log(ApplicationNumber,Name,Qualification,College,CourseName,CollegePassingYear,WorkingStatus,"
	// 	parentsquerylog += "EmploymentType,Designation,Occupation,MobileNumber,EmailID,Organization,AnnualIncome,OrganizationAddress,"
	// 	parentsquerylog += "SchoolAlumini,SchoolPassingYear,CreatedBy,ProfilePicture,ParentType,Salutation,CreatedOn,TimeStamp,OperationType) values"

	// 	for _, pi := range s.ParentInformation {
	// 		parentsquerylog += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp(),1),"
	// 		valslog = append(valslog, s.StudentInformation.Application_Number, pi.Name, pi.Qualification, pi.College, pi.Course_Name, pi.College_Passing_Year, pi.Working_Status, pi.Employment_Type)
	// 		valslog = append(valslog, pi.Designation, pi.Occupation, pi.Mobile_Number, pi.Email_ID, pi.Organization, pi.Annual_Income, pi.Organization_Address)
	// 		valslog = append(valslog, pi.School_Alumini, pi.School_Passing_year, pi.Created_By, pi.Profile_Picture, pi.Parent_Type, pi.Salutation)
	// 	}
	// 	parentsquery = parentsquery[0 : len(parentsquery)-1]
	// 	parentsquery += ";"
	// 	fmt.Println(parentsquery)

	// 	res, err = txn.Exec(parentsquery, vals...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	// for log
	// 	parentsquerylog = parentsquerylog[0 : len(parentsquerylog)-1]
	// 	parentsquerylog += ";"
	// 	fmt.Println(parentsquerylog)
	// 	res, err = txn.Exec(parentsquerylog, valslog...)
	// 	if err != nil {
	// 		tempsr.Data = "SQL Query Error"
	// 		tempsr.Value = -301
	// 		srs = append(srs, tempsr)
	// 		txn.Rollback()
	// 		continue
	// 	}
	// 	ra = res.RowsAffected
	// 	fmt.Println("Rows affected: ", &ra)

	// 	txn.Commit()
	// 	tempsr.Value = 1
	// 	srs = append(srs, tempsr)
	// }
}



// handler
func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/createstop", createstop).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// main function
func main() {
	handleRequests()
}

// generate JWT tokens
func GenerateJWT() {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s" + err.Error())
		//return "", err
	}

	fmt.Println("Token is ", tokenString)
}


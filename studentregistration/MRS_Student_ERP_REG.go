package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Student_Details

type Student_Details struct {
	StudentInformation Student_Information   `json:"StudentInformation"`
	ParentInformation  []Parent_Information  `json:"ParentInformation"`
	AddressInformation []Address_Information `json:"AddressInformation"`
	GuardianDetails    Guardian_Details      `json:"GuardianDetails"`
	SecondaryExam      Secondary_Exam        `json:"SecondaryExam"`
}

type Student_Information struct {
	Application_Number   *int    `json:"ApplicationNumber"`
	Application_Status   *int    `json:"ApplicationStatus"`
	First_Name           *string `json:"FirstName"`
	Middle_Name          *string `json:"MiddleName"`
	Last_Name            *string `json:"LastName"`
	Class                *int    `json:"Class"`
	DOB                  *string `json:"Dob"`
	Batch                *int    `json:"Batch"`
	Staff_Ward           *int    `json:"StaffWard"`
	Residential_Category *int    `json:"ResidentialCategory"`
	Admission_Type       *int    `json:"AdmissionType"`
	Medical              *string `json:"Medical"`
	Medical_Precaution   *string `json:"MedicalPrecaution"`
	Nationality          *int    `json:"Nationality"`
	Gender               *int    `json:"Gender"`
	Blood_Group          *int    `json:"BloodGroup"`
	Emergency_Number     *string `json:"EmergencyNumber"`
	Mother_Tongue        *int    `json:"MotherTongue"`
	Religion             *string `json:"Religion"`
	Aadhar_Number        *string `json:"AadharNumber"`
	Social_Category      *int    `json:"SocialCategory"`
	Locality             *string `json:"Locality"`
	Birth_Place          *string `json:"BirthPlace"`
	Transport_Facilities *int    `json:"TransportFacilities"`
	Previous_School      *string `json:"PreviousSchool"`
	Student_Allergy      *string `json:"StudentAllergy"`
	Allergy_Precaution   *string `json:"AllergyPrecaution"`
	Student_Status       *int    `json:"StudentStatus"`
	Payment_Status       *int    `json:"PaymentStatus"`
	Profile_Picture      *string `json:"ProfilePicture"`
	Applied_By           *int    `json:"AppliedBy"`
	Institute_Id         *int    `json:"InstituteId"`
	Created_By           *int    `json:"CreatedBy"`
	Created_On           *string `json:"CreatedOn"`
	Modified_By          *int    `json:"ModifiedBy"`
	Modified_On          *string `json:"ModifiedOn"`
}
type Parent_Information struct {
	Application_Number   *int    `json:"ApplicationNumber"`
	Name                 *string `json:"Name"`
	Qualification        *int    `json:"Qualification"`
	College              *string `json:"College"`
	Course_Name          *string `json:"CourseName"`
	College_Passing_Year *int    `json:"CollegPassingYear"`
	Working_Status       *int    `json:"WorkingStatus"`
	Employment_Type      *int    `json:"EmploymentType"`
	Designation          *string `json:"Designation"`
	Mobile_Number        *string `json:"MobileNumber"`
	Email_ID             *string `json:"EmailId"`
	Organization         *string `json:"Organization"`
	Annual_Income        *int    `json:"AnnualIncome"`
	Organization_Address *string `json:"OrganinationAddress"`
	School_Alumini       *int    `json:"SchoolAlumini"`
	School_Passing_year  *int    `json:"SchoolPassingYear"`
	Created_By           *int    `json:"CreatedBy"`
	Created_On           *string `json:"CreatedOn"`
	Modified_By          *int    `json:"ModifiedBy"`
	Modified_On          *string `json:"ModifiedOn"`
	Profile_Picture      *string `json:"ProfilePicture"`
	Parent_Type          *int    `json:"ParentType"`
	Salutation           *int    `json:"Salutation"`
	Occupation           *string `json:"Occupation"`
}
type Address_Information struct {
	Application_Number *int    `json:"ApplicationNumber"`
	Address            *string `json:"Address"`
	Pincode            *int    `json:"Pincode"`
	City               *string `json:"City"`
	State              *string `json:"State"`
	Country            *int    `json:"Country"`
	Phone_Number       *string `json:"PhoneNumber"`
	Address_Type       *int    `json:"AddressType"`
	Created_By         *int    `json:"CreatedBy"`
	Created_On         *string `json:"CreatedOn"`
	Modified_By        *int    `json:"ModifiedBy"`
	Modified_On        *string `json:"ModifiedOn"`
}
type Guardian_Details struct {
	Application_Number       *int                  `json:"ApplicationNumber"`
	Guardian_Name            *string               `json:"GuardianName"`
	Guardian_Address         *string               `json:"GuardianAddress"`
	Mobile_Number            *string               `json:"MobileNumber"`
	Activities               []Guardian_Activities `json:"Activities"`
	Languages                []Guardian_Languages  `json:"Languages"`
	Emergency_Contact_Person *int                  `json:"EmergencyContactPerson"`
	Photograph               *string               `json:"Photograph"`
	Created_By               *int                  `json:"CreatedBy"`
	Created_On               *string               `json:"CreatedOn"`
	Modified_By              *int                  `json:"ModifiedBy"`
	Modified_On              *string               `json:"ModifiedOn"`
}
type Secondary_Exam struct {
	Application_Number                         *int     `json:"ApplicationNumber"`
	Passing_Year                               *int     `json:"PassingYear"`
	School_Name                                *string  `json:"SchoolName"`
	School_Address                             *string  `json:"SchoolAddress"`
	School_status                              *int     `json:"Schoolstatus"`
	Year_Attended                              *int     `json:"YearAttended"`
	Board_Roll_Number                          *string  `json:"BoardRollNumber"`
	Reason_not_taking_Admission_in_same_School *string  `json:"ReasonNotTakingAdmissionInSameSchool"`
	Board                                      *int     `json:"Board"`
	Result_in_percentage                       *float64 `json:"ResultInPercentage"`
	Location                                   *string  `json:"Location"`
	Preference_1                               *int     `json:"Preference1"`
	Preference_2                               *int     `json:"Preference2"`
	Preference_3                               *int     `json:"Preference3"`
	Created_By                                 *int     `json:"CreatedBy"`
	Created_On                                 *string  `json:"CreatedOn"`
	Modified_By                                *int     `json:"ModifiedBy"`
	Modified_On                                *string  `json:"ModifiedOn"`
}
type StandardResponse struct {
	Data  string `json:"Data"`
	Value int    `json:"Value"`
}
type Guardian_Languages struct {
	Application_Number *int    `json:"ApplicationNumber"`
	Language           *int    `json:"Language"`
	CreatedBy          *int    `json:"CreatedBy"`
	CreatedOn          *string `json:"CreatedOn"`
}
type Guardian_Activities struct {
	Application_Number *int    `json:"ApplicationNumber"`
	Activity           *int    `json:"Activity"`
	CreatedBy          *int    `json:"CreatedBy"`
	CreatedOn          *string `json:"CreatedOn"`
}

// JWt Token

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")

}

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

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register called")
	var srs []StandardResponse
	var sr StandardResponse
	defer func() {
		srs = append([]StandardResponse{sr}, srs...)
		json.NewEncoder(w).Encode(srs)
	}()
	sr.Value = 0
	//Set header type and initialize connection string
	w.Header().Set("Content-Type", "application/json")
	connString := "server=localhost;userid=sa;password=ravidev2018;port=1433;database=transport"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		sr.Value = -1
		return
	}

	//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.

	defer r.Body.Close()

	var inputArray []Student_Details

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
		var tempsr StandardResponse
		tempsr.Value = 0
		tempsr.Data = strconv.Itoa(*s.StudentInformation.Application_Number)
		//Create and initialize transaction
		txn, err := db.Begin()
		if err != nil {
			//sr.Data = err.Error()
			tempsr.Data = "The database could not be connected."
			tempsr.Value = -101
			srs = append(srs, tempsr)
			txn.Rollback()
			continue

		}

		//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.
		defer txn.Rollback()
		//Query to insert student information data
		query := "insert into stopmaster(ApplicationNumber,ApplicationStatus,FirstName,MiddleName,LastName,Class,DOB,AppliedBatch,"
		query = query + "StaffWard,ResidentialCategory,AdmissionType,Medical,MedicalPrecaution,Nationality,Gender,BloddGroup,EmergencyNumber,"
		query = query + "MotherTongue,AadharNumber,Religion,SocialCategory,Locality,BirthPlace,Transport,PreviousSchool,Allergy,"
		query = query + "AllergyPrecaution,StudentStatus,PaymentStatus,ProfilePicture,AppliedBy,InstituteId,CreatedBy,CreatedOn) values"
		query = query + "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp());"

		//Query to insert student information data into log
		queryLog := "insert into reg_student_information_log(ApplicationNumber,ApplicationStatus,FirstName,MiddleName,LastName,Class,DOB,AppliedBatch,"
		queryLog = queryLog + "StaffWard,ResidentialCategory,AdmissionType,Medical,MedicalPrecaution,Nationality,Gender,BloddGroup,EmergencyNumber,"
		queryLog = queryLog + "MotherTongue,AadharNumber,Religion,SocialCategory,Locality,BirthPlace,Transport,PreviousSchool,Allergy,"
		queryLog = queryLog + "AllergyPrecaution,StudentStatus,PaymentStatus,ProfilePicture,AppliedBy,InstituteId,CreatedBy,CreatedOn,Timestamp,OperationType) values"
		queryLog = queryLog + "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp(),1);"

		//Initialize variable vals to contain parameters
		vals := []interface{}{} // empty interface holds any type of value of vals
		valslog := []interface{}{}

		//Start appending parameters to vals
		vals = append(vals, s.StudentInformation.Application_Number, s.StudentInformation.Application_Status, s.StudentInformation.First_Name, s.StudentInformation.Middle_Name)
		vals = append(vals, s.StudentInformation.Last_Name, s.StudentInformation.Class, s.StudentInformation.DOB, s.StudentInformation.Batch)
		vals = append(vals, s.StudentInformation.Staff_Ward, s.StudentInformation.Residential_Category, s.StudentInformation.Admission_Type)
		vals = append(vals, s.StudentInformation.Medical, s.StudentInformation.Medical_Precaution, s.StudentInformation.Nationality)
		vals = append(vals, s.StudentInformation.Gender, s.StudentInformation.Blood_Group, s.StudentInformation.Emergency_Number)
		vals = append(vals, s.StudentInformation.Mother_Tongue, s.StudentInformation.Aadhar_Number, s.StudentInformation.Religion)
		vals = append(vals, s.StudentInformation.Social_Category, s.StudentInformation.Locality, s.StudentInformation.Birth_Place)
		vals = append(vals, s.StudentInformation.Transport_Facilities, s.StudentInformation.Previous_School)
		vals = append(vals, s.StudentInformation.Student_Allergy, s.StudentInformation.Allergy_Precaution)
		vals = append(vals, s.StudentInformation.Student_Status, s.StudentInformation.Payment_Status, s.StudentInformation.Profile_Picture)
		vals = append(vals, s.StudentInformation.Applied_By, s.StudentInformation.Institute_Id, s.StudentInformation.Created_By)

		//Start appending parameters to valslog

		valslog = append(valslog, s.StudentInformation.Application_Number, s.StudentInformation.Application_Status, s.StudentInformation.First_Name, s.StudentInformation.Middle_Name)
		valslog = append(valslog, s.StudentInformation.Last_Name, s.StudentInformation.Class, s.StudentInformation.DOB, s.StudentInformation.Batch)
		valslog = append(valslog, s.StudentInformation.Staff_Ward, s.StudentInformation.Residential_Category, s.StudentInformation.Admission_Type)
		valslog = append(valslog, s.StudentInformation.Medical, s.StudentInformation.Medical_Precaution, s.StudentInformation.Nationality)
		valslog = append(valslog, s.StudentInformation.Gender, s.StudentInformation.Blood_Group, s.StudentInformation.Emergency_Number)
		valslog = append(valslog, s.StudentInformation.Mother_Tongue, s.StudentInformation.Aadhar_Number, s.StudentInformation.Religion)
		valslog = append(valslog, s.StudentInformation.Social_Category, s.StudentInformation.Locality, s.StudentInformation.Birth_Place)
		valslog = append(valslog, s.StudentInformation.Transport_Facilities, s.StudentInformation.Previous_School)
		valslog = append(valslog, s.StudentInformation.Student_Allergy, s.StudentInformation.Allergy_Precaution)
		valslog = append(valslog, s.StudentInformation.Student_Status, s.StudentInformation.Payment_Status, s.StudentInformation.Profile_Picture)
		valslog = append(valslog, s.StudentInformation.Applied_By, s.StudentInformation.Institute_Id, s.StudentInformation.Created_By)

		// Execution of vals

		res, err := txn.Exec(query, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra := res.RowsAffected
		fmt.Println("The result of first query is ", &ra)

		fmt.Println("Rows affected: ", &ra)

		// Execution of valslog

		res, err = txn.Exec(queryLog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			//fmt.Println(err.Error())
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("The result of first query is ", &ra)

		fmt.Println("Rows affected: ", &ra)

		//Guardian details. The process is same as above
		query2 := "insert into reg_guardian_details(ApplicationNumber,LocalGuardianName,LocalGuardiaAddress,GuardiaMobileNumber,"
		query2 = query2 + "EmergencyContactPerson,GuardianPhotograph,CreatedBy,CreatedOn) values(?,?,?,?,?,?,?,current_timestamp());"

		//Log Guardian details. The process is same as above
		logquery2 := "insert into reg_guardian_details_log(ApplicationNumber,LocalGuardianName,LocalGuardiaAddress,GuardiaMobileNumber,"
		logquery2 = logquery2 + "EmergencyContactPerson,GuardianPhotograph,CreatedBy,CreatedOn,Timestamp,OperationType) values(?,?,?,?,?,?,?,current_timestamp(),"
		logquery2 = logquery2 + "current_timestamp(),1);"

		fmt.Println(query2)
		fmt.Println(logquery2)

		vals = []interface{}{}

		valslog = []interface{}{}

		vals = append(vals, s.StudentInformation.Application_Number, s.GuardianDetails.Guardian_Name, s.GuardianDetails.Guardian_Address, s.GuardianDetails.Mobile_Number)
		vals = append(vals, s.GuardianDetails.Emergency_Contact_Person, s.GuardianDetails.Photograph, s.GuardianDetails.Created_By)

		valslog = append(valslog, s.StudentInformation.Application_Number, s.GuardianDetails.Guardian_Name, s.GuardianDetails.Guardian_Address, s.GuardianDetails.Mobile_Number)
		valslog = append(valslog, s.GuardianDetails.Emergency_Contact_Person, s.GuardianDetails.Photograph, s.GuardianDetails.Created_By)

		//fmt.Println(query2)
		res, err = txn.Exec(query2, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		//fmt.Println(logquery2)

		res, err = txn.Exec(logquery2, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		//Guardian Languages.

		vals = []interface{}{}

		valslog = []interface{}{}

		glquery := "insert into reg_guardian_languages(ApplicationNumber,Language,CreatedBy,CreatedOn) values"
		//(last_insert_id(),?,?,current_timestamp());
		for _, gl := range s.GuardianDetails.Languages {
			glquery += "(?,?,?,current_timestamp()),"
			vals = append(vals, s.StudentInformation.Application_Number, gl.Language, gl.CreatedBy)
		}
		glquery = glquery[0 : len(glquery)-1]
		glquery += ";"
		fmt.Println(glquery)

		//log

		glquerylog := "insert into reg_guardian_languages_log(ApplicationNumber,Language,CreatedBy,CreatedOn,Timestamp,OperationType) values"
		//(last_insert_id(),?,?,current_timestamp());
		for _, gl := range s.GuardianDetails.Languages {
			glquerylog += "(?,?,?,current_timestamp(),current_timestamp(),1),"
			valslog = append(valslog, s.StudentInformation.Application_Number, gl.Language, gl.CreatedBy)
		}
		glquerylog = glquerylog[0 : len(glquerylog)-1]
		glquerylog += ";"
		fmt.Println(glquerylog)

		res, err = txn.Exec(glquery, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		// for log

		res, err = txn.Exec(glquerylog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		//Guardian activities.

		vals = []interface{}{}

		valslog = []interface{}{}

		gaquery := "insert into reg_guardian_activities(ApplicationNumber,Activity,CreatedBy,CreatedOn) values"
		//(last_insert_id(),?,?,current_timestamp());
		for _, ga := range s.GuardianDetails.Activities {
			gaquery += "(?,?,?,current_timestamp()),"
			vals = append(vals, s.StudentInformation.Application_Number, ga.Activity, ga.CreatedBy)
		}
		gaquery = gaquery[0 : len(gaquery)-1]
		gaquery += ";"
		fmt.Println(gaquery)

		// for log

		gaquerylog := "insert into reg_guardian_activities_log(ApplicationNumber,Activity,CreatedBy,CreatedOn,Timestamp,OperationType) values"
		//(last_insert_id(),?,?,current_timestamp());
		for _, ga := range s.GuardianDetails.Activities {
			gaquerylog += "(?,?,?,current_timestamp(),current_timestamp(),1),"
			valslog = append(valslog, s.StudentInformation.Application_Number, ga.Activity, ga.CreatedBy)
		}
		gaquerylog = gaquerylog[0 : len(gaquerylog)-1]
		gaquerylog += ";"
		fmt.Println(gaquerylog)

		res, err = txn.Exec(gaquery, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		// for log

		res, err = txn.Exec(gaquerylog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		//Secondary Exam. The process is same
		query3 := "insert into reg_secondary_exam_details(ApplicationNumber,PassingYear,SchoolName,SchoolAddress,Schoolstatus,YearAttended,"
		query3 = query3 + "BoardRollNumber,ReasonNotTakingAdmissionInSameSchool,Board,ResultInPercentage,Location,Preference1,Preference2,"
		query3 = query3 + "Preference3,CreatedBy,CreatedOn) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,Current_Timestamp());"

		//Secondary Exam Log. The process is same
		query3log := "insert into reg_secondary_exam_details_log(ApplicationNumber,PassingYear,SchoolName,SchoolAddress,Schoolstatus,YearAttended,"
		query3log = query3log + "BoardRollNumber,ReasonNotTakingAdmissionInSameSchool,Board,ResultInPercentage,Location,Preference1,Preference2,"
		query3log = query3log + "Preference3,CreatedBy,CreatedOn,Timestamp,OperationType) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,Current_Timestamp(),Current_Timestamp(),1);"

		vals = []interface{}{}

		valslog = []interface{}{}

		vals = append(vals, s.StudentInformation.Application_Number, s.SecondaryExam.Passing_Year, s.SecondaryExam.School_Name, s.SecondaryExam.School_Address)
		vals = append(vals, s.SecondaryExam.School_status, s.SecondaryExam.Year_Attended, s.SecondaryExam.Board_Roll_Number)
		vals = append(vals, s.SecondaryExam.Reason_not_taking_Admission_in_same_School, s.SecondaryExam.Board)
		vals = append(vals, s.SecondaryExam.Result_in_percentage, s.SecondaryExam.Location, s.SecondaryExam.Preference_1)
		vals = append(vals, s.SecondaryExam.Preference_2, s.SecondaryExam.Preference_3, s.SecondaryExam.Created_By)

		valslog = append(valslog, s.StudentInformation.Application_Number, s.SecondaryExam.Passing_Year, s.SecondaryExam.School_Name, s.SecondaryExam.School_Address)
		valslog = append(valslog, s.SecondaryExam.School_status, s.SecondaryExam.Year_Attended, s.SecondaryExam.Board_Roll_Number)
		valslog = append(valslog, s.SecondaryExam.Reason_not_taking_Admission_in_same_School, s.SecondaryExam.Board)
		valslog = append(valslog, s.SecondaryExam.Result_in_percentage, s.SecondaryExam.Location, s.SecondaryExam.Preference_1)
		valslog = append(valslog, s.SecondaryExam.Preference_2, s.SecondaryExam.Preference_3, s.SecondaryExam.Created_By)

		fmt.Println(query3)

		fmt.Println(query3log)

		res, err = txn.Exec(query3, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		// for log

		res, err = txn.Exec(query3log, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		//Address Info. The process is same.
		vals = []interface{}{}

		valslog = []interface{}{}

		addressQuery := "insert into reg_address_information(ApplicationNumber,Address,Pincode,City,State,Country,PhoneNumber,AddressType,"
		addressQuery += "CreatedBy,CreatedOn) values"
		for _, ai := range s.AddressInformation {
			addressQuery += "(?,?,?,?,?,?,?,?,?,current_timestamp()),"
			vals = append(vals, s.StudentInformation.Application_Number, ai.Address, ai.Pincode, ai.City, ai.State, ai.Country, ai.Phone_Number, ai.Address_Type, ai.Created_By)
		}
		// for log
		addressQuerylog := "insert into reg_address_information_log(ApplicationNumber,Address,Pincode,City,State,Country,PhoneNumber,AddressType,"
		addressQuerylog += "CreatedBy,CreatedOn,Timestamp,OperationType) values"
		for _, ai := range s.AddressInformation {
			addressQuerylog += "(?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp,1),"
			valslog = append(valslog, s.StudentInformation.Application_Number, ai.Address, ai.Pincode, ai.City, ai.State, ai.Country, ai.Phone_Number, ai.Address_Type, ai.Created_By)
		}
		addressQuery = addressQuery[0 : len(addressQuery)-1]
		addressQuery += ";"
		res, err = txn.Exec(addressQuery, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}

		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		addressQuerylog = addressQuerylog[0 : len(addressQuerylog)-1]
		addressQuerylog += ";"
		res, err = txn.Exec(addressQuerylog, valslog...)

		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		//Parents info
		vals = []interface{}{}

		valslog = []interface{}{}

		parentsquery := "insert into reg_parent_information(ApplicationNumber,Name,Qualification,College,CourseName,CollegePassingYear,WorkingStatus,"
		parentsquery += "EmploymentType,Designation,Occupation,MobileNumber,EmailID,Organization,AnnualIncome,OrganizationAddress,"
		parentsquery += "SchoolAlumini,SchoolPassingYear,CreatedBy,ProfilePicture,ParentType,Salutation,CreatedOn) values"

		for _, pi := range s.ParentInformation {
			parentsquery += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp()),"
			vals = append(vals, s.StudentInformation.Application_Number, pi.Name, pi.Qualification, pi.College, pi.Course_Name, pi.College_Passing_Year, pi.Working_Status, pi.Employment_Type)
			vals = append(vals, pi.Designation, pi.Occupation, pi.Mobile_Number, pi.Email_ID, pi.Organization, pi.Annual_Income, pi.Organization_Address)
			vals = append(vals, pi.School_Alumini, pi.School_Passing_year, pi.Created_By, pi.Profile_Picture, pi.Parent_Type, pi.Salutation)
		}

		// for log

		parentsquerylog := "insert into reg_parent_information_log(ApplicationNumber,Name,Qualification,College,CourseName,CollegePassingYear,WorkingStatus,"
		parentsquerylog += "EmploymentType,Designation,Occupation,MobileNumber,EmailID,Organization,AnnualIncome,OrganizationAddress,"
		parentsquerylog += "SchoolAlumini,SchoolPassingYear,CreatedBy,ProfilePicture,ParentType,Salutation,CreatedOn,TimeStamp,OperationType) values"

		for _, pi := range s.ParentInformation {
			parentsquerylog += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp(),1),"
			valslog = append(valslog, s.StudentInformation.Application_Number, pi.Name, pi.Qualification, pi.College, pi.Course_Name, pi.College_Passing_Year, pi.Working_Status, pi.Employment_Type)
			valslog = append(valslog, pi.Designation, pi.Occupation, pi.Mobile_Number, pi.Email_ID, pi.Organization, pi.Annual_Income, pi.Organization_Address)
			valslog = append(valslog, pi.School_Alumini, pi.School_Passing_year, pi.Created_By, pi.Profile_Picture, pi.Parent_Type, pi.Salutation)
		}
		parentsquery = parentsquery[0 : len(parentsquery)-1]
		parentsquery += ";"
		fmt.Println(parentsquery)

		res, err = txn.Exec(parentsquery, vals...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		// for log
		parentsquerylog = parentsquerylog[0 : len(parentsquerylog)-1]
		parentsquerylog += ";"
		fmt.Println(parentsquerylog)
		res, err = txn.Exec(parentsquerylog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)

		txn.Commit()
		tempsr.Value = 1
		srs = append(srs, tempsr)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	//Set header type and initialize connection string
	fmt.Println("List called")
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:shubham@tcp(127.0.0.1:3306)/mrerp_student_registration")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("List api called")

	defer r.Body.Close()

	var input *Student_Details = &Student_Details{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	//Decode input json and insert it into s
	err = json.Unmarshal(body, &input)

	if err != nil {
		panic(err.Error())
	}

	//The filters will be stored in vals now
	filters := ""
	filterApplied := false

	//Query to insert student information data
	query := "select ApplicationNumber,ApplicationStatus,FirstName,MiddleName,LastName,Class,DOB,AppliedBatch,StaffWard,"
	query += "ResidentialCategory,AdmissionType,Medical,MedicalPrecaution,Nationality,Gender,BloddGroup,EmergencyNumber,MotherTongue,"
	query += "AadharNumber,Religion,SocialCategory,Locality,BirthPlace,Transport,PreviousSchool,Allergy,AllergyPrecaution,StudentStatus,"
	query += "PaymentStatus,ProfilePicture,AppliedBy,InstituteId,CreatedBy,CreatedOn,ModifiedBy,ModifiedOn from reg_student_information where "

	vals := []interface{}{}
	if input.StudentInformation.Application_Number != nil {
		filters += "ApplicationNumber = ? and "
		vals = append(vals, input.StudentInformation.Application_Number)
		filterApplied = true
	}

	if input.StudentInformation.Application_Status != nil {
		filters += "ApplicationStatus = ? and "
		vals = append(vals, input.StudentInformation.Application_Status)
		filterApplied = true
	}

	if input.StudentInformation.First_Name != nil {
		filters += "FirstName like ('%'+?+'%') and "
		vals = append(vals, input.StudentInformation.First_Name)
		filterApplied = true
	}

	if input.StudentInformation.Middle_Name != nil {
		filters += "MiddleName like ('%'+?+'%') and "
		vals = append(vals, input.StudentInformation.Middle_Name)
		filterApplied = true
	}

	if input.StudentInformation.Last_Name != nil {
		filters += "LastName like ('%'+?+'%') and "
		vals = append(vals, input.StudentInformation.Last_Name)
		filterApplied = true
	}

	if input.StudentInformation.Class != nil {
		filters += "Class = ? and "
		vals = append(vals, input.StudentInformation.Class)
		filterApplied = true
	}
	if input.StudentInformation.Batch != nil {
		filters += "Batch = ? and "
		vals = append(vals, input.StudentInformation.Batch)
		filterApplied = true
	}
	if input.StudentInformation.Staff_Ward != nil {
		filters += "StaffWard = ? and "
		vals = append(vals, input.StudentInformation.Staff_Ward)
		filterApplied = true
	}
	if input.StudentInformation.Admission_Type != nil {
		filters += "AdmissionType = ? and "
		vals = append(vals, input.StudentInformation.Admission_Type)
		filterApplied = true
	}
	if input.StudentInformation.Gender != nil {
		filters += "Gender = ? and "
		vals = append(vals, input.StudentInformation.Gender)
		filterApplied = true
	}
	if input.StudentInformation.Social_Category != nil {
		filters += "SocialCategory = ? and "
		vals = append(vals, input.StudentInformation.Social_Category)
		filterApplied = true
	}
	if input.StudentInformation.Transport_Facilities != nil {
		filters += "Transport = ? and "
		vals = append(vals, input.StudentInformation.Transport_Facilities)
		filterApplied = true
	}
	if input.StudentInformation.Institute_Id != nil {
		filters += "InstituteId = ? and "
		vals = append(vals, input.StudentInformation.Institute_Id)
		filterApplied = true
	}

	if filterApplied {
		filters = filters[0 : len(filters)-5]
		filters += ";"
		query += filters
	} else {
		query = query[0 : len(query)-7]
		query += ";"
	}

	results, err := db.Query(query, vals...)
	if err != nil {
		panic(err.Error())
	}

	var students []Student_Details
	for results.Next() {
		var student Student_Details
		err := results.Scan(&student.StudentInformation.Application_Number, &student.StudentInformation.Application_Status, &student.StudentInformation.First_Name, &student.StudentInformation.Middle_Name, &student.StudentInformation.Last_Name, &student.StudentInformation.Class, &student.StudentInformation.DOB, &student.StudentInformation.Batch, &student.StudentInformation.Staff_Ward, &student.StudentInformation.Residential_Category, &student.StudentInformation.Admission_Type, &student.StudentInformation.Medical, &student.StudentInformation.Medical_Precaution, &student.StudentInformation.Nationality, &student.StudentInformation.Gender, &student.StudentInformation.Blood_Group, &student.StudentInformation.Emergency_Number, &student.StudentInformation.Mother_Tongue, &student.StudentInformation.Aadhar_Number, &student.StudentInformation.Religion, &student.StudentInformation.Social_Category, &student.StudentInformation.Locality, &student.StudentInformation.Birth_Place, &student.StudentInformation.Transport_Facilities, &student.StudentInformation.Previous_School, &student.StudentInformation.Student_Allergy, &student.StudentInformation.Allergy_Precaution, &student.StudentInformation.Student_Status, &student.StudentInformation.Payment_Status, &student.StudentInformation.Profile_Picture, &student.StudentInformation.Applied_By, &student.StudentInformation.Institute_Id, &student.StudentInformation.Created_By, &student.StudentInformation.Created_On, &student.StudentInformation.Modified_By, &student.StudentInformation.Modified_On)
		if err != nil {
			panic(err.Error())
		}

		//Secondary Exam Filters
		sequery := "select ApplicationNumber,PassingYear,SchoolName,SchoolAddress,Schoolstatus,YearAttended,BoardRollNumber,"
		sequery += "ReasonNotTakingAdmissionInSameSchool,ResultInPercentage,Location,Preference1,CreatedBy,CreatedOn,"
		sequery += "ModifiedBy,ModifiedOn,Preference2,Preference3,Board from reg_secondary_exam_details where ApplicationNumber = ?;"

		seresults, err := db.Query(sequery, *student.StudentInformation.Application_Number)

		if err != nil {
			panic(err.Error())
		}

		for seresults.Next() {
			err := seresults.Scan(&student.SecondaryExam.Application_Number, &student.SecondaryExam.Passing_Year, &student.SecondaryExam.School_Name, &student.SecondaryExam.School_Address, &student.SecondaryExam.School_status, &student.SecondaryExam.Year_Attended, &student.SecondaryExam.Board_Roll_Number, &student.SecondaryExam.Reason_not_taking_Admission_in_same_School, &student.SecondaryExam.Result_in_percentage, &student.SecondaryExam.Location, &student.SecondaryExam.Preference_1, &student.SecondaryExam.Created_By, &student.SecondaryExam.Created_On, &student.SecondaryExam.Modified_By, &student.SecondaryExam.Modified_On, &student.SecondaryExam.Preference_2, &student.SecondaryExam.Preference_3, &student.SecondaryExam.Board)

			if err != nil {
				panic(err.Error())
			}
		}

		//Guardian Details
		gdquery := "select ApplicationNumber,LocalGuardianName, LocalGuardiaAddress, GuardiaMobileNumber,EmergencyContactPerson,"
		gdquery += "GuardianPhotograph,CreatedBy,	CreatedOn, ModifiedBy,	ModifiedOn from reg_guardian_details where ApplicationNumber = ?;"

		gdresults, err := db.Query(gdquery, *student.StudentInformation.Application_Number)

		if err != nil {
			panic(err.Error())
		}

		for gdresults.Next() {
			err := gdresults.Scan(&student.GuardianDetails.Application_Number, &student.GuardianDetails.Guardian_Name, &student.GuardianDetails.Guardian_Address, &student.GuardianDetails.Mobile_Number, &student.GuardianDetails.Emergency_Contact_Person, &student.GuardianDetails.Photograph, &student.GuardianDetails.Created_By, &student.GuardianDetails.Created_On, &student.GuardianDetails.Modified_By, &student.GuardianDetails.Modified_On)

			if err != nil {
				panic(err.Error())
			}

			//Guardian Languages. One to many
			glquery := "select ApplicationNumber,CreatedBy,CreatedOn,Language from reg_Guardian_Languages where ApplicationNumber=?;"

			glresults, err := db.Query(glquery, *student.StudentInformation.Application_Number)

			if err != nil {
				panic(err.Error())
			}

			for glresults.Next() {
				var lang Guardian_Languages
				err := glresults.Scan(&lang.Application_Number, &lang.CreatedBy, &lang.CreatedOn, &lang.Language)
				if err != nil {

				}
				student.GuardianDetails.Languages = append(student.GuardianDetails.Languages, lang)
			}

			//Guardian Activities. One to many
			gaquery := "select ApplicationNumber,CreatedBy,CreatedOn,Activity from reg_Guardian_Activities where ApplicationNumber=?;"

			garesults, err := db.Query(gaquery, *student.StudentInformation.Application_Number)

			if err != nil {
				panic(err.Error())
			}

			for garesults.Next() {
				var activity Guardian_Activities
				err := garesults.Scan(&activity.Application_Number, &activity.CreatedBy, &activity.CreatedOn, &activity.Activity)
				if err != nil {
					panic(err.Error())
				}
				student.GuardianDetails.Activities = append(student.GuardianDetails.Activities, activity)
			}
		}

		// Parent Filter
		paquery := "select ApplicationNumber, Name, Qualification, College, CourseName, CollegePassingYear, WorkingStatus, EmploymentType,"
		paquery += "Designation, Occupation, MobileNumber, EmailId, Organization, AnnualIncome, OrganizationAddress,SchoolAlumini,"
		paquery += " CreatedBy,CreatedOn,ModifiedBy,ModifiedOn,ProfilePicture,ParentType,Salutation,SchoolPassingYear "
		paquery += " from reg_parent_Information where ApplicationNumber=?;"

		paresults, err := db.Query(paquery, *student.StudentInformation.Application_Number)

		if err != nil {
			panic(err.Error())
		}

		for paresults.Next() {
			var parent Parent_Information
			err := paresults.Scan(&parent.Application_Number, &parent.Name, &parent.Qualification, &parent.College, &parent.Course_Name, &parent.College_Passing_Year, &parent.Working_Status, &parent.Employment_Type, &parent.Designation, &parent.Occupation, &parent.Mobile_Number, &parent.Email_ID, &parent.Organization, &parent.Annual_Income, &parent.Organization_Address, &parent.School_Alumini, &parent.Created_By, &parent.Created_On, &parent.Modified_By, &parent.Modified_On, &parent.Profile_Picture, &parent.Parent_Type, &parent.Salutation, &parent.School_Passing_year)
			if err != nil {
				panic(err.Error())
			}
			student.ParentInformation = append(student.ParentInformation, parent)
		}

		//Address Filters
		adquery := "select ApplicationNumber, Address,Pincode,City,State,Country,PhoneNumber,AddressType,CreatedBy,CreatedOn,ModifiedBy,ModifiedOn"
		adquery += " from reg_address_Information where ApplicationNumber=?;"

		adresults, err := db.Query(adquery, *student.StudentInformation.Application_Number)

		if err != nil {
			panic(err.Error())
		}

		for adresults.Next() {
			var adrress Address_Information
			err := adresults.Scan(&adrress.Application_Number, &adrress.Address, &adrress.Pincode, &adrress.City, &adrress.State, &adrress.Country, &adrress.Phone_Number, &adrress.Address_Type, &adrress.Created_By, &adrress.Created_On, &adrress.Modified_By, &adrress.Modified_On)
			if err != nil {
				panic(err.Error())
			}
			student.AddressInformation = append(student.AddressInformation, adrress)
		}

		students = append(students, student)

	}

	json.NewEncoder(w).Encode(students)
}

func edit(w http.ResponseWriter, r *http.Request) {

	var srs []StandardResponse
	var sr StandardResponse
	defer func() {
		srs = append([]StandardResponse{sr}, srs...)
		json.NewEncoder(w).Encode(srs)
	}()
	sr.Value = 0

	//Set header type and initialize connection string
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:shubham@tcp(127.0.0.1:3306)/mrerp_student_registration")
	if err != nil {
		sr.Value = -1
		return
	}

	//This will always be called after the function has finished executing. In case the transaction is not commited, it will rollback otherwise it will have no effect.

	defer r.Body.Close()

	var inputArray []Student_Details

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sr.Data = "Error Reading Input"
		sr.Value = -201
		return
	}

	//Decode input json and insert it into s
	err = json.Unmarshal(body, &inputArray)

	//fmt.Println("Edit api called")

	if err != nil {
		sr.Data = "Error Parsing Input"
		sr.Value = -202
		return
	}

	for _, input := range inputArray {
		var tempsr StandardResponse
		tempsr.Value = 0
		tempsr.Data = strconv.Itoa(*input.StudentInformation.Application_Number)
		//Create and initialize transaction
		txn, err := db.Begin()
		if err != nil {
			//sr.Data = err.Error()
			tempsr.Data = "The database could not be connected."
			tempsr.Value = -101
			srs = append(srs, tempsr)

			continue

		}

		//Query to edit student information data

		defer txn.Rollback()
		query := "update reg_Student_Information set "

		vals := []interface{}{}

		if input.StudentInformation.First_Name != nil {
			query += "FirstName = ?, "
			vals = append(vals, input.StudentInformation.First_Name)
		}

		if input.StudentInformation.Middle_Name != nil {
			query += "MiddleName = ?, "
			vals = append(vals, input.StudentInformation.Middle_Name)
		}

		if input.StudentInformation.Last_Name != nil {
			query += "LastName = ?, "
			vals = append(vals, input.StudentInformation.Last_Name)
		}

		if input.StudentInformation.Class != nil {
			query += "Class = ?, "
			vals = append(vals, input.StudentInformation.Class)
		}

		if input.StudentInformation.DOB != nil {
			query += "DOB = ?, "
			vals = append(vals, input.StudentInformation.DOB)
		}

		if input.StudentInformation.Batch != nil {
			query += "AppliedBatch = ?, "
			vals = append(vals, input.StudentInformation.Batch)
		}

		if input.StudentInformation.Staff_Ward != nil {
			query += "StaffWard = ?, "
			vals = append(vals, input.StudentInformation.Staff_Ward)
		}

		if input.StudentInformation.Residential_Category != nil {
			query += "ResidentialCategory = ?, "
			vals = append(vals, input.StudentInformation.Residential_Category)
		}

		if input.StudentInformation.Admission_Type != nil {
			query += "AdmissionType = ?, "
			vals = append(vals, input.StudentInformation.Admission_Type)
		}

		if input.StudentInformation.Medical != nil {
			query += "Medical = ?, "
			vals = append(vals, input.StudentInformation.Medical)
		}

		if input.StudentInformation.Medical_Precaution != nil {
			query += "MedicalPrecaution = ?, "
			vals = append(vals, input.StudentInformation.Medical_Precaution)
		}

		if input.StudentInformation.Nationality != nil {
			query += "Nationality = ?, "
			vals = append(vals, input.StudentInformation.Nationality)
		}

		if input.StudentInformation.Gender != nil {
			query += "Gender = ?, "
			vals = append(vals, input.StudentInformation.Gender)
		}

		if input.StudentInformation.Blood_Group != nil {
			query += "BloddGroup = ?, "
			vals = append(vals, input.StudentInformation.Blood_Group)
		}

		if input.StudentInformation.Emergency_Number != nil {
			query += "EmergencyNumber = ?, "
			vals = append(vals, input.StudentInformation.Emergency_Number)
		}

		if input.StudentInformation.Mother_Tongue != nil {
			query += "MotherTongue = ?, "
			vals = append(vals, input.StudentInformation.Mother_Tongue)
		}

		if input.StudentInformation.Aadhar_Number != nil {
			query += "AadharNumber = ?, "
			vals = append(vals, input.StudentInformation.Aadhar_Number)
		}

		if input.StudentInformation.Religion != nil {
			query += "Religion = ?, "
			vals = append(vals, input.StudentInformation.Religion)
		}

		if input.StudentInformation.Social_Category != nil {
			query += "SocialCategory = ?, "
			vals = append(vals, input.StudentInformation.Social_Category)
		}

		if input.StudentInformation.Locality != nil {
			query += "Locality = ?, "
			vals = append(vals, input.StudentInformation.Locality)
		}

		if input.StudentInformation.Birth_Place != nil {
			query += "BirthPlace = ?, "
			vals = append(vals, input.StudentInformation.Birth_Place)
		}

		if input.StudentInformation.Transport_Facilities != nil {
			query += "Transport = ?, "
			vals = append(vals, input.StudentInformation.Transport_Facilities)
		}

		if input.StudentInformation.Previous_School != nil {
			query += "PreviousSchool = ?, "
			vals = append(vals, input.StudentInformation.Previous_School)
		}

		if input.StudentInformation.Student_Allergy != nil {
			query += "Allergy = ?, "
			vals = append(vals, input.StudentInformation.Student_Allergy)
		}

		if input.StudentInformation.Allergy_Precaution != nil {
			query += "AllergyPrecaution = ?, "
			vals = append(vals, input.StudentInformation.Allergy_Precaution)
		}

		if input.StudentInformation.Student_Status != nil {
			query += "StudentStatus = ?, "
			vals = append(vals, input.StudentInformation.Student_Status)
		}

		if input.StudentInformation.Payment_Status != nil {
			query += "PaymentStatus = ?, "
			vals = append(vals, input.StudentInformation.Payment_Status)
		}

		if input.StudentInformation.Profile_Picture != nil {
			query += "ProfilePicture = ?, "
			vals = append(vals, input.StudentInformation.Profile_Picture)
		}

		if input.StudentInformation.Modified_By != nil {
			query += "ModifiedBy = ?, "
			vals = append(vals, input.StudentInformation.Modified_By)
		}

		if input.StudentInformation.Modified_On != nil {
			query += "ModifiedOn = ?, "
			vals = append(vals, input.StudentInformation.Modified_On)
		}

		if input.StudentInformation.Institute_Id != nil {
			query += "InstituteId = ?, "
			vals = append(vals, input.StudentInformation.Institute_Id)
		}

		query = query[0 : len(query)-2]
		query += " where ApplicationNumber = ?;"
		vals = append(vals, input.StudentInformation.Application_Number)
		_, err = txn.Exec(query, vals...)

		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}

		// log

		//Query to insert student information data into log
		queryLog := "insert into reg_student_information_log(ApplicationNumber,ApplicationStatus,FirstName,MiddleName,LastName,Class,DOB,AppliedBatch,"
		queryLog = queryLog + "StaffWard,ResidentialCategory,AdmissionType,Medical,MedicalPrecaution,Nationality,Gender,BloddGroup,EmergencyNumber,"
		queryLog = queryLog + "MotherTongue,AadharNumber,Religion,SocialCategory,Locality,BirthPlace,Transport,PreviousSchool,Allergy,"
		queryLog = queryLog + "AllergyPrecaution,StudentStatus,PaymentStatus,ProfilePicture,AppliedBy,InstituteId,CreatedBy,CreatedOn,Timestamp,OperationType) values"
		queryLog = queryLog + "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp(),2);"

		valslog := []interface{}{}

		//Start appending parameters to valslog

		valslog = append(valslog, input.StudentInformation.Application_Number, input.StudentInformation.Application_Status, input.StudentInformation.First_Name, input.StudentInformation.Middle_Name)
		valslog = append(valslog, input.StudentInformation.Last_Name, input.StudentInformation.Class, input.StudentInformation.DOB, input.StudentInformation.Batch)
		valslog = append(valslog, input.StudentInformation.Staff_Ward, input.StudentInformation.Residential_Category, input.StudentInformation.Admission_Type)
		valslog = append(valslog, input.StudentInformation.Medical, input.StudentInformation.Medical_Precaution, input.StudentInformation.Nationality)
		valslog = append(valslog, input.StudentInformation.Gender, input.StudentInformation.Blood_Group, input.StudentInformation.Emergency_Number)
		valslog = append(valslog, input.StudentInformation.Mother_Tongue, input.StudentInformation.Aadhar_Number, input.StudentInformation.Religion)
		valslog = append(valslog, input.StudentInformation.Social_Category, input.StudentInformation.Locality, input.StudentInformation.Birth_Place)
		valslog = append(valslog, input.StudentInformation.Transport_Facilities, input.StudentInformation.Previous_School)
		valslog = append(valslog, input.StudentInformation.Student_Allergy, input.StudentInformation.Allergy_Precaution)
		valslog = append(valslog, input.StudentInformation.Student_Status, input.StudentInformation.Payment_Status, input.StudentInformation.Profile_Picture)
		valslog = append(valslog, input.StudentInformation.Applied_By, input.StudentInformation.Institute_Id, input.StudentInformation.Created_By)

		// Execution of valslog

		_, err = txn.Exec(queryLog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		/*ra = res.RowsAffected
		fmt.Println("The result of first query is ", &ra)

		fmt.Println("Rows affected: ", &ra)*/

		//TODO: update secondary exam

		if input.SecondaryExam.Application_Number != nil {
			if *input.SecondaryExam.Application_Number != *input.StudentInformation.Application_Number {
				panic("The student info and secondary exam details data does not match.")
			}
			sequery := "update reg_Secondary_Exam_details set "
			vals = []interface{}{}
			if input.SecondaryExam.Passing_Year != nil {
				sequery += "PassingYear = ?, "
				vals = append(vals, input.SecondaryExam.Passing_Year)
			}
			if input.SecondaryExam.School_Name != nil {
				sequery += "SchoolName = ?, "
				vals = append(vals, input.SecondaryExam.School_Name)
			}
			if input.SecondaryExam.School_Address != nil {
				sequery += "SchoolAddress = ?, "
				vals = append(vals, input.SecondaryExam.School_Address)
			}

			if input.SecondaryExam.School_status != nil {
				sequery += "SchoolStatus = ?, "
				vals = append(vals, input.SecondaryExam.School_status)
			}

			if input.SecondaryExam.Year_Attended != nil {
				sequery += "YearAttened = ?, "
				vals = append(vals, input.SecondaryExam.Year_Attended)
			}
			if input.SecondaryExam.Board_Roll_Number != nil {
				sequery += "BoardRollNumber = ?, "
				vals = append(vals, input.SecondaryExam.Board_Roll_Number)
			}
			if input.SecondaryExam.Reason_not_taking_Admission_in_same_School != nil {
				sequery += "ReasonNotTakingAdmissionInSameSchool = ?, "
				vals = append(vals, input.SecondaryExam.Reason_not_taking_Admission_in_same_School)
			}
			if input.SecondaryExam.Board != nil {
				sequery += "Board = ?, "
				vals = append(vals, input.SecondaryExam.Board)
			}
			if input.SecondaryExam.Result_in_percentage != nil {
				sequery += "ResultInPercentage = ?, "
				vals = append(vals, input.SecondaryExam.Result_in_percentage)
			}
			if input.SecondaryExam.Location != nil {
				sequery += "Location = ?, "
				vals = append(vals, input.SecondaryExam.Location)
			}
			if input.SecondaryExam.Preference_1 != nil {
				sequery += "Preference1 = ?, "
				vals = append(vals, input.SecondaryExam.Preference_1)
			}
			if input.SecondaryExam.Preference_2 != nil {
				sequery += "Preference2 = ?, "
				vals = append(vals, input.SecondaryExam.Preference_2)
			}
			if input.SecondaryExam.Preference_3 != nil {
				sequery += "Preference3 = ?, "
				vals = append(vals, input.SecondaryExam.Preference_3)
			}
			if input.SecondaryExam.Modified_By != nil {
				sequery += "ModifiedBy = ?, "
				vals = append(vals, input.SecondaryExam.Modified_By)
			}
			if input.SecondaryExam.Modified_On != nil {
				sequery += "ModifiedOn = ?, "
				vals = append(vals, input.SecondaryExam.Modified_On)
			}
			sequery = sequery[0 : len(sequery)-2]
			sequery += " where ApplicationNumber = ?;"

			fmt.Println(sequery)

			vals = append(vals, input.SecondaryExam.Application_Number)
			_, err := txn.Exec(sequery, vals...)

			if err != nil {
				tempsr.Data = "SQL Query Error"
				tempsr.Value = -301
				srs = append(srs, tempsr)
				txn.Rollback()
				continue
			}
		}
		// For Log

		//Secondary Exam Log. The process is same
		query3log := "insert into reg_secondary_exam_details_log(ApplicationNumber,PassingYear,SchoolName,SchoolAddress,Schoolstatus,YearAttended,"
		query3log = query3log + "BoardRollNumber,ReasonNotTakingAdmissionInSameSchool,Board,ResultInPercentage,Location,Preference1,Preference2,"
		query3log = query3log + "Preference3,CreatedBy,CreatedOn,Timestamp,OperationType) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,Current_Timestamp(),Current_Timestamp(),2);"

		valslog = []interface{}{}

		valslog = append(valslog, input.StudentInformation.Application_Number, input.SecondaryExam.Passing_Year, input.SecondaryExam.School_Name, input.SecondaryExam.School_Address)
		valslog = append(valslog, input.SecondaryExam.School_status, input.SecondaryExam.Year_Attended, input.SecondaryExam.Board_Roll_Number)
		valslog = append(valslog, input.SecondaryExam.Reason_not_taking_Admission_in_same_School, input.SecondaryExam.Board)
		valslog = append(valslog, input.SecondaryExam.Result_in_percentage, input.SecondaryExam.Location, input.SecondaryExam.Preference_1)
		valslog = append(valslog, input.SecondaryExam.Preference_2, input.SecondaryExam.Preference_3, input.SecondaryExam.Created_By)

		_, err = txn.Exec(query3log, valslog...)

		fmt.Println(query3log)

		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		/*ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)*/

		//Log Guardian details. The process is same as above

		logquery2 := "insert into reg_guardian_details_log(ApplicationNumber,LocalGuardianName,LocalGuardiaAddress,GuardiaMobileNumber,"
		logquery2 = logquery2 + "EmergencyContactPerson,GuardianPhotograph,CreatedBy,CreatedOn,Timestamp,OperationType) values(?,?,?,?,?,?,?,current_timestamp(),"
		logquery2 = logquery2 + "current_timestamp(),2);"

		fmt.Println(logquery2)

		valslog = []interface{}{}

		valslog = append(valslog, input.StudentInformation.Application_Number, input.GuardianDetails.Guardian_Name, input.GuardianDetails.Guardian_Address, input.GuardianDetails.Mobile_Number)
		valslog = append(valslog, input.GuardianDetails.Emergency_Contact_Person, input.GuardianDetails.Photograph, input.GuardianDetails.Created_By)

		//fmt.Println(logquery2)

		_, err = txn.Exec(logquery2, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		/*ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)*/

		//Guardian Languages.

		valslog = []interface{}{}

		//log

		glquerylog := "insert into reg_guardian_languages_log(ApplicationNumber,Language,CreatedBy,CreatedOn,Timestamp,OperationType) values"
		//(last_insert_id(),?,?,current_timestamp());
		for _, gl := range input.GuardianDetails.Languages {
			glquerylog += "(?,?,?,current_timestamp(),current_timestamp(),2),"
			valslog = append(valslog, input.StudentInformation.Application_Number, gl.Language, gl.CreatedBy)
		}
		glquerylog = glquerylog[0 : len(glquerylog)-1]
		glquerylog += ";"
		fmt.Println(glquerylog)

		// for log

		_, err = txn.Exec(glquerylog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		/*ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)*/

		//Guardian activities.

		valslog = []interface{}{}

		// for log

		gaquerylog := "insert into reg_guardian_activities_log(ApplicationNumber,Activity,CreatedBy,CreatedOn,Timestamp,OperationType) values"
		//(last_insert_id(),?,?,current_timestamp());
		for _, ga := range input.GuardianDetails.Activities {
			gaquerylog += "(?,?,?,current_timestamp(),current_timestamp(),2),"
			valslog = append(valslog, input.StudentInformation.Application_Number, ga.Activity, ga.CreatedBy)
		}
		gaquerylog = gaquerylog[0 : len(gaquerylog)-1]
		gaquerylog += ";"
		fmt.Println(gaquerylog)

		// for log

		_, err = txn.Exec(gaquerylog, valslog...)
		if err != nil {
			tempsr.Data = "SQL Query Error"
			tempsr.Value = -301
			srs = append(srs, tempsr)
			txn.Rollback()
			continue
		}
		/*ra = res.RowsAffected
		fmt.Println("Rows affected: ", &ra)*/

		//Update parent details

		for _, parent := range input.ParentInformation {

			if parent.Application_Number != nil {
				if *parent.Application_Number != *input.StudentInformation.Application_Number {
					panic("The student info and parent details data does not match.")
				}
				paquery := "update reg_parent_information set "

				vals = []interface{}{}

				if parent.Name != nil {
					paquery += "Name = ?, "
					vals = append(vals, parent.Name)
				}

				if parent.Qualification != nil {
					paquery += "Qualification = ?, "
					vals = append(vals, parent.Qualification)
				}

				if parent.College != nil {
					paquery += "College = ?, "
					vals = append(vals, parent.College)
				}

				if parent.Course_Name != nil {
					paquery += "CourseName = ?, "
					vals = append(vals, parent.Course_Name)
				}

				if parent.College_Passing_Year != nil {
					paquery += "CollegePassingYear = ?, "
					vals = append(vals, parent.College_Passing_Year)
				}

				if parent.Working_Status != nil {
					paquery += "WorkingStatus = ?, "
					vals = append(vals, parent.Working_Status)
				}

				if parent.Employment_Type != nil {
					paquery += "EmploymentType = ?, "
					vals = append(vals, parent.Employment_Type)
				}

				if parent.Designation != nil {
					paquery += "Designation = ?, "
					vals = append(vals, parent.Designation)
				}

				if parent.Occupation != nil {
					paquery += "Occupation = ?, "
					vals = append(vals, parent.Occupation)
				}

				if parent.Mobile_Number != nil {
					paquery += "MobileNumber = ?, "
					vals = append(vals, parent.Mobile_Number)
				}

				if parent.Email_ID != nil {
					paquery += "EmailID = ?, "
					vals = append(vals, parent.Email_ID)
				}

				if parent.Organization != nil {
					paquery += "Organization = ?, "
					vals = append(vals, parent.Organization)
				}

				if parent.Annual_Income != nil {
					paquery += "AnnualIncome = ?, "
					vals = append(vals, parent.Annual_Income)
				}

				if parent.Organization_Address != nil {
					paquery += "OrganizationAddress = ?, "
					vals = append(vals, parent.Organization_Address)
				}

				if parent.School_Alumini != nil {
					paquery += "SchoolAlumni = ?, "
					vals = append(vals, parent.School_Alumini)
				}

				if parent.Modified_By != nil {
					paquery += "ModifiedBy = ?, "
					vals = append(vals, parent.Modified_By)
				}

				if parent.Modified_On != nil {
					paquery += "ModifiedOn = current_timestamp(), "
				}

				if parent.Profile_Picture != nil {
					paquery += "ProfilePicture = ?, "
					vals = append(vals, parent.Profile_Picture)
				}

				if parent.Salutation != nil {
					paquery += "Salutation = ?, "
					vals = append(vals, parent.Salutation)
				}

				if parent.School_Passing_year != nil {
					paquery += "SchoolPassingYear = ?, "
					vals = append(vals, parent.School_Passing_year)
				}

				paquery = paquery[0 : len(paquery)-2]
				paquery += " where ApplicationNumber = ? and ParentType = ?;"

				fmt.Println(paquery)

				vals = append(vals, input.StudentInformation.Application_Number, parent.Parent_Type)
				_, err := txn.Exec(paquery, vals...)

				if err != nil {
					tempsr.Data = "SQL Query Error"
					tempsr.Value = -301
					srs = append(srs, tempsr)
					txn.Rollback()
					continue
				}

			}
			// for log

			valslog = []interface{}{}

			parentsquerylog := "insert into reg_parent_information_log(ApplicationNumber,Name,Qualification,College,CourseName,CollegePassingYear,WorkingStatus,"
			parentsquerylog += "EmploymentType,Designation,Occupation,MobileNumber,EmailID,Organization,AnnualIncome,OrganizationAddress,"
			parentsquerylog += "SchoolAlumini,SchoolPassingYear,CreatedBy,ProfilePicture,ParentType,Salutation,CreatedOn,TimeStamp,OperationType) values"
			//values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
			for _, pi := range input.ParentInformation {
				parentsquerylog += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,current_timestamp(),current_timestamp(),2),"
				valslog = append(valslog, input.StudentInformation.Application_Number, pi.Name, pi.Qualification, pi.College, pi.Course_Name, pi.College_Passing_Year, pi.Working_Status, pi.Employment_Type)
				valslog = append(valslog, pi.Designation, pi.Occupation, pi.Mobile_Number, pi.Email_ID, pi.Organization, pi.Annual_Income, pi.Organization_Address)
				valslog = append(valslog, pi.School_Alumini, pi.School_Passing_year, pi.Created_By, pi.Profile_Picture, pi.Parent_Type, pi.Salutation)
			}
			parentsquerylog = parentsquerylog[0 : len(parentsquerylog)-1]
			parentsquerylog += ";"
			fmt.Println(parentsquerylog)
			_, err = txn.Exec(parentsquerylog, valslog...)

			// for log

			if err != nil {
				tempsr.Data = "SQL Query Error"
				tempsr.Value = -301
				srs = append(srs, tempsr)
				txn.Rollback()
				continue
			}
			/*ra = res.RowsAffected
			fmt.Println("Rows affected: ", &ra)*/

		}

		txn.Commit()
		tempsr.Value = 1
		srs = append(srs, tempsr)
	}

}

func handler() {
	router := mux.NewRouter()
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/list", list).Methods("POST")
	router.HandleFunc("/edit", edit).Methods("POST")
	//router.HandleFunc("/logregister", register).Methods("POST")
	//router.HandleFunc("/", isAuthorized(homePage)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8002", router))
}

func main() {
	//GenerateJWT()
	handler()

}

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

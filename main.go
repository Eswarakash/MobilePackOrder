package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId                int
	CustomerName              string
	CustomerNumber            int
	CustomerAlternativeNumber int
	CustomerDob               string
	CustomerGender            string
	CustomerAddress           string
	CustomerPincode           int
	CustomerPlanActive        string
}
type Plan struct {
	PlanId             int
	PlanDuration       int
	PlanVoiceCallLimit int
	PlanDataLimit      int
	PlanSmsLimit       int
	PlanCost           int
}

type CustomerSubscription struct {
	CustomerId    int
	PlanId        int
	PlanStartDate time.Time
	PlanEndDate   time.Time
	AmountPaid    int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "golangapp"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MAIN PAGE")
	tmpl.ExecuteTemplate(w, "MainPage", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Index")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Customer ORDER BY customerId DESC")
	if err != nil {
		panic(err.Error())
	}
	cust := Customer{}
	res := []Customer{}
	for selDB.Next() {
		var id, number, alternativeNumber, pincode int
		var name, dob, gender, address, plan string
		err = selDB.Scan(&id, &name, &number, &alternativeNumber, &dob, &gender, &address, &pincode, &plan)
		if err != nil {
			panic(err.Error())
		}
		cust.CustomerId = id
		cust.CustomerName = name
		cust.CustomerNumber = number
		cust.CustomerAlternativeNumber = alternativeNumber
		cust.CustomerDob = dob
		cust.CustomerGender = gender
		cust.CustomerAddress = address
		cust.CustomerPincode = pincode
		cust.CustomerPlanActive = plan
		res = append(res, cust)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}
func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside show")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Customer ORDER BY customerId DESC")
	if err != nil {
		panic(err.Error())
	}
	cust := Customer{}
	//res := []Customer{}
	for selDB.Next() {
		var id, number, alternativeNumber, pincode int
		var name, dob, gender, address, plan string
		err = selDB.Scan(&id, &name, &number, &alternativeNumber, &dob, &gender, &address, &pincode, &plan)
		if err != nil {
			panic(err.Error())
		}
		cust.CustomerId = id
		cust.CustomerName = name
		cust.CustomerNumber = number
		cust.CustomerAlternativeNumber = alternativeNumber
		cust.CustomerDob = dob
		cust.CustomerGender = gender
		cust.CustomerAddress = address
		cust.CustomerPincode = pincode
		cust.CustomerPlanActive = plan

		//res = append(res, cust)
	}
	tmpl.ExecuteTemplate(w, "Show", cust)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("CustomerName")
		number := r.FormValue("CustomerNumber")
		alternativeNumber := r.FormValue("AlternativeNumber")
		dob := r.FormValue("AlternativeNumber")
		gender := r.FormValue("AlternativeNumber")
		address := r.FormValue("AlternativeNumber")
		pincode := r.FormValue("AlternativeNumber")
		plan := "NOT_ACTIVE"
		insForm, err := db.Prepare("INSERT INTO Customer(customerName, customerNumber, customerAlternativeNumber, customerDob, customerGender, customerAddress, customerPincode, customerPlanActive) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, number, alternativeNumber, dob, gender, address, pincode, plan)
		log.Println("INSERT: Name: " + name + " | Number: " + number + " | Alternate Number: " + alternativeNumber + " | DOB: " + dob + " | Gender: " + gender + " | Address: " + address + " | Pincode: " + pincode + " | Plan Active: " + plan)
	}
	defer db.Close()
	http.Redirect(w, r, "Index", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from Customer where customerId=?", nId)
	if err != nil {
		panic(err.Error())
	}
	cust := Customer{}
	for selDB.Next() {
		var id, number, alternativeNumber, pincode int
		var name, dob, gender, address, plan string
		err = selDB.Scan(&id, &name, &number, &alternativeNumber, &dob, &gender, &address, &pincode, &plan)
		if err != nil {
			panic(err.Error())
		}
		cust.CustomerId = id
		cust.CustomerName = name
		cust.CustomerNumber = number
		cust.CustomerAlternativeNumber = alternativeNumber
		cust.CustomerDob = dob
		cust.CustomerGender = gender
		cust.CustomerAddress = address
		cust.CustomerPincode = pincode
		cust.CustomerPlanActive = plan
		fmt.Println(cust.CustomerId)
		fmt.Println(cust)
	}
	tmpl.ExecuteTemplate(w, "Edit", cust)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE 1")
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("customerName")
		number := r.FormValue("customerNumber")
		alternativeNumber := r.FormValue("customerAlternativeNumber")
		dob := r.FormValue("customerDob")
		gender := r.FormValue("customerGender")
		address := r.FormValue("customerAddress")
		pincode := r.FormValue("customerPincode")
		plan := r.FormValue("customerPlanActive")
		id := r.FormValue("customerId")

		insForm, err := db.Prepare("UPDATE Customer SET customerName=?, customerNumber=?, customerAlternativeNumber=?, customerDob=?, customerGender= ? , customerAddress=?, customerPincode=?, customerPlanActive=?  WHERE customerId=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, number, alternativeNumber, dob, gender, address, pincode, plan, id)
		log.Println("UPDATE: Name: " + name + " | Number: " + number + " | Alternate Number: " + alternativeNumber + " | Dob: " + dob + " | Gender: " + gender + " | Address: " + address + " | Pincode: " + pincode + " | Active Plan : " + plan)
	}

	defer db.Close()
	http.Redirect(w, r, "Index", 301)
}

func Searchtemplate(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Search", nil)
}

func SearchCustomer(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	fmt.Println("search")
	res := []Customer{}
	if r.Method == "POST" {
		id := r.FormValue("customerId")
		name := r.FormValue("customerName")
		alternatenumber := r.FormValue("customerNumber")
		number := r.FormValue("customerAlternativeNumber")
		address := r.FormValue("customerAddress")
		pincode := r.FormValue("customerPincode")

		selDB, err := db.Query("SELECT * FROM Customer where customerId=? or customerName=? or customerNumber=? or customerNumber=? or customerAddress=? or customerPincode=?", id, name, number, alternatenumber, address, pincode)
		if err != nil {
			panic(err.Error())
		}
		customer := Customer{}

		for selDB.Next() {
			err = selDB.Scan(&customer.CustomerId, &customer.CustomerName, &customer.CustomerNumber, &customer.CustomerAlternativeNumber, &customer.CustomerDob, &customer.CustomerGender, &customer.CustomerAddress, &customer.CustomerPincode, &customer.CustomerPlanActive)
			if err != nil {
				panic(err.Error())
			}
			res = append(res, customer)
		}
	}
	tmpl.ExecuteTemplate(w, "Index", res)

	defer db.Close()
}

func Plans(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "PlanRegistration", nil)
}

func PlanRegistration(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		duration := r.FormValue("planDuration")
		voice := r.FormValue("planVoiceCallLimit")
		data := r.FormValue("planDataLimit")
		sms := r.FormValue("planSmsLimit")
		cost := r.FormValue("planCost")
		insForm, err := db.Prepare("INSERT INTO Plan(planDuration, planVoiceCallLimit, PlanDataLimit, planSmsLimit, planCost) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(duration, voice, data, sms, cost)
		log.Println("INSERT:  Duration: " + duration + " | VOice: " + voice + " | Data : " + data + " | Sms: " + sms + " | Cost: " + cost)
	}
	defer db.Close()
	http.Redirect(w, r, "Plan", 301)
}
func ShowPlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside plan show")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Plan ORDER BY planId DESC")
	if err != nil {
		panic(err.Error())
	}
	cust := Plan{}
	res := []Plan{}
	for selDB.Next() {
		var id, duration, voice, data, sms, cost int
		err = selDB.Scan(&id, &duration, &voice, &data, &sms, &cost)
		if err != nil {
			panic(err.Error())
		}
		cust.PlanId = id
		cust.PlanDuration = duration
		cust.PlanVoiceCallLimit = voice
		cust.PlanDataLimit = data
		cust.PlanSmsLimit = sms
		cust.PlanCost = cost
		res = append(res, cust)
	}

	tmpl.ExecuteTemplate(w, "PlanIndex", res)
	fmt.Println("end inside plan show")
	defer db.Close()
}

func EditPlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit Plan")
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("select * from Plan where planId=?", nId)
	if err != nil {
		panic(err.Error())
	}
	cust := Plan{}
	//res := []Plan{}
	for selDB.Next() {
		var id, duration, voice, data, sms, cost int
		err = selDB.Scan(&id, &duration, &voice, &data, &sms, &cost)
		if err != nil {
			panic(err.Error())
		}
		cust.PlanId = id
		cust.PlanDuration = duration
		cust.PlanVoiceCallLimit = voice
		cust.PlanDataLimit = data
		cust.PlanSmsLimit = sms
		cust.PlanCost = cost
		//res = append(res, cust)
	}
	tmpl.ExecuteTemplate(w, "EditPlan", cust)
	fmt.Println("End Edit Plan")

	defer db.Close()
}
func UpdatePlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE 1")
	db := dbConn()
	if r.Method == "POST" {
		duration := r.FormValue("planDuration")
		voice := r.FormValue("planVoiceCallLimit")
		data := r.FormValue("planDataLimit")
		sms := r.FormValue("planSmsLimit")
		cost := r.FormValue("planCost")
		id := r.FormValue("planId")

		insForm, err := db.Prepare("UPDATE Plan SET planDuration=?, planVoiceCallLimit=?, planDataLimit=?, planSmsLimit=?, planCost= ? WHERE planId=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(duration, voice, data, sms, cost, id)
		log.Println("UPDATE: Duration: " + duration + " | Voice: " + voice + " | Data : " + data + " | Sms: " + sms + " | Cost: " + cost)
	}

	defer db.Close()
	http.Redirect(w, r, "PlanIndex", 301)
}

func CustomerRegSubscription(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "RegisterPlan", nil)
}

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		cid := r.FormValue("customerId")
		pid := r.FormValue("planId")
		planStartDate := time.Now()

		var plancost, planduration int

		insForm, err := db.Query("Select planCost,planDuration from Plan  WHERE planId=?", pid)
		if err != nil {
			panic(err.Error())
		}
		for insForm.Next() {
			err := insForm.Scan(&plancost, &planduration)
			if err != nil {
				panic(err.Error())
			}
		}

		planEndDate := planStartDate.AddDate(0, 0, planduration)

		str, err := db.Prepare("INSERT INTO CustomerSubscription(customerId,planId,planStartDate,planEndDate,amountPaid )  VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		str.Exec(cid, pid, planStartDate, planEndDate, plancost)

		str1, err := db.Prepare("UPDATE Customer SET customerPlanActive=? where customerId=? ")
		if err != nil {
			panic(err.Error())
		}
		planactive := "ACTIVE"
		str1.Exec(planactive, cid)

		//log.Println("UPDATE: Duration: " + duration + " | Voice: " + voice + " | Data : " + data + " | Sms: " + sms + " | Cost: " + cost)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func SearchExpiry(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	plan := "NOT_ACTIVE"

	search, err := db.Query("Select * from Customer where customerPlanActive= ?", plan)
	if err != nil {
		panic(err.Error())
	}
	cust := Customer{}
	res := []Customer{}

	for search.Next() {
		var id, number, alternativeNumber, pincode int
		var name, dob, gender, address, plan string
		err = search.Scan(&id, &name, &number, &alternativeNumber, &dob, &gender, &address, &pincode, &plan)
		if err != nil {
			panic(err.Error())
		}
		cust.CustomerId = id
		cust.CustomerName = name
		cust.CustomerNumber = number
		cust.CustomerAlternativeNumber = alternativeNumber
		cust.CustomerDob = dob
		cust.CustomerGender = gender
		cust.CustomerAddress = address
		cust.CustomerPincode = pincode
		cust.CustomerPlanActive = plan
		fmt.Println(cust.CustomerId)
		fmt.Println(cust)
		res = append(res, cust)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
}

func SearchActive(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	plan := "NO"

	search, err := db.Query("Select * from Customer where customerPlanActive= ?", plan)
	if err != nil {
		panic(err.Error())
	}
	cust := Customer{}
	res := []Customer{}

	for search.Next() {
		var id, number, alternativeNumber, pincode int
		var name, dob, gender, address, plan string
		err = search.Scan(&id, &name, &number, &alternativeNumber, &dob, &gender, &address, &pincode, &plan)
		if err != nil {
			panic(err.Error())
		}
		cust.CustomerId = id
		cust.CustomerName = name
		cust.CustomerNumber = number
		cust.CustomerAlternativeNumber = alternativeNumber
		cust.CustomerDob = dob
		cust.CustomerGender = gender
		cust.CustomerAddress = address
		cust.CustomerPincode = pincode
		cust.CustomerPlanActive = plan
		fmt.Println(cust.CustomerId)
		fmt.Println(cust)
		res = append(res, cust)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
}

func SetExpiry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HIIIIIIIi")
	db := dbConn()
	plan := "EXPIRED"
	search, err := db.Prepare("UPDATE Customer SET customerPlanActive=? WHERE customerId>0")
	if err != nil {
		panic(err.Error())
	}
	search.Exec(plan)
	fmt.Println("HIIIIIIIi")

	http.Redirect(w, r, "/viewCustomers", 301)
	//	tmpl.ExecuteTemplate(w, "Index", nil)
	defer db.Close()
}
func SearchByDate(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "SearchByDate", nil)
}

func SearchExpiryCustomer(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	cust := Customer{}
	res := []Customer{}
	if r.Method == "POST" {
		expiryDate := r.FormValue("ExpiryDate")

		insForm, err := db.Query("SELECT a.customerId,a.customerName,a.customerNumber,a.customerAlternativeNumber,a.customerDob,a.customerGender,a.customerAddress,a.customerPincode,a.customerPlanActive FROM Customer a, CustomerSubscription b where a.customerId=b.customerId and b.planEndDate=?", expiryDate)
		if err != nil {
			panic(err.Error())
		}

		var id, number, alternativeNumber, pincode int
		var name, dob, gender, address, plan string
		//var planStartDate,planEndDate,amountPaid int

		for insForm.Next() {

			err = insForm.Scan(&id, &name, &number, &alternativeNumber, &dob, &gender, &address, &pincode, &plan)
			if err != nil {
				panic(err.Error())
			}
			cust.CustomerId = id
			cust.CustomerName = name
			cust.CustomerNumber = number
			cust.CustomerAlternativeNumber = alternativeNumber
			cust.CustomerDob = dob
			cust.CustomerGender = gender
			cust.CustomerAddress = address
			cust.CustomerPincode = pincode
			cust.CustomerPlanActive = plan
			fmt.Println(cust.CustomerId)
			fmt.Println(cust)
			res = append(res, cust)
		}
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

}

func RefreshList(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	plan := "EXPIRED"
	upform, err := db.Prepare("UPDATE Customer a  INNER JOIN CustomerSubscription b ON a.customerId=b.customerId SET a.CustomerPlanActive=? WHERE DATEDIFF(b.planEndDate,CURDATE())<0")

	if err != nil {
		panic(err.Error())
	}

	upform.Exec(plan)
	http.Redirect(w, r, "/viewCustomers", 301)

}

func main() {
	log.Println("Server started on: http://localhost:30000")
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/index", Index)
	http.HandleFunc("/viewCustomers", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/search", Searchtemplate)
	http.HandleFunc("/searchCustomer", SearchCustomer)
	http.HandleFunc("/insertPlan", PlanRegistration)
	http.HandleFunc("/plan", Plans)
	http.HandleFunc("/viewPlans", ShowPlan)
	http.HandleFunc("/editPlan", EditPlan)
	http.HandleFunc("/updatePlan", UpdatePlan)
	http.HandleFunc("/registerPlan", CustomerRegSubscription)
	http.HandleFunc("/updateRegister", RegisterCustomer)
	http.HandleFunc("/searchExpiry", SearchExpiry)
	//http.HandleFunc("/searchACtive", SearchActive)
	http.HandleFunc("/expireSet", SetExpiry)
	http.HandleFunc("/searchByDate", SearchByDate)
	http.HandleFunc("/searchCust", SearchExpiryCustomer)
	http.HandleFunc("/refreshList", RefreshList)
	http.ListenAndServe(":30000", nil)
}

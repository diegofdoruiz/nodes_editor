package main

import (
	"context"
	"errors"
	"flag"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"strconv"

	"github.com/dgraph-io/dgo/v210"
    "github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type CancelFunc func()

func getDgraphClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	ctx := context.Background()

	for {

		err = dg.Login(ctx, "groot", "password")
		if err == nil || !strings.Contains(err.Error(), "Please retry") {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func main() {
	//log.Printf(dbCreateUser("diego", "test"))
	//log.Printf(dbCreateUser("mario", "test"))

	flag.Parse()
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// RESTy routes for "modules" resource
	r.Route("/modules", func(r chi.Router) {
		r.Post("/", ListModules)
		r.Post("/create", CreateModule)
		r.Post("/search", SearchModuleByName)
		r.Route("/{moduleUID}", func(r chi.Router) {
			r.Put("/", ClearModule) // Clear /modules/123
			r.Delete("/", DeleteModule) // DELETE /modules/123
		})
	})

	// RESTy routes for "nodes" resource
	r.Route("/nodes", func(r chi.Router) {
		r.Post("/create", CreateNode)
		r.Post("/data", DataNode)
		r.Route("/{nodeUID}", func(r chi.Router) {
			r.Put("/", PositionNode) // Change position /nodes/123
			r.Delete("/", DeleteNode) // DELETE /nodes/123
		})
		r.Route("/connections", func(r chi.Router) {
			r.Post("/create", CreateConnection)
			r.Post("/delete", DeleteConnections)
		})
	})

	// RESTy routes for "user" resource
	r.Route("/user", func(r chi.Router) {
		r.Post("/login", SignIn)
	})

	http.ListenAndServe(":3333", r)
}

/***************** Models ******************/
type Data struct {
	Uid			string		`json:"uid,omitempty"` 
	Name		string		`json:"name,omitempty"`
	Value		string		`json:"value,omitempty"`
	Operator	string		`json:"operator,omitempty"`
	DgraphType	string 		`json:"dgraph.type,omitempty"`
}

type Connection struct {
	Uid			string		`json:"uid,omitempty"` 
	NodeNumber	string		`json:"node_number,omitempty"` //12
	Port		string		`json:"port,omitempty"` // output_1
	DgraphType	string 		`json:"dgraph.type,omitempty"`
}

type InputOutput struct {
	Uid			string			`json:"uid,omitempty"` 
	NodeUID		string			`json:"node_uid,omitempty"`
	Name		string			`json:"name,omitempty"` //input_1
	Type		string			`json:"type,omitempty"`   //input / output
	Connections []*Connection	`json:"connections,omitempty"`
	DgraphType	string 			`json:"dgraph.type,omitempty"`	
}

type Node struct {
	Uid				string			`json:"uid,omitempty"`
	ModuleUID		string			`json:"module_uid,omitempty"`
	Id				int				`json:"id,omitempty"`
	Name			string			`json:"name,omitempty"`
	Data 			Data			`json:"data,omitempty"`
	InputsOutputs	[]*InputOutput	`json:"inputs_outputs,omitempty"`
	Class			string			`json:"class,omitempty"`
	Html			string			`json:"html,omitempty"`
	Typenode		bool			`json:"typenode"`
	PosX			float32			`json:"pos_x,omitempty"`
	PosY			float32			`json:"pos_y,omitempty"`
	DgraphType		string 			`json:"dgraph.type,omitempty"`
}

type Module struct {
	Uid			string		`json:"uid,omitempty"`
	Owner		string		`json:"owner,omitempty"`
	Name		string		`json:"name,omitempty"`
	DgraphType	string 		`json:"dgraph.type,omitempty"`
}

// User model
type User struct {
	Uid        	string		`json:"uid,omitempty"`
	Username	string 		`json:"username,omitempty"`
	Password   	string  	`json:"password,omitempty"`
	DgraphType 	string    	`json:"dgraph.type,omitempty"`
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code
	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}


/******************************************************************************
********************************* Start database ******************************
******************************************************************************/
func dbCreateUser(username string, password string) string {
	var result string
	result = "USERNAME_ALREADY_REGISTERED"

	//Check if exists an user with this username
	users := getUsersByUsername(username)
	if len(users) > 0 {
		return result
	}

	dg, cancel := getDgraphClient()
	defer cancel()
	
	user := User{
		Username: username,
		Password: password,
		DgraphType: "User",
	}

	ou := &api.Operation{}
	ou.Schema = `
		username: string @index(exact) .
		password: string .
		type: string .
		type User {
			username: string
			password: string
		}
	`
	ctx := context.Background()
	if err := dg.Alter(ctx, ou); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	ub, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = ub
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}
	
	result = "USER_NOT_CREATED"
	//Get created program uid
	for _, value := range response.Uids {
		result = value
	}

	return result
}

func getAllUsers(username string) string {
	dg, cancel := getDgraphClient()
	defer cancel()

	vars := make(map[string]string)
	vars["$enusername"] = username
	q := `query wanghausers($enusername: string){
		users(func: eq(username, $enusername)){
			uid
			expand(_all_)
		}
	}`

	ctx := context.Background()

	resp, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]User `json:"users"`
	}

	var r arrays
	err = json.Unmarshal([]byte(resp.Json), &r)
	if err != nil{
		log.Println(err)
	}

	//Return users
	fmt.Printf("response: %v",resp)
	return string(resp.Json)
}

func getUsersByUsername(username string) []User {
	dg, cancel := getDgraphClient()
	defer cancel()

	vars := make(map[string]string)
	vars["$enusername"] = username
	q := `query wanghausers($enusername: string){
		users(func: eq(username, $enusername)){
			uid
			expand(_all_)
		}
	}`

	ctx := context.Background()

	resp, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]User `json:"users"`
	}

	var r arrays
	err = json.Unmarshal([]byte(resp.Json), &r)
	if err != nil{
		log.Println(err)
	}

	//Return users as string
	//return string(resp.Json)

	//Return User array
	return r.Uids
}

func dbDeleteAnyByUidType(uid string, borrar string) bool {
	dg, cancel := getDgraphClient()
	defer cancel()
	ctx := context.Background()

	d := map[string]string{"uid":uid}
	ub, err := json.Marshal(d)

	mu := &api.Mutation{
		CommitNow: true,
		DeleteJson: ub,
	}

	resp, err := dg.NewTxn().Mutate(ctx, mu)
	if resp != nil {}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func dbGetModuleByName(name string, username string) []Module {
	dg, cancel := getDgraphClient()
	defer cancel()

	vars := make(map[string]string)
	vars["$enname"] = name
	vars["$enusername"] = username
	q := `query wanghamodules($enname: string, $enusername: string){
		modules(func: type(Module)) @filter(eq(name, $enname) and eq(owner, $enusername) ){
			uid
			expand(_all_)
		}
	}`

	ctx := context.Background()

	resp, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]Module `json:"modules"`
	}

	var r arrays
	err = json.Unmarshal([]byte(resp.Json), &r)
	if err != nil{
		log.Println(err)
	}

	return r.Uids
}

func dbCreateModule(module *Module) (string, error) {
	dg, cancel := getDgraphClient()
	defer cancel()

	new_module := Module{
		Name: module.Name,
		Owner: module.Owner,
		DgraphType: "Module",
	}
	
	mo := &api.Operation{}
	mo.Schema = `
		name: string @index(exact) .
		owner: string @index(exact) . 
		type: string .
		type Module {
			name:		string
			owner: 	string
		}
	`

	ctx := context.Background()
	if err := dg.Alter(ctx, mo); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	mb, err := json.Marshal(new_module)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = mb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	var uid string
	uid = ""
	//Get created module uid
	for _, value := range response.Uids {
		uid = value
	}

	return uid, nil
}

func dbUserGetModules(username string) []*Module  {
	dg, cancel := getDgraphClient()
	defer cancel()

	vars := make(map[string]string)
	vars["$username"] = username
	q := `query usermodules($username: string){
		modules(func: type(Module)) @filter(eq(owner, $username)) {
			uid
			expand(_all_)
		}
	}`

	ctx := context.Background()

	resp, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]*Module `json:"modules,omitempty"`
	}

	var modules arrays

	err = json.Unmarshal([]byte(resp.Json), &modules)
	if err != nil{
		log.Println(err)
	}

	return modules.Uids
}

func dbDeleteModule(uid string) int {

	//Firts delete nodes
	nodes := dbModuleGetNodes(uid)
	for _, node := range nodes {
		dbDeleteAnyByUidType(node.Uid, "Node")
	}

	dg, cancel := getDgraphClient()
	defer cancel()
	ctx := context.Background()

	d := map[string]string{"uid":uid, "dgraph.type":"Module"}
	ub, err := json.Marshal(d)

	mu := &api.Mutation{
		CommitNow: true,
		DeleteJson: ub,
	}

	assign, err := dg.NewTxn().Mutate(ctx, mu)
	log.Println(assign.Metrics.NumUids)
	if err != nil{
		return 0
	}
	return 1
}

func dbCreateNode(node *Node) (Node, error) {
	dg, cancel := getDgraphClient()
	defer cancel()

	no := &api.Operation{}
	no.Schema = `
		module_uid: string @index(exact) .
		name: string @index(exact) .
		id: int .
		class: string .
		html: string .
		typenode: bool .
		pos_x: float .
		pos_y: float .
		value: string .
		operator: string .
		port: string .
		input: string .
		type: string .
		Data: uid @reverse .
		InputOutput: [uid] @reverse .
		Connection: [uid] .
		node_number: string .
		port: string .
		type Node {
			module_uid: string
			id: int
			name: string
			class: string
			html: string
			typenode: bool
			pos_x: float
			pos_y: float
			Data: Data
			InputOutput: [InputOutput]
		}
		type Data {
			name:	string
			value: string
			operator: string
		}
		type InputOutput {
			name: string
			type: string
			Connection: [Connection]
		}
		type Connection {
			node_number: string
			port: string
		}
	`

	ctx := context.Background()
	if err := dg.Alter(ctx, no); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	nb, err := json.Marshal(node)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = nb
	resp, err := dg.NewTxn().Mutate(ctx, mu)
	if (err != nil && resp == nil)  {
		log.Fatal(err)
	}

	id := strconv.Itoa(node.Id)
	vars := make(map[string]string)
	vars["$moduleuid"] = node.ModuleUID
	vars["$id"] = id
	q := `query getcreatednode($moduleuid: string, $id: string){
		nodes(func: type(Node)) @filter(eq(module_uid, $moduleuid) and eq(id, $id) ){
			uid
			expand(_all_)
			data{
				uid
				expand(_all_)
			}
			inputs_outputs{
				uid
				expand(_all_)
				connections{
					uid
			  		expand(_all_)
				}
		  	}
		}
	}`

	resp1, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]Node `json:"nodes"`
	}

	var r arrays
	err = json.Unmarshal([]byte(resp1.Json), &r)
	if err != nil{
		log.Println(err)
	}

	if len(r.Uids) > 0 {
		return r.Uids[0], nil
	}

	return Node{}, nil
}

func dbModuleGetNodes(module_uid string) []*Node {
	dg, cancel := getDgraphClient()
	defer cancel()

	vars := make(map[string]string)
	vars["$module_uid"] = module_uid
	q := `query modulenodes($module_uid: string){
		nodes(func: type(Node)) @filter(eq(module_uid, $module_uid)) {
			uid
			expand(_all_)
			data{
				uid
				expand(_all_)
			}
			inputs_outputs{
				uid
				expand(_all_)
				connections{
					uid
			  		expand(_all_)
				}
		  	}
		}
	}`

	ctx := context.Background()

	resp, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]*Node `json:"nodes,omitempty"`
	}

	var nodes arrays

	err = json.Unmarshal([]byte(resp.Json), &nodes)
	if err != nil{
		log.Println(err)
	}

	log.Println(nodes)

	return nodes.Uids
}

func dbDeleteNode(uid string) bool {
	dg, cancel := getDgraphClient()
	defer cancel()

	vars := make(map[string]string)
	vars["$uid"] = uid
	q := `query getnodes($uid: string){
				nodes(func: type(Node)) @filter(uid($uid)) {
			uid
			expand(_all_)
			data{
				uid
				expand(_all_)
			}
			inputs_outputs{
				uid
				expand(_all_)
				connections{
					uid
			  		expand(_all_)
				}
		  	}
		}
	}`

	ctx := context.Background()

	resp, err := dg.NewTxn().QueryWithVars(ctx,q,vars)
	if err != nil {
		log.Println(err)
	}

	type arrays struct{
		Uids	[]*Node `json:"nodes,omitempty"`
	}

	var nodes arrays
	var deleted bool
	deleted = false

	err = json.Unmarshal([]byte(resp.Json), &nodes)
	if err != nil{
		log.Println(err)
	}

	//Get nodes
	for _, node := range nodes.Uids {
		//Delete node data
		data_uid := node.Data.Uid
		if data_uid != ""{
			dbDeleteAnyByUidType(data_uid, "Data")
		}
		//Delete node inputs_outputs
		inputs_outputs := node.InputsOutputs
		for _, input_output := range inputs_outputs{
			input_output_uid := input_output.Uid

			//Delete input_output connections
			connections := input_output.Connections
			for _, connection := range connections{
				connection_uid := connection.Uid
				if connection_uid != "" {
					dbDeleteAnyByUidType(connection_uid, "Connection")
				}
			}

			//Delete this input_output
			if input_output_uid != "" {
				dbDeleteAnyByUidType(input_output_uid, "InputOutput")
			}
		}

		//Finally delete this node 
		node_uid := node.Uid 
		if node_uid != "" {
			deleted = dbDeleteAnyByUidType(node_uid, "Node")
		}
	}
	return deleted
}

func dbNodeUpdatePosition(node_uid string, pos_x float32, pos_y float32) bool {

	dg, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()

	mu := &api.Mutation{
		CommitNow: true,
	}

	t1 := fmt.Sprintf("<%s> <pos_x> \"%g\" .",node_uid,pos_x)
	t2 := fmt.Sprintf("<%s> <pos_y> \"%g\" .",node_uid,pos_y)
	t := fmt.Sprintf(t1+"\n"+t2)
	mu.SetNquads = []byte(t)

	_,err := dg.NewTxn().Mutate(ctx,mu)
	if err != nil{
		log.Println(err)
		return false
	}

	return true
}


func dbUpdateData(data *Data) bool {

	dg, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()

	mu := &api.Mutation{
		CommitNow: true,
	}

	t1 := fmt.Sprintf("<%s> <name> \"%s\" .",data.Uid,data.Name)
	t2 := fmt.Sprintf("<%s> <value> \"%s\" .",data.Uid,data.Value)
	t3 := fmt.Sprintf("<%s> <operator> \"%s\" .",data.Uid,data.Operator)
	t := fmt.Sprintf(t1+"\n"+t2+"\n"+t3)
	mu.SetNquads = []byte(t)

	_,err := dg.NewTxn().Mutate(ctx,mu)
	if err != nil{
		log.Println(err)
		return false
	}

	return true
}

func dbcreateConnection(connection *Connection) string {
	dg, cancel := getDgraphClient()
	defer cancel()

	co := &api.Operation{}
	co.Schema = `
		node_number: string .
		port: string .
		type Connection {
			node_number: string
			port: string
		}
	`

	ctx := context.Background()
	if err := dg.Alter(ctx, co); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	cb, err := json.Marshal(connection)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = cb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	var uid string
	uid = ""
	//Get created program uid
	for _, value := range response.Uids {
		uid = value
	}

	return uid
}

func dbInputOutputAddConnection(input_output_uid string, connection_uid string){
	dg, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()

	mu := &api.Mutation{
		CommitNow: true,
	}

	t := fmt.Sprintf("<%s> <connections> <%s> .",input_output_uid,connection_uid)
	mu.SetNquads = []byte(t)

	assign,err := dg.NewTxn().Mutate(ctx,mu)
	if assign == nil {}
	if err != nil{
		log.Println(err)
	}
}

func dbDeleteConnection(parent_uid string, connection *Connection){
	dg, cancel := getDgraphClient()
	defer cancel()
	ctx := context.Background()

	//Remove relation between OutputConnection and Connection
	type UidStruct struct {
		Uid string	`json:"uid"`
	}
	connection_uids := [1]UidStruct {{connection.Uid}} // [{uid:"0x4"}]
	d := map[string]interface{}{
        "uid": parent_uid,
        "connections": &connection_uids,
    }

    pb, err := json.Marshal(d)
    if err != nil {
        log.Fatal(err)
    }

    mu := &api.Mutation{
        CommitNow:  true,
        DeleteJson: pb,
    }	

	resp, err := dg.NewTxn().Mutate(ctx, mu)
	if resp == nil {}

	dbDeleteAnyByUidType(connection.Uid, "Connection");
}
/******************************************************************************
********************************* End database ********************************
******************************************************************************/



/******************************************************************************
********************************* Start Api Rest ******************************
******************************************************************************/

/******************************* Start Users ********************************/
type UserRequest struct {
	*User
}

type custom_error struct {
	Field 	string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type LoginResponse struct {
	Token 		string 			`json:"token,omitempty"`
	Username	string			`json:"username,omitempty"` 
	Errors 		[]custom_error	`json:"errors,omitempty"`
}

func (rd *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.User == nil {
		return errors.New("missing required User fields.")
	}
	return nil
}

// CreateArticle persists the posted Article and returns it
// back to the client as an acknowledgement.
func SignIn(w http.ResponseWriter, r *http.Request) {
	data := &UserRequest{}
	var errors []custom_error
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp := &LoginResponse{}

		if(data.Username == ""){
			errors = append(errors, custom_error {
				Field: "username", 
				Message: "El nombre de usuario es obligatorio",
			})
		}
		if(data.Password == ""){
			errors = append(errors, custom_error {
				Field: "password", 
				Message: "La contraseña es obligatoria",
			})
		}

		//Get user from database
		users := getUsersByUsername(data.Username)
		if len(users) > 0 {
			user := users[0]
			if user.Password == data.Password{
				resp.Username = user.Username
				resp.Token = "12345"
			}else{
				errors = append(errors, custom_error {
					Field: "password", 
					Message: "Contraseña incorrecta",
				})
			}
		}else if data.Username != "" {
			errors = append(errors, custom_error {
				Field: "username", 
				Message: "Este nombre de usuario no se encuentra registrado",
			})
		}

	resp.Errors = errors

	render.Status(r, http.StatusCreated)
	render.Render(w, r, resp)
}
/********************************* End Users *********************************/

/***************************** Start Modules *********************************/
type ModuleRequest struct {
	Module	*Module		`json:"module,omitempty"`
	Token	string		`json:"token,omitempty"`
}
func (a *ModuleRequest) Bind(r *http.Request) error {

	if a.Module == nil {
		return errors.New("missing required Module fields.")
	}

	if a.Module.Name == "" {
		return errors.New("missing required Module Name.")
	}

	return nil
}

type SearchModuleResponse struct {
	Success 	bool 	`json:"success,omitempty"`
	Isset		bool	`json:"isset,omitempty"` 
}

func (rd *SearchModuleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func SearchModuleByName(w http.ResponseWriter, r *http.Request){
	data := &ModuleRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	modules := dbGetModuleByName(data.Module.Name, data.Module.Owner)

	var isset bool
	isset = true
	if(len(modules) <= 0){
		isset = false
	}

	resp := &SearchModuleResponse{Success: true, Isset: isset}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, resp)
}

type CreateModuleResponse struct {
	Uid 		string 	`json:"uid,omitempty"`
}

func (rd *CreateModuleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func CreateModule(w http.ResponseWriter, r *http.Request){
	data := &ModuleRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	uid, _ := dbCreateModule(data.Module)

	resp := &CreateModuleResponse{Uid: uid}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, resp)
}

type ListModuleRequest struct {
	Username	string		`json:"username,omitempty"`
	Token		string		`json:"token,omitempty"`
}

func (a *ListModuleRequest) Bind(r *http.Request) error {

	if a.Username == "" {
		return errors.New("missing required Username field.")
	}

	if a.Token == "" {
		return errors.New("missing required Token field.")
	}

	return nil
}

func ListModules(w http.ResponseWriter, r *http.Request) {
	data := &ListModuleRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	modules := dbUserGetModules(data.Username)
	resp := &ModuleListResponse{Success: true}

	if err := render.RenderList(w, r, NewModuleListResponse(modules)); err != nil {
		render.Render(w, r, resp)
		render.Render(w, r, ErrRender(err))
		return
	}
}

type ModuleListResponse struct {
	Success		bool 					`json:"success,omitempty"`
	//Modules		[]render.Renderer		`json:"modules,omitempty"` 
}

func (rd *ModuleListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewModuleListResponse(modules []*Module) []render.Renderer {
	list := []render.Renderer{}
	for _, module := range modules {
		list = append(list, NewModuleResponse(module))
	}
	return list
}

type ModuleResponse struct {
	*Module
	Nodes	[]*Node	`json:"nodes,omitempty"`
}

func (rd *ModuleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Module.Owner = ""
	rd.Nodes = dbModuleGetNodes(rd.Module.Uid)
	return nil
}

func NewModuleResponse(module *Module) *ModuleResponse {
	resp := &ModuleResponse{Module: module}
	return resp
}


// DeleteModule removes an existing Module from our persistent store.
func DeleteModule(w http.ResponseWriter, r *http.Request) {
	
	module_uid := chi.URLParam(r, "moduleUID")

	deleted := dbDeleteModule(module_uid)

	if(deleted <= 0){
		render.Render(w, r, ErrNotFound)
		return
	}else{

		resp := &DeleteModuleResponse{Deleted: true, Uid:module_uid}

		render.Status(r, http.StatusAccepted)
		render.Render(w, r, resp)

	}
}
type DeleteModuleResponse struct {
	Deleted		bool	`json:"deleted,omitempty"`
	Uid 		string 	`json:"uid,omitempty"`
}

func (rd *DeleteModuleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}


// ClearModule delete all nodes from an existing Module from our persistent store.
func ClearModule(w http.ResponseWriter, r *http.Request) {
	module_uid := chi.URLParam(r, "moduleUID")
	nodes := dbModuleGetNodes(module_uid)
	for _, node := range nodes {
		if node.Uid != "" {
			dbDeleteAnyByUidType(node.Uid, "Node")
		}
	}
	render.Status(r, http.StatusAccepted)
}
/****************************** End Modules **********************************/

/***************************** Start Nodes ***********************************/
type NodeRequest struct {
	Node			*Node			`json:"node,omitempty"`
	Token			string			`json:"token,omitempty"`
}

func (a *NodeRequest) Bind(r *http.Request) error {
	if a.Node == nil {
		return errors.New("missing required Node fields.")
	}
	return nil
}

func CreateNode(w http.ResponseWriter, r *http.Request){
	data := &NodeRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	node, _ := dbCreateNode(data.Node)
	resp := &CreateNodeResponse{Created: true, Node:node}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, resp)
}

type CreateNodeResponse struct {
	Created		bool	`json:"created,omitempty"`
	Node 		Node 	`json:"node,omitempty"`
}

func (rd *CreateNodeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// DeleteNode removes an existing Node from our persistent store.
func DeleteNode(w http.ResponseWriter, r *http.Request) {
	
	node_uid := chi.URLParam(r, "nodeUID")

	deleted := dbDeleteNode(node_uid)

	if(!deleted){
		render.Render(w, r, ErrNotFound)
		return
	}else{

		resp := &DeleteNodeResponse{Deleted: deleted, Uid:node_uid}

		render.Status(r, http.StatusAccepted)
		render.Render(w, r, resp)

	}
}
type DeleteNodeResponse struct {
	Deleted		bool	`json:"deleted,omitempty"`
	Uid 		string 	`json:"uid,omitempty"`
}

func (rd *DeleteNodeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type PositionNodeRequest struct {
	PosX	float32		`json:"pos_x,omitempty"`
	PosY	float32		`json:"pos_y,omitempty"`
}

func (a *PositionNodeRequest) Bind(r *http.Request) error {
	return nil
}

func PositionNode(w http.ResponseWriter, r *http.Request){
	node_uid := chi.URLParam(r, "nodeUID")

	data := &PositionNodeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	
	updated := dbNodeUpdatePosition(node_uid, data.PosX, data.PosY)

	resp := &PositionNodeResponse{Updated: updated, Uid:node_uid}
	render.Status(r, http.StatusAccepted)
	render.Render(w, r, resp)
}

type PositionNodeResponse struct {
	Updated		bool	`json:"updated,omitempty"`
	Uid 		string 	`json:"uid,omitempty"`
}

func (rd *PositionNodeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type DataNodeRequest struct {
	NodeData	*Data		`json:"node_data,omitempty"`
}

func (a *DataNodeRequest) Bind(r *http.Request) error {
	if (a.NodeData == nil) {
		return errors.New("missing required Data fields.")
	}
	return nil
}

func DataNode(w http.ResponseWriter, r *http.Request){

	data := &DataNodeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	
	updated := dbUpdateData(data.NodeData)

	resp := &DataNodeResponse{Updated: updated}
	render.Status(r, http.StatusAccepted)
	render.Render(w, r, resp)
}

type DataNodeResponse struct {
	Updated		bool	`json:"updated,omitempty"`
}
func (rd *DataNodeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type CreateConnectionRequest struct {
	OutputConnection		*Connection		`json:"output_connection,omitempty"`
	InputConnection			*Connection		`json:"input_connection,omitempty"`
	OutputInputOutputUID	string 			`json:"output_input_output_uid,omitempty"`
	InputInputOutputUID		string 			`json:"input_input_output_uid,omitempty"`
	Token					string			`json:"token,omitempty"`
}

func (a *CreateConnectionRequest) Bind(r *http.Request) error {
	if (a.OutputConnection == nil || a.InputConnection == nil) {
		return errors.New("missing required Connection fields.")
	}
	return nil
}

func CreateConnection(w http.ResponseWriter, r *http.Request){
	data := &CreateConnectionRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	output_connection_uid := dbcreateConnection(data.OutputConnection)
	input_connection_uid := dbcreateConnection(data.InputConnection)

	dbInputOutputAddConnection(data.OutputInputOutputUID, output_connection_uid)
	dbInputOutputAddConnection(data.InputInputOutputUID, input_connection_uid)

	resp := &CreateConnectionResponse{ConnectionOutputUID: output_connection_uid, ConnectionInputUID:input_connection_uid}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, resp)
}

type CreateConnectionResponse struct {
	ConnectionOutputUID		string		`json:"connection_output_uid,omitempty"`
	ConnectionInputUID		string		`json:"connection_input_uid,omitempty"`
}

func (rd *CreateConnectionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}


type DeleteConnectionsRequest struct {
	OutputConnection			*Connection		`json:"output_connection,omitempty"`
	InputConnection				*Connection		`json:"input_connection,omitempty"`
	OutputConnectionParentId	string			`json:"output_connection_parent_uid,omitempty"`
	InputConnectionParentId		string			`json:"input_connection_parent_uid,omitempty"`
	Token						string			`json:"token,omitempty"`
}

func (a *DeleteConnectionsRequest) Bind(r *http.Request) error {
	if (a.OutputConnection == nil || a.InputConnection == nil) {
		return errors.New("missing required two connections fields.")
	}
	return nil
}

func DeleteConnections(w http.ResponseWriter, r *http.Request){
	data := &DeleteConnectionsRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	dbDeleteConnection(data.OutputConnectionParentId, data.OutputConnection);
	dbDeleteConnection(data.InputConnectionParentId, data.InputConnection);
	
	render.Status(r, http.StatusAccepted)
}

/****************************** End Nodes ************************************/

/******************************************************************************
********************************* End Api Rest ********************************
******************************************************************************/
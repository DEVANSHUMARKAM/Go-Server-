**SIMPLE WEB SERVER**</br>

This project creates a basic web server using the Go programming language. The server can handle two things.</br>
1. It shows a welcome message on a special webpage
2. It processes forms where users can submit information like their name, address, and email.</br>

1. Starting the Web Server</br>
At the start of the program, we make the server listen for requests from the internet.</br></br>
**CODE:</br>
fileServer := http.FileServer(http.Dir("./static"))</br>
http.Handle("/", fileServer)</br>
http.HandleFunc("/form", formHandler)</br>
http.HandleFunc("/hello", helloHandler)</br>
fmt.Printf("Starting Server at port 8080\n")</br>
if err := http.ListenAndServ(":8080", nil); err != nil{</br>
log.Fatal(err)</br>
}</br></br>**

fileServer: This part tells the server to show any files stored in a folder called static. So if you have an HTML file inside this folder, the server can show it.</br>
http.HandleFunc("/form", formHandler): This tells the server to look for a special page. When someone visits that page, it will use a function called formhandler to handle their request.</br>
http.HandleFunc("/hello", helloHandler): This is similar to /form, but this one looks for /hello. When someone visits that URL, it uses the helloHandler function to show a message.</br>
ListenAndServer(":8080", nil): This starts the server at localhost:8080. The server will run on your computer at port 8080, so you can open it in browser like this: HTTP://localhost:8080.</br>

2. Handling the form</br>
This part of the program is responsible for handling the data submitted by a user through a form like their name, address, and email</br></br>
**CODE:</br>
func formHandler(w HTTP:ResponseWriter, r *http.Request){</br>
if err := r.ParseForm(); err != nil{</br>
fmt.Fprintf(w, "ParseForm() err: %v", err)</br>
return</br>
}</br>
fmt.Fprintf(w, "POST  request successful")</br>
name := r.FromValue("name")</br>
address := r.FormValue("address")</br>
email := r.FormValue("email")</br>

fmt.Fprintf(w, "Name = %s\n", name)</br>
fmt.Fprintf(w, "Address = "%s\n", address)</br>
fmt.Fprintf(w, "Email = %s\n", email)</br>
}</br></br>**

r.ParseForm(): This line takes the information a user typed into the form and makes it usable by the program. If there is an error when trying to understand the form, it will show a message to the user.</br>
r.FormValue("name"): This sends a response back to the user's browser. For example, after submitting a form, the user will see "POST" request successful" and the details they submitted.</br>

3.Saying Hello(helloHandler)</br>
This function shows a simple message "hello!" when someone visits the /hello page.</br></br>
**CODE:</br>
func helloHandler(w http.ResponseWriter, r *http.Request){</br>
if r.URL.Path != "/hello"{</br>
http.Error(w, "404 not found", http.StatusNotFound)
return</br>
}</br>
if r.method != "GET"{</br>
http.Error(w, "method is not supported", http.StatusNotFound)</br>
return</br>
}</br>
fmt.Fprintf(w, "hello")</br>
}</br></br>**

r.URL.Path != "/hello": This checks if the person is visiting the correct page(/hello). If not, it shows a "404 not found" message.</br>
r.Method != "GET": This checks that the user is using the correct method to request the page. If they use a different method, like POST, the serer says it's not supported.</br></br>

Summary of Functions:</br>
1. formHandler: This function handles the form, gets the user's name, address, and email. and shows that information back to the user.</br>
2. helloHandler: This function shows simple "hello!" message wen the user visits the /hello page.

How to Run the Project:</br>
Save the code in a file like main.go.</br>
Make a folder called static and put any HTML files you want to serve in it.</br>
Run the program (go run main.go).</br>
Open your browser and go to http://localhost:8080 to see your files or http://localhost:8080/hello for the "hello!" message.</br>
Go to http://localhost:8080/form to submit a form with your name, address, and email.</br></br>


CONTRIBUTOR : DEVANSHU MARKAM












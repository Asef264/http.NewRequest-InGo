# http.NewRequest-InGo
***xhmoazedi@gmail.com****
///khaledMoazedi***go back developer///\\\

After Running the snippet :

in Output
server: GET /
client: got response!
client: status code: 200
client: response body: {"message": "hello!"}


in result notice this : 



In the first line, you see that the server is still receiving a GET request to the / path. The client
also receives a 200 response from the server, but it’s also reading and printing the Body of the
server’s response. In a more complex program, you could then take the {"message":
"hello!"} value you received as the body from the server and process it as JSON using the
(https://pkg.go.dev/encoding/json) package .

// Public non login
// get doctor simple data
http://localhost:5000/api/customer/doctor
body json sample
{
	"Email": "hendra@email.com",
	"Password": "123456"
}

// register new customer
http://localhost:5000/api/register
body json sample
{
	"Name": "Charles David",
	"Gender": "Male",
  	"Mobile": "+6281145784756",
	"Email": "hendra2@email.com",
	"Password": "123456",
	"RePassword": "123456"
}


//login and get token
http://localhost:5000/api/login

//check user
//using header Authorization Bearer <token>
http://localhost:5000/api/customer/user

// get doctor customer
//using header Authorization Bearer <token>
http://localhost:5000/api/customer/doctor

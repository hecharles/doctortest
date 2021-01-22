	Public non login 
	get doctor simple data
	GET: http://localhost:5000/api/public/doctor 

	login and get token
	POST: http://localhost:5000/api/login     
	body json sample
	{
		"Email": "hendra@email.com",
		"Password": "123456"
	}

	register new customer
	POST: http://localhost:5000/api/register    
	body json sample
	{
		"Name": "Charles David",
		"Gender": "Male",
		"Mobile": "+6281145784756",
		"Email": "hendra2@email.com",
		"Password": "123456",
		"RePassword": "123456"
	}




	check user
	using header Authorization Bearer <token>
	GET: http://localhost:5000/api/customer/user         

	get doctor customer
	using header Authorization Bearer <token>
	GET: http://localhost:5000/api/customer/doctor

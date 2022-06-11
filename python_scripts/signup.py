import requests

URL = "http://localhost:8080/signup"

data = {
    "name": "Samirs shop",
    "email": "samir@email.com",
    "password": "I am here",
    "phone_number": "+213795524594"
}

res = requests.post(URL, json=data)
print(res.status_code)
print(res.content)

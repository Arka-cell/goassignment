import requests

URL = "http://localhost:8080/login"

data = {
    "email": "samir@email.com",
    "password": "I am here",
}

res = requests.post(URL, json=data)
print(res.status_code)
print(res.content)

import requests

URL = "http://localhost:8080/login"

data = {
    "email": "samir@email2.com",
    "password": "Iamhere",
}

res = requests.post(URL, json=data)
print(res.status_code)
print(res.content)

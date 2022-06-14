import requests

URL = "http://localhost:8080/signup"

data = {
    "name": "Michele ",
    "email": "michele@email.com",
    "password": "Iamhere",
    "phone_number": "+213795524593"
}

res = requests.post(URL, json=data)
print(res.status_code)
print(res.content)

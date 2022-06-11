import requests

TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTQ5NDUwNzgsInNob3BfaWQiOjF9.cIT2DwwWRXu8ZXTJmQPO8mu1tVw9Wso968X31bmY9GQ"

HEADERS = {
    "Authorization": f"JWT {TOKEN}"
}

URL = "http://localhost:8080/products"

FILES = [
  ('file_input',('image.jpg',open('image.jpg','rb'),'image/*')),
]

res = requests.post(URL, headers=HEADERS, files=FILES)
print(res.status_code)
print(res.content)
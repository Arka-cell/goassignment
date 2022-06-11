import requests

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTQ5NzIxNTcsInNob3BfaWQiOjF9.WC8oA_7qjLQhogit-a1QN6fLlvptMShoxk3HiED-sRk"

HEADERS = {
    "Authorization": f"{token}"
}

data = {
    "title": "Sephora lipsticks",
    "description": "It will make your lips look good",
    "shop_id": 1
    
}

res = requests.post(f"http://localhost:8080/products?token={token}", json=data)
print(res.status_code)
print(res.content)

res = requests.get("http://localhost:8080/products")
print(res.status_code)
print(res.content)

res = requests.get("http://localhost:8080/products/1")
print(res.status_code)
print(res.content)
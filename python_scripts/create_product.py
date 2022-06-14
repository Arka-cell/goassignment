import requests

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTUxOTYyOTAsInNob3BfaWQiOjF9.7sJGocuwaW8_svlNQvDJweI0I1IOUvNT7iAdPzmOzQk"

HEADERS = {
    "Authorization": f"{token}"
}

data = {
    "title": "Sephora lipsticks ssx",
    "description": "It will make your lips look good",
    "shop_id": 1
    
}

res = requests.post(f"http://localhost:8080/products?token={token}", json=data)
print(res.status_code)
print(res.content)
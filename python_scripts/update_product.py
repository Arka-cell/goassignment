import requests

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTUxOTYyOTAsInNob3BfaWQiOjF9.7sJGocuwaW8_svlNQvDJweI0I1IOUvNT7iAdPzmOzQk"
HEADERS = {
    "Authorization": f"{token}"
}

data = {
    "title": "okay",
    "description": "",
    "shop_id": 1,
    "image_url": ""
    
}

res = requests.put(f"http://localhost:8080/products/1?token={token}", json=data, headers=HEADERS)
print(res.status_code)
print(res.content)

import requests

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTUwNDgyNDIsInNob3BfaWQiOjF9.-1lAVs2kBeMNPCf1JtgNR-aBuLA1kiVeCVnfh2XuiIA"

HEADERS = {
    "Authorization": f"{token}"
}




res = requests.get("http://localhost:8080/products")
print(res.status_code)
print(res.content)

res = requests.get("http://localhost:8080/products/1")
print(res.status_code)
print(res.content)
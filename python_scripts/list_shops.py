import requests

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTQ5Njk4OTgsInNob3BfaWQiOjR9.p8m6R1V5LdZjyVKOvaaFEw-RrQxD72UQ63ufAy9L7AM"

HEADERS = {
    "Authorization": f"Token {token}"
}

res = requests.get("http://localhost:8080/shops")
print(res.status_code)
print(res.content)

res = requests.get("http://localhost:8080/shops/4")
print(res.status_code)
print(res.content)

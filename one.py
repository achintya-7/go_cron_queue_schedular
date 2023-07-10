import requests

url = "http://localhost:8080"

response = requests.request("GET", url)
print(response.text)



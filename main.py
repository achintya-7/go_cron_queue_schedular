import requests

for i in range(10):
    url = "http://localhost:8080"
    
    if i == 0:
        utl = url + "?name=J"
    
    if i == 1:
        utl = url + "?name=Ma"

    if i == 2:
        utl = url + "?name=Pet"

    if i == 3:
        utl = url + "?name=Paul"

    if i == 4:
        utl = url + "?name=12345"

    if i == 5:
        utl = url + "?name=Simone"

    if i == 6:
        utl = url + "?name=Bernard"

    if i == 7:
        utl = url + "?name=Achintya"

    if i == 8:
        utl = url + "?name=123456789"

    if i == 9:
        utl = url + "?name=Decagon333"

    response = requests.request("GET", utl)
    print(response.text)



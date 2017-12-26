import requests
import json

pre_user_id = ""
for i in range(2):
    data = {
        "name": "liuzhenzhong"
    }
    res = requests.post("http://127.0.0.1:8080/users", json=data)
    print(res.text, res)

    value = json.loads(res.text)
    if pre_user_id:
        url = "http://127.0.0.1:8080/users/"+value["id"]+"/relationships/"+pre_user_id
        data = {"state": "liked"}
        res = requests.put(url, json=data)
        print(res.text, res)

        url = "http://127.0.0.1:8080/users/"+pre_user_id+"/relationships/"+value["id"]
        data = {"state": "disliked"}
        res = requests.put(url, json=data)
        print(res.text, res)

    print("------get relations:")
    url = "http://127.0.0.1:8080/users/"+value["id"]+"/relationships"
    res = requests.get(url)
    print(res.text, res)

    if pre_user_id:
        url = "http://127.0.0.1:8080/users/"+pre_user_id+"/relationships"
        res = requests.get(url)
        print(res.text, res)     

  

    pre_user_id = value["id"]
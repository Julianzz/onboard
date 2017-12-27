import requests
import json

def put_relation(u1, u2, state):
    url = "http://127.0.0.1:8080/users/"+u1+"/relationships/"+u2
    data = {"state": state}
    res = requests.put(url, json=data)
    print(res.text, res)

def get_users():
    res = requests.get("http://127.0.0.1:8080/users")
    print(res.text, res)

def test_state(state1, state2, num=2):
    pre_user_id = ""
    for i in range(num):
        data = {
            "name": "liuzhenzhong"
        }
        res = requests.post("http://127.0.0.1:8080/users", json=data)
        #print(res.text, res)

        value = json.loads(res.text)
        if pre_user_id:
            put_relation(value["id"], pre_user_id, state1)
            put_relation(pre_user_id, value["id"], state2)

        if pre_user_id:

            print("------get relations:")
            url = "http://127.0.0.1:8080/users/"+value["id"]+"/relationships"
            res = requests.get(url)
            print(res.text, res)

            url = "http://127.0.0.1:8080/users/"+pre_user_id+"/relationships"
            res = requests.get(url)
            print(res.text, res)     
        pre_user_id = value["id"]

test_state("liked", "disliked")
print("")
print("")
test_state("liked", "liked")

print("----------putwront userid")
put_relation("23", "25", "liked")

print("----------putwront state")
test_state("matched", "liked")
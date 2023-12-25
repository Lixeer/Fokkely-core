import requests
data={"id":"sk-1"}
rp=requests.post(url="http://127.0.0.1:5700/get/skill",json=data)
print(rp.text)

while True:
    rp=requests.post(url="http://127.0.0.1:5700/get/skill",json=data)
    print(rp.text)

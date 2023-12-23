import requests
data={"f":{ "each":"hello"},"b":["1","2"]}
rp=requests.post(url="http://127.0.0.1:5700/process",json=data)
print(rp.text)

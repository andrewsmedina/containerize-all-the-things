$ docker run -d -p 8000 python:3.5.1-alpine python3 -m http.server

$ docker ps | grep python
f98ea7d0213e        python:3.5.1-alpine   "python3 -m http.serv"   About a minute ago   Up About a minute   0.0.0.0:32769->8000/tcp   pedantic_panini

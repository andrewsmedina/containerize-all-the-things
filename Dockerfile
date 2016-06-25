FROM python:3.5.1-alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY requirements.txt /usr/src/app/
RUN pip  install -v -r requirements.txt

COPY app.py /usr/src/app

CMD ["python", "app.py"]

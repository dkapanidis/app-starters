FROM python:3-alpine

RUN adduser -D python
USER python
WORKDIR /home/python

WORKDIR /app
COPY --chown=python:python . .
EXPOSE 5000
CMD [ "python3", "-u" , "app.py"]

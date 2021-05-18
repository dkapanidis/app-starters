FROM python:3-alpine

RUN adduser -D python
USER python
WORKDIR /home/python

WORKDIR /app
COPY --chown=python:python . .
RUN python -m nose2
ARG DEBUG=0
EXPOSE 5000
CMD [ "python3", "-u" , "src/app.py"]


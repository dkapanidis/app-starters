FROM python:3-alpine

RUN adduser -D python
USER python
WORKDIR /home/python

WORKDIR /app
COPY --chown=python:python requirements.txt .
RUN pip install --user -r requirements.txt
COPY --chown=python:python . .
RUN python -m nose2
ARG DEBUG=0
ENV FLASK_DEBUG $DEBUG
ENV FLASK_APP=src/app.py
EXPOSE 5000
CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]


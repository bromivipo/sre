FROM python:3.9-slim

RUN pip install requests
RUN pip install flask

COPY app.py /app.py

EXPOSE 8080

CMD ["python", "app.py"]


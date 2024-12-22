from flask import Flask
import os
import requests

app = Flask(__name__)

@app.route('/')
def version():
    requests.get('http://metrics:5000/v1')
    pod_name = os.getenv('POD_NAME', 'unknown')
    version_info = f"Version 1.0, Pod: {pod_name}"
    return version_info, 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)


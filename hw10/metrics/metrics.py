from prometheus_client import start_http_server, Counter
from flask import Flask, request

app = Flask(__name__)

http_requests_total_v1 = Counter('http_requests_total_v1', 'Total number of HTTP requests')
http_requests_total_v2 = Counter('http_requests_total_v2', 'Total number of HTTP requests')

@app.route('/v1')
def v1():
    http_requests_total_v1.inc()
    return ""

@app.route('/v2')
def v2():
    http_requests_total_v2.inc()
    return ""

if __name__ == '__main__':
    start_http_server(8000)
    app.run(host='0.0.0.0', port=5000)


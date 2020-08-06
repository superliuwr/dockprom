import http.server
import logging
import random
import time
from prometheus_client import start_http_server, Counter, Gauge, Summary, Histogram

# Counters
REQUESTS = Counter('hello_worlds_total', 'Hello Worlds requested.')
SALES = Counter('hello_world_sales_euro_total', 'Euros made serving Hello World.')
EXCEPTIONS = Counter('hello_world_exceptions_total', 'Exceptions serving Hello World.')

# Gauges
INPROGRESS = Gauge('hello_worlds_inprogress', 'Number of Hello Worlds in progress.')
LAST = Gauge('hello_world_last_time_seconds', 'The last time a Hello World was served.')

# Summaries
LATENCY = Summary('hello_world_latency_seconds', 'Time for a request Hello World.')

# Histograms
LATENCY_H = Histogram('hello_world_latency_seconds_h', 'Time for a request Hello World.')

class MyHandler(http.server.BaseHTTPRequestHandler):
  @LATENCY_H.time()
  def do_GET(self):
    logging.warning('MyHandler called once')
    start = time.time()
    REQUESTS.inc()
    INPROGRESS.inc()
    euros = random.random()
    SALES.inc(euros)
    # with EXCEPTIONS.count_exceptions():
    #   if random.random() < 0.2:
    #     raise Exception
    self.send_response(200)
    self.end_headers()
    self.wfile.write(b"Hello World")
    LAST.set_to_current_time()
    INPROGRESS.inc()
    LATENCY.observe(time.time() - start)

if __name__ == "__main__":
  start_http_server(8100)
  server = http.server.HTTPServer(('0.0.0.0', 8000), MyHandler)
  server.serve_forever()
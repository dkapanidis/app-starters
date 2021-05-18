import http.server
import socketserver

PORT = 5000

class Handler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        self.wfile.write("Hello world".encode('utf-8'))

with socketserver.TCPServer(("", PORT),  Handler) as httpd:
    print("serving at port", PORT)
    httpd.serve_forever()


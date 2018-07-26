#!/usr/bin/env python
"""
Very simple HTTP server in python.
Usage::
    ./dummy-web-server.py [<port>]
Send a GET request::
    curl http://localhost
Send a HEAD request::
    curl -I http://localhost
Send a POST request::
    curl -d "foo=bar&bin=baz" http://localhost
"""
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
import SocketServer
import json

class S(BaseHTTPRequestHandler):
    def _set_headers(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()

    def do_GET(self):
        output = "{}"

        if self.path.startswith("/_matrix/app/unstable/thirdparty/protocol/"):
            output = self.retreive_protocol_def()
        elif self.path.startswith("/_matrix/app/unstable/thirdparty/location"):
            output = self.location()
        elif self.path.startswith("/_matrix/app/unstable/thirdparty/user"):
            output = self.user()

        self._set_headers()
        self.wfile.write(output)

    def retreive_protocol_def(self):
        return '''{
  "user_fields": [
    "network",
    "nickname"
  ],
  "location_fields": [
    "network",
    "channel"
  ],
  "icon": "mxc://example.org/aBcDeFgH",
  "field_types": {
    "network": {
      "regexp": "([a-z0-9]+\\\.)*[a-z0-9]+",
      "placeholder": "irc.example.org"
    },
    "nickname": {
      "regexp": "[^\\\s]+",
      "placeholder": "username"
    },
    "channel": {
      "regexp": "#[^\\\s]+",
      "placeholder": "#foobar"
    }
  },
  "instances": [
    {
      "desc": "Freenode",
      "icon": "mxc://example.org/JkLmNoPq",
      "fields": {
        "network": "freenode.net"
      }
    }
  ]
}
'''

    def user(self):
        return '''[
  {
    "userid": "@_gitter_jim:matrix.org",
    "protocol": "gitter",
    "fields": {
      "user": "jim"
    }
  }
]
'''

    def location(self):
        return '''[
  {
    "alias": "#freenode_#matrix:matrix.org",
    "protocol": "irc",
    "fields": {
      "network": "freenode",
      "channel": "#matrix"
    }
  }
]
'''

    def do_HEAD(self):
        self._set_headers()
        
    def do_POST(self):
        # Doesn't do anything with posted data
        self._set_headers()
        content_len = int(self.headers.getheader('content-length', 0))
        post_body = self.rfile.read(content_len)
        json_body = json.loads(post_body)
        print json.dumps(json_body, indent=2, sort_keys=False)
        self.wfile.write("<html><body><h1>POST!</h1></body></html>")
        
def run(server_class=HTTPServer, handler_class=S, port=80):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    print 'Starting httpd...'
    httpd.serve_forever()

if __name__ == "__main__":
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()

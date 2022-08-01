import os, flask
MESSAGE = os.environ['MESSAGE']
PORT = int(os.environ['PORT'])

app = flask.Flask('app server')
@app.route('/')
def index():
  return MESSAGE
app.run(debug=True, host='0.0.0.0', port=PORT)

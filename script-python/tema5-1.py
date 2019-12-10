from flask import Flask
import os
app = Flask(__name__)

@app.route("/conf/env")
def env_vars():
    return dict(os.environ)

if (__name__ == "__main__"):
    app.run(port = 8080)
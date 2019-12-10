from flask import Flask
import os
app = Flask(__name__)

@app.route("/env/<name>/<value>", methods=['POST'])
def env_vars(name, value):
    with open(os.environ["HOME"]+"/.bashrc", "a") as new_Var:
        new_Var.write("export "+name + "=" + value)
    return "Variável adicionada !"

if (__name__ == "__main__"):
    app.run(port = 8080)
from flask import Flask, request
import os
import json


app = Flask(__name__)


@app.route("/discounts")
def sale():
    cost = int(request.args.get('cost'))
    end_cost = get_sale(cost)
    return  json.dumps({'result': end_cost})


def get_sale(cost):
    result = 0.5*cost
    return result


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=int(os.getenv('LISTEN_PORT', 5000)))

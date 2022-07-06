from flask import Flask, render_template
import os


app = Flask(__name__, template_folder='template', static_folder='static')


flag = open('static/.k0u5927a9t0ca025c98g2187i7dc2e46c3da945cp', 'w')
flag.write(os.getenv('FLAG'))
flag.close()


@app.route('/')
def home():
    return render_template('index.php', xss='PoC')


if __name__ == "__main__":
    port = int(os.environ.get('PORT', 5002))
    app.run(debug=True, host='0.0.0.0', port=port)

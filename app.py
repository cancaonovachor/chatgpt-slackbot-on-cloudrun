import os
from flask import Flask, request, jsonify
from slack_sdk.signature import SignatureVerifier
import openai

# Initialize OpenAI API
openai.api_key = os.environ["OPENAI_API_KEY"]

# Initialize Flask App
app = Flask(__name__)


def verify_request(req):
    signature_verifier = SignatureVerifier(os.environ["SLACK_SIGNING_SECRET"])
    return signature_verifier.is_valid(
        body=req.get_data(as_text=True),
        timestamp=req.headers.get("X-Slack-Request-Timestamp"),
        signature=req.headers.get("X-Slack-Signature")
    )


@app.route("/events", methods=["POST"])
def handle_event():
    if not verify_request(request):
        return jsonify({"error": "Invalid request"}), 400

    payload = request.json
    if payload["event"].get("type") == "app_mention":
        text = payload["event"]["text"]
        reply = chat_gpt(text)
        return jsonify({"text": reply})

    return jsonify({})


def chat_gpt(text):
    response = openai.Completion.create(
        engine="gpt-3.5-turbo",
        prompt=f"{text}\n",
        temperature=0.5,
        max_tokens=100,
        top_p=1,
        frequency_penalty=0,
        presence_penalty=0
    )

    return response.choices[0].text.strip()


if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=int(os.environ.get("PORT", 8080)))

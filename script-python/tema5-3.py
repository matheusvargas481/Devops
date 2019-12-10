import os, slack, sys

def send_message():
    client = slack.WebClient(token=sys.argv[1])
    client.chat_postMessage(
    channel=sys.argv[2],
    text=os.popen("ps -e -o pid,cmd").read())

if __name__ == "__main__":
    send_message()
import os
import openai

from fastapi import FastAPI
from fastapi import Form
from pydantic import BaseModel


app = FastAPI()


@app.post("/ask/")
async def ask(api_key: str = Form(), ask_content: str = Form()):
    res = ask_chatgpt(api_key, ask_content)
    return res


def ask_chatgpt(api_key, ask_content):
    openai.api_key = api_key
    res = openai.ChatCompletion.create(
      model="gpt-3.5-turbo",
      messages=[
        {"role": "user", "content": ask_content}
      ]
    )

    try:
        answer = res.get('choices')[0].get('message').get('content')
    except:
        answer = '[出错了，请重试]'
    
    return answer
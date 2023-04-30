import os
import openai

from fastapi import FastAPI
from pydantic import BaseModel


class AskParams(BaseModel):
    api_key: str
    ask_content: str


app = FastAPI()


@app.post("/ask/")
async def ask(ask_params: AskParams):
    res = ask_chatgpt(ask_params.api_key, ask_params.ask_content)
    return res


def ask_chatgpt(api_key, ask_content):
    openai.api_key = api_key
    res = openai.ChatCompletion.create(
      model="gpt-3.5-turbo",
      messages=[
        {"role": "user", "content": ask_content}
      ]
    )

    return res
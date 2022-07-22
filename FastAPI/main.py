#!/usr/bin/env python

from database_session import session as Session
from database_session import Users
from fastapi import FastAPI
from sqlalchemy import select

app = FastAPI()

@app.get("/users/")
async def GetAllUsers():
    async with Session() as session:
        result = await session.execute(
            select(Users)
        )
        result = result.scalars().all()
    return {"data": {"user": result}}

@app.get("/users/{user_id:int}")
async def GetAllUsers(user_id: int):
    async with Session() as session:
        result = await session.execute(
            select(Users).where(Users.user_id==user_id)
        )
        result = result.scalars().first()
    return {"data": {"user": result}}
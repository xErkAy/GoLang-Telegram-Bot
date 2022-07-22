from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.asyncio import create_async_engine
from sqlalchemy.orm import declarative_base
from sqlalchemy import Column
from sqlalchemy import Integer
from sqlalchemy import String

Base = declarative_base()

engine = create_async_engine("postgresql+asyncpg://admin:admin@db:5432/gotest")
session = sessionmaker(
  engine, expire_on_commit = False, class_ = AsyncSession
)

class Users(Base):
  __tablename__ = "users"
  user_id = Column(Integer, primary_key=True)
  username = Column(String)
  language = Column(String)
  first_name = Column(String)
from fastapi import FastAPI

app = FastAPI()

@app.get("/users/{user_id}")
def get_user(user_id: str):
    return {"id": user_id, "name": "John Doe", "email": "john@example.com"}

@app.get("/users")
def list_users():
    return [{"id": "1", "name": "Alice"}, {"id": "2", "name": "Bob"}]
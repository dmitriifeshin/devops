### task_manager/app.py

from datetime import datetime
from sqlalchemy.orm import Session
from models import Task
from database import SessionLocal, engine, Base
import os

Base.metadata.create_all(bind=engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

def add_task(title, description, due_date):
    db = next(get_db())
    task = Task(
        title=title,
        description=description,
        due_date=due_date,
    )
    db.add(task)
    db.commit()
    print(f"Task '{title}' added.")

def list_tasks():
    db = next(get_db())
    tasks = db.query(Task).all()
    for task in tasks:
        status = "Done" if task.completed else "Pending"
        print(f"{task.id}: {task.title} | Due: {task.due_date} | Status: {status}")

def mark_task_done(task_id):
    db = next(get_db())
    task = db.query(Task).filter(Task.id == task_id).first()
    if task:
        task.completed = True
        db.commit()
        print(f"Task {task_id} marked as done.")
    else:
        print("Task not found.")

def delete_task(task_id):
    db = next(get_db())
    task = db.query(Task).filter(Task.id == task_id).first()
    if task:
        db.delete(task)
        db.commit()
        print(f"Task {task_id} deleted.")
    else:
        print("Task not found.")

def main():
    while True:
        print("\nTask Manager")
        print("1. Add Task")
        print("2. List Tasks")
        print("3. Mark Task as Done")
        print("4. Delete Task")
        print("5. Exit")
        choice = input("Enter choice: ")

        if choice == '1':
            title = input("Title: ")
            description = input("Description: ")
            due = input("Due Date (YYYY-MM-DD): ")
            due_date = datetime.strptime(due, "%Y-%m-%d").date()
            add_task(title, description, due_date)
        elif choice == '2':
            list_tasks()
        elif choice == '3':
            task_id = int(input("Enter task ID to mark as done: "))
            mark_task_done(task_id)
        elif choice == '4':
            task_id = int(input("Enter task ID to delete: "))
            delete_task(task_id)
        elif choice == '5':
            break
        else:
            print("Invalid choice.")

if __name__ == "__main__":
    main()
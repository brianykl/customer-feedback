from dotenv import load_dotenv 
import os
import psycopg
from psycopg import OperationalError, DatabaseError
import uuid

load_dotenv()
DATABASE_URL = os.getenv('DATABASE_URL')

def get_db_connection():
    print("DATABASE_URL:", DATABASE_URL)
    try:
        return psycopg.connect(DATABASE_URL)
    except OperationalError as e:
        print(f'error connecting to the data base{e}')

def execute_query(query, args = None, fetch = False):
    try:
        with get_db_connection() as conn:
            with conn.cursor(row_factory=psycopg.rows.dict_row) as cur:
                cur.execute(query, args)
                if fetch:
                    return cur.fetchall()
            
    except DatabaseError as e:
        print(f'database query error {e}')
            
def add_record(table, record: dict):
    columns = ', '.join(record.keys())
    values = tuple(record.values())
    placeholders = ', '.join(['%s'] * len(record))
    query = f"INSERT INTO {table} ({columns}) VALUES ({placeholders})"
    execute_query(query, values)

def get_records(table, condition, args=None):
    query = f"SELECT * FROM {table} WHERE {condition}"
    return execute_query(query, args)

def update_record(table, update: dict, condition, args = None):
    update_string = ', '.join([f'{key} = %s' for key in update.keys()])
    values = list(update.values())
    if args:
        values.extend(args)

    query = f"UPDATE {table} SET {update_string} WHERE {condition}"
    return execute_query(query, values)

def delete_record(table, condition, args = None):
    query = f"DELETE FROM {table} WHERE {condition}"
    return execute_query(query, args)

def id_generator():
    return str(uuid.uuid4())

def create_feedback_table():
    create_table_query = """ 
    CREATE TABLE IF NOT EXISTS feedback (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL,
    category TEXT NOT NULL,
    feedback_text TEXT NOT NULL,
    submission_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );"""

    try:
        with get_db_connection() as conn:
            with conn.cursor() as cur:
                cur.execute(create_table_query)
                print("feedback table created successfully")
    
    except DatabaseError as e:
        print(f'error creating feedback table: {e}')

def create_feedback_analysis_table():
    create_feedback_analysis_table_query = """
    CREATE TABLE IF NOT EXISTS feedback_analysis (
    id UUID PRIMARY KEY,
    topic VARCHAR (255),
    sentiment VARCHAR (255),
    FOREIGN KEY (id) REFERENCES feedback(id)
    );"""

    try:
        with get_db_connection() as conn:
            with conn.cursor() as cur:
                cur.execute(create_feedback_analysis_table_query)
                print("feedback_analysis table created successfully")
    
    except DatabaseError as e:
        print(f'error creating feedback_analysis table: {e}')

if __name__ == '__main__':
    create_feedback_analysis_table()


import os
import psycopg
from psycopg import OperationalError, DatabaseError


DATABASE_URL = os.getenv('DATABASE_URL')

def get_db_connection():
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
            
def add_record(table, record):
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



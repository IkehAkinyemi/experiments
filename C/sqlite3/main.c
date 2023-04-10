#include <stdio.h>
#include <stdlib.h>
#include <sqlite3.h>

int main(int argc, char* argv[]) {
    sqlite3 *db;
    char *err_msg = 0;

    int rc = sqlite3_open("test.db", &db);

    if (rc != SQLITE_OK) {
        fprintf(stderr, "Cannot open database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        return 1;
    }

    char *sql = "CREATE TABLE users(id INT, name TEXT);"
                "INSERT INTO users VALUES(1, 'John Doe');";

    rc = sqlite3_exec(db, sql, 0, 0, &err_msg);

    if (rc != SQLITE_OK ) {
        fprintf(stderr, "SQL error: %s\n", err_msg);
        sqlite3_free(err_msg);
    }

    sqlite3_close(db);
    return 0;
}

Welcome to JiraniAPI

This will be fun amigo....this should be re-written

# Setting up database
> **Assumption**: You have postgresql or mysql installed and running in  your machine
- Access the terminal and enter:
```bash
sudo -i -u postgres psql
```
## Postgres or MYSQL Environment
> Steps:
1. Create user
```bash
CREATE USER target_name WITH PASSWORD 'eg...12345';
# Enter different password
```
2. Create database
```bash
CREATE DATABASE [IF NOT EXISTS] db_name;
```
3. Grant privileges to newly created user, also we want this user to own this database.
```bash
GRANT ALL PRIVILEGES ON DATABASE database_name TO target_name;

# Changing ownership
ALTER DATABASE db_name
OWNER TO target_user;
```
# Resources

[Setting Up Google OAuth2](https://www.loginradius.com/blog/engineering/google-authentication-with-golang-and-goth/)
[Setting up Google OAuth2 V2](https://blog.boot.dev/golang/how-to-implement-sign-in-with-google-in-golang/)

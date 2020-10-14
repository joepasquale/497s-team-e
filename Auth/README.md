Authentication using JWT(JSON Web Token)

What is JWT?

JWT allows you encrpt user info and use it while sending request.

![JWT workflow](https://dzone.com/storage/temp/8886724-jwt-communication.png)

Upon successfully login, the client will receive a token from the server. The front-end will store the token locally and send subsequent requests with the token in the header.

You can think of JWT as a way to carry out authentication at potentially every request but a million times faster than database.

This is because the token can only be decrypted if you know the SECRET_KEY used to encrpt it (and of course it is on the server).

However, if you keep trying, one day you might happen to create a valid token, and this is really bad. To solve this problem, we need to periodically update the valid token pools so no matter how hard you try, the token will become invalid before you generate it.

Thanks to radis, we can do that rather easily.

The idea is to create two tokens upon login. One is "access_token" which should be put in the header. The other is called "refresh_token", which is used to get a new access_token when the "access_token" expires. When both tokens become invalid, the user will need to login again.

Ok, now is how to use it.

You can simply run the following command to start the service:
docker-compose up --build

Authentication service will be working on port: 5555.
It has five routers which all accept only POST requests:
1. localhost:5555/register
expecting input:
{
    "username": username,
    "password": password
}
Username has to be unique. Password will be encrpted using bcrypt before store into mongoDB.

2. localhost:5555/login 
expecting input:
{
    "username": username,
    "password": password
}
Upon successfully login, client will recieve:
{
    "access_token": access_token_string,
    "refresh_token": refresh_token_string
}

3. localhost:5555/check
Before you access this route, remember to store the token at the right place.
If you are using "postman", copy paste the access_token at bearer token: 
![postman](https://assets.postman.com/postman-docs/authorization-types.jpg)
What check does is it look at token and search it in redis to see if it still exists.
If the token is not a valid one, what you will recieve is
{
    "userID":  "",
    "islogin": "false",
    "status":  "invalid token",
}

If the token is valid but it has expires, you will get
{
    "userID":  "",
    "islogin": "false",
    "status":  "access token expired",
}

If the token is still valid:
{		"userID":  userID,
		"islogin": "true",
		"status":  "ok",
}

The status will be helpful for the front-end to determine what to do next.

4. localhost:5555/refresh
expect:
{
    "refresh_token": refresh_token_string,
}
If refresh_token is valid and not expired, the client will recieve:
{
    "access_token": new_access_token_string,
    "refresh_token": same_refresh_token_string
}

5. localhost:5555/logout
Yeah, you still need to give something before you can logout, and that is your token.
Similarly, keep the access_token in the header. It allows the server to delete your credentials in redis instantly.

Storing JWT at front-end: https://hasura.io/blog/best-practices-of-using-jwt-with-graphql/#basics_client_setup




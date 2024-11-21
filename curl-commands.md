# Sign Up 

curl -X POST http://localhost:8080/signup \
-H "Content-Type: application/json" \
-d '{"email": "test@example.com", "password": "password123"}'

Method: POST
URL: /signup
Body: JSON with email and password.

Response 
{
  "message": "User created successfully"
}


# Sign in  

curl -X POST http://localhost:8080/signin \
-H "Content-Type: application/json" \
-d '{"email": "test@example.com", "password": "password123"}'

Method: POST
URL: /signin
Body: JSON with email and password.

{
  "message": "Login successful",
  "token": "your-jwt-token-here"
}


curl -X GET http://localhost:8080/protected \
-H "Authorization: Bearer your-jwt-token-here"


# protected  

Method: GET
URL: /protected
Header: Authorization: Bearer <JWT_TOKEN> (replace <JWT_TOKEN> with the actual token obtained during sign-in).

{
  "message": "Welcome, user!",
  "user_id": 1
}


Failure Case : 

{
  "error": "Invalid token"
}

# Refresh Token 

curl -X POST http://localhost:8080/refresh \
-H "Content-Type: application/json" \
-d '{"refresh_token": "your-refresh-token-here"}'


Method: POST
URL: /refresh
Body: JSON with refresh_token.


{
  "message": "Token refreshed",
  "token": "new-jwt-token-here"
}

Invalid Token 

{
  "error": "Invalid refresh token"
}

# Token Expired Status

Message for Token Expired 

{
  "error": "Token has expired"
}
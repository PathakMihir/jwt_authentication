# Role Based Authentication using JWT Authentication 

The project demonstrates following features:
 
* JWT Authentication Implementation
* SignIn/Login Api Implementation
* Role Based Authentication

Data Model:

>API Requests:

1.SignIn API

```
    POST api/v1/user/signIn 
    Body: {
        "first_name":
        "last_name":
        "email":
        "phone_number":
        "password":
        "address":
        "pincode":
        "type":
    }
```    

Implementation:

User should be able to sign in and load all the details into single entry in mongoDB with all validations    

Task:
* API Defination and Models ->Done
* Validation Logic and DB insertion -> Done
* Return Response and Error Handling->Done


2.Login API
```
    POST api/v1/users/login
    Body: {
        "email":
        "password":
    }

    Response:
    {
        user_details:{}
        "token":
        "refresh_token":
    }
```    

Implementation:

User should be able to verify the password and email as per in the database and generate a new token and refresh token and return back in the response.    

3.All other API have middleware implementation which verifies the token before granting access to the particular api.
 
```
    GET api/v1/profiles/users/ + token
    
    Response:
    {
        users_list
    }

```
This api shows the use of JWT Authentication middleware...    

4.Refresh Token :
 
 ```
 
 GET api/v1/refreshToken + HTTP-Only cookie with refresh_token without expiry

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Impob25AZ21haWwuY29tIiwiRmlyc3ROYW1lIjoiamhvbiIsIkxhc3ROYW1lIjoiMjIiLCJVc2VySUQiOiIiLCJleHAiOjE2NzU1NjQyMTR9.RP-nfcl03bJf0K1tTEhXNmIxRUj6TQCCvdT9Q4ppvFc"
}

```

5.





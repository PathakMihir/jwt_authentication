# jwt_authentication
Simple JWT Authentication Implementation
Data Model:



API Requests:
1.
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

Implementation:
User should be able to sign in and load all the details into single entry in mongoDB with all validations    
Task:
    1.API Defination and Models
    2.Validation Logic and DB insertion
    3.Return Response and Error Handling


2.
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
Implementation:
User should be able to verify the password and email as per in the database and generate a new token and refresh token and return back in the response.    

3.All other API have middleware implementation which verifies the token before granting access to the particular api.





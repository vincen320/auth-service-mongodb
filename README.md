# auth-service-mongodb
test code simple microservices, you can see how i write code in go this service provides login to app, after login user get token(JWT) to identify user to other service(product-service) and provides refresh to refresh token when expired didn't implement refresh token yet, just use the same token for practice
Same as auth-service actually (i'm really just copy and paste all code), just change the model for ID from int to primitive.ObjectID
# Architecture
Controller -> Service -> Repository

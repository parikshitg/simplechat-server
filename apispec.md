# SimpleChat API Specification

## Structures

### Struct Status
```
{
    "success": bool,
    "error": string,
}
```

### Struct User
```
{
    "id": string,
    "name": string,
    "email:" string
    "photo": string,
} 
```


# API

## 1. Register

Endpoint ```/register```

### RegisterRequest

```
{
    "email": string
}
    
```

### RegisterResponse
```
{
    "status": Status,
    "auth_token": string
}
    
```


## 2. Login

Endpoint ```/login```

### LoginRequest
```
{
    "auth_token": string,
    "password": string    
}
```

### LoginResponse
```
{
    "status": Status,
    "access_token": string
}
```

## 3. User Profile

Endpoint ```/profile```

### ProfileRequest
```
{
    "access_token": string,
    "id": string
}
```

### ProfileResponse
```
{
    "status": Status,
    "profile": User,
}
```


## 4. Friend List

Endpoint ```/friends```

### FriendListRequest
```
{
    "access_token": string,
}
```

### FriendListResponse
```
{
    "status": Status,
    "friends": []User
}
```

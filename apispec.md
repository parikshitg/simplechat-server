# SimpleChat API Specification

## Structures

### 1. Status
```
{
    "success": bool,
    "error_message": string,
}
```


### 2. User
```
{
    "userid": string,
    "name": string,
    "email": string,
    "mobile": string,
    "photo": string,
    "status": string,
} 
```


### 3. Message
```
{
    "msgid": string,
    "text": string,
    "date": string
    "sender_id": string,
    "reciever_id": string,
} 
```




# API

## 1. Login

Endpoint ```POST /login```

### 1.1 LoginRequest

```
{
	"2fa": bool,
    "email": string,
    "password": string,
    "mobile": string,
    "otp": string,
    "auth_token": string,
}
    
```

### 1.2 LoginResponse
```
{
    "status": Status,
    "auth_token": string,
    "access_token": string,
}
    
```


## 2. User Profile

Endpoint ```POST /user```

### 2.1 UserProfileRequest
```
{
    "access_token": string,
    "userid": string,
    "email": stirng,
    "mobile": string
}
```

### 2.2 UserProfileResponse
```
{
    "status": Status,
    "user": User,
}
```


## 3. Friend List

Endpoint ```POST /friends```

### 3.1 FriendListRequest
```
{
	"access_token": string,
	"query" string, 
	"limit": int,
	"skip": int,
}
```

### 3.2 FriendListResponse
```
{
    "status": Status,
    "users":  [
	    {
	    	"user": User,
	    	"lastmsg": Message
	    },
	    {
	    	"user": User,
	    	"lastmsg": Message
	    }	    
    ]
    "total": int,
}
```


## 4. Messages

### 4.1 Receive Messages

Endpoint ```POST /recv```

#### 4.1.1 ReceiveMessagesRequest
```
{
    "access_token": string,
    "userid": string,
}
```

#### 4.1.2 ReceiveMessagesResponse
```
{
    "status": Status,
    "user": User,
    "messages": []Message,
}
```


### 4.2 Send New Message

Endpoint ```POST /send```

#### 4.2.1 SendMessageRequest
```
{
	"access_token": string,
	"message": Message
}
```

#### 4.2.2 SendMessageResponse
```
{
    "status": Status,
}
```

## 5. Get User List

Endpoint ```POST /users```

### 5.1 UserListRequest
```
{
	"access_token": string,
	"query" string, 
	"limit": int,
	"skip": int,
}
```

### 5.2 UserListResponse
```
{
    "status": Status,
    "users": []User,
    "total": int,
}
```


## 6. Add Friend

Endpoint ```POST /add-friend```

### 6.1 AddFriendRequest
```
{
	"access_token": string,
	"userid": string
}
```

### 6.1 AddFriendResponse
```
{
    "status": Status,
}
```

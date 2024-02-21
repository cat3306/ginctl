syntax = "v1"

type Request {
  Name string `path:"name,options=you|me"`
}

type Response {
  Message string `json:"message"`
}

type UserCreate {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	IdCard string `json:"id_card"`
}
type UserCreateRsp {
	UserId string `json:"user_id"`
}

service {{.name}}-api {
  @handler {{.handler}}Handler
  get /from/:name(Request) returns (Response)

  @handler UserCreate
	post /api/user/create(UserCreate)returns (UserCreateRsp)
}




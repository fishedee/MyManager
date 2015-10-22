package user;

type TypeData struct{
	ADMIN int
	USER int
}

var Type = &TypeData{ADMIN:1,USER:2};

func (this *TypeData)Names()(map[string]string){
	return map[string]string{
		"1":"ADMIN",
		"2":"USER",
	};
}

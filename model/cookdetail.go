package model

type Url string

type User struct {
	NickName    string
	HomepageUrl Url
	Avatar      Url
}

type Material struct {
	Material string
	Num      string
}

type Step struct {
	PicUrl  Url
	Operate string
}

type CookDetail struct {
	UserInfo     User
	StepList     []Step
	FoodMaterial []Material
}

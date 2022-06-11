package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"` //id是视频的主键
	Author        User   `json:"author" gorm:"foreignKey:UserId"`
	UserId        int64  `gorm:"not null"` //上传视频的用户的id，好取
	PlayUrl       string `json:"play_url" json:"play_url,omitempty" gorm:"not null"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"not null"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	//喜欢的总数量
	CommentCount int64 `json:"comment_count,omitempty"`
	//评论的总数量
	IsFavorite bool `json:"is_favorite,omitempty"`
	//是否被当前用户喜欢
	Title string `json:"title,omitempty"` // 视频名称
}

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey"`
	User       User   `json:"user" gorm:"foreignKey:UserId"`
	UserId     int64  `gorm:"not null"` // 评论对应的用户 Id
	Video      Video  `gorm:"foreignKey:VideoId"`
	VideoId    int64  `gorm:"not null"` // 评论对应的视频 Id
	Content    string `json:"content,omitempty" gorm:"not null"`
	CreateDate string `json:"create_date,omitempty" gorm:"not null"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name          string `json:"name,omitempty" gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

//喜欢列表

type Favorite struct {
	Id      int64 `gorm:"primaryKey"`
	User    User  `gorm:"foreignKey:UserId"`
	UserId  int64 `gorm:"not null"`
	Video   Video `gorm:"foreignKey:VideoId"`
	VideoId int64 `gorm:"not null"`
}

type Relation struct {
	Id      int64 `gorm:"primaryKey"`
	UserA   User  `gorm:"foreignKey:UserAId"`
	UserAId int64 `gorm:"notnull"`
	UserB   User  `gorm:"foreignKey:UserBId"`
	UserBId int64 `gorm:"notnull"`
}

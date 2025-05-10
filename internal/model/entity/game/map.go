package game

type Room struct {
	Id      int64
	Map     Map
	Players []User
}

type Map struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`
	MaxPlayer    int64
	FileName     string `db:"fileName"`
	MonsterNum   string `db:"monsterNum"`
	MonsterFrame string `db:"monsterFrame"`
}

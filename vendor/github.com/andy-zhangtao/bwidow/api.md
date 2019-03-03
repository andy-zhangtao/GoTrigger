# bwidow
--
    import "github.com/andy-zhangtao/bwidow"


## Usage

```go
const (
	BW_MONGO_ENDPOINT = "BW_ENV_MONGO_ENDPOINT"
	BW_MONGO_USER     = "BW_ENV_MONGO_USER"
	BW_MONGO_PASSWD   = "BW_ENV_MONGO_PASSWD"
	BW_MONGO_DB       = "BW_ENV_MONGO_DB"
)
```

```go
const (
	BW_VERSION = "0.1.0-Alpha"
)
```

```go
const (
	//DRIVER_MONGO Mongo驱动
	DRIVER_MONGO = iota
)
```

#### type BW

```go
type BW struct {
}
```


```go
var Widow *BW
```
##### Example

```go

    //BWidow Init
    bw := bwidow.GetWidow()
    err := bw.Driver(bwidow.DRIVER_MONGO)
    if err != nil {
    	logrus.WithFields(logrus.Fields{"Init Error": err}).Errorln(ModuleName)
    	return
    }
    //Setting Model And Table
    bw.Map(User{}, "devex_user_copy")
    u := User{
    	CurrentAuthority: "BW",
    }

    //Delete All Matched Record
    num, err := bw.Delete(&u, []string{"currentauthority"})
    if err != nil {
    	logrus.Errorln(err)
    	return
    }

    logrus.WithFields(logrus.Fields{"change": num}).Info(ModuleName)

```

#### func  GetWidow

```go
func GetWidow() *BW
```
GetWidow 获取当前全局Widow. 如果没有则创建

#### func (*BW) CheckIndex

```go
func (this *BW) CheckIndex(uPtr interface{}) *BW
```

#### func (*BW) Delete

```go
func (this *BW) Delete(uPtr interface{}, field []string) (num int, err error)
```
Delete 删除命中的所有数据 uPtr 供定位记录的数据 field 用于筛选的字段

##### Example

```go

    u := User{
    	CurrentAuthority: "BW",
    }

    //Delete All Matched Record
    num, err := bw.Delete(&u, []string{"currentauthority"})
    if err != nil {
    	logrus.Errorln(err)
    	return
    }

```

#### func (*BW) DeleteAll

```go
func (this *BW) DeleteAll(uPtr interface{}) (num int, err error)
```

#### func (*BW) Driver

```go
func (this *BW) Driver(driver int) *BW
```
Driver 设置使用的数据库类型 当前支持的类型为: DRIVER_MONGO - Mongo

##### Example

```go

    	err := bw.Driver(bwidow.DRIVER_MONGO)
        if err != nil {
        	logrus.WithFields(logrus.Fields{"Init Error": err}).Errorln(ModuleName)
        	return
        }

```

#### func (*BW) Error

```go
func (this *BW) Error() error
```
Error 返回当前Error信息

#### func (*BW) FindAll

```go
func (this *BW) FindAll(u interface{}, aPtr interface{}) (err error)
```
FindAll 通过u的字段查询所有数据 BW会解析u的字段,然后将所有非空字段作为查询条件进行查询，同时将查询到的数据赋值给a u 查找记录使用的数据
a必须为array/slice类型的指针

##### Example

``` go

    u := User{
    	CurrentAuthority: "dev",
    }

    var allUser []User
    err = bw.FindAll(u, &allUser)
    if err != nil {
    	logrus.Error(err)
    	return
    }

    logrus.WithFields(logrus.Fields{"FindAll": allUser}).Info(ModuleName)

```

#### func (*BW) FindAllWithSort

```go
func (this *BW) FindAllWithSort(uPtr interface{}, aArray interface{}, sortField []string) (err error)
```
FindAllWithSort 通过u的字段查询所有数据并且按照给定的条件进行排序
BW会解析u的字段,然后将所有非空字段作为查询条件进行查询，同时将查询到的数据赋值给a u 查找记录使用的数据 a 必须为array/slice类型的指针
sortField 需要排序的字段数组

##### Example

```go

    u := User{
    	CurrentAuthority: "dev",
    }
    var allUser []User
    err = bw.FindAllWithSort(u, &allUser, []string{"+name"})
    if err != nil {
    	logrus.Error(err)
    	return
    }

    logrus.WithFields(logrus.Fields{"FindAllWithSort": allUser}).Info(ModuleName)

```

#### func (*BW) FindOne

```go
func (this *BW) FindOne(uPtr interface{}) (err error)
```
FindOne 通过u的字段查询数据 BW会解析u的字段,然后将所有非空字段作为查询条件进行查询，同时将查询到的数据赋值给u u必须为指针

##### Example

```go

    u := User{
    	Name: "ztao8607@gmail.com",
    }
    logrus.WithFields(logrus.Fields{"Query": u}).Info(ModuleName)

    err = bw.FindOne(&p)
    if err != nil {
    	logrus.Error(err)
    	return
    }

```

#### func (*BW) First

```go
func (this *BW) First(uPtr interface{}) (err error)
```
First 查询与u绑定的表中的首条记录 u 数据结构体指针

#### Example

```go

    u := User{}
    err = bw.First(&u)
    if err != nil {
    	logrus.Error(err)
    	return
    }

    logrus.WithFields(logrus.Fields{"First": u}).Info(ModuleName)

```

#### func (*BW) Map

```go
func (this *BW) Map(u interface{}, name string) *BW
```
Map 将u与数据表进行绑定 u 数据结构体 name 数据表名

##### Example

```go

    bw.Map(User{}, "devex_user_copy")

```

#### func (*BW) Save

```go
func (this *BW) Save(u interface{}) (err error)
```
Save 插入单条数据

##### Example

```go

    u := User{
    	ID:   bson.NewObjectId(),
    	Name: "fromBW@gmail.com",
    }

    err = bw.Save(u)
    if err != nil {
    	logrus.Errorln(err)
    	return
    }

```

#### func (*BW) SaveAll

```go
func (this *BW) SaveAll(uArray interface{}) (err error)
```
SaveAll 插入一批次的数据 支持数组内存在不同数据类型 u 为数组

##### Example

```go

    al := []User{
    	{
    		ID:   bson.NewObjectId(),
    		Name: "fromBW1@gmail.com",
    	},
    	{
    		ID:   bson.NewObjectId(),
    		Name: "fromBW2@gmail.com",
    	},
    	{
    		ID:   bson.NewObjectId(),
    		Name: "fromBW3@gmail.com",
    	},
    }

    err = bw.SaveAll(al)
    if err != nil {
    	logrus.Errorln(err)
    	return
    }

```

#### func (*BW) Update

```go
func (this *BW) Update(uPtr interface{}, field []string) (num int, err error)
```
Update 更新命中的所有数据. uPtr 供定位记录的数据 field 用于筛选的字段

##### Example

```go

    u := User{
    	Name:             "fromBW2@gmail.com",
    	CurrentAuthority: "BW",
    }

    num, err := bw.Update(&u, []string{"name"})
    if err != nil {
    	logrus.Errorln(err)
    	return
    }

    logrus.WithFields(logrus.Fields{"change": num}).Info(ModuleName)

```

#### func (*BW) Version

```go
func (this *BW) Version() string
```

#### type BWDriver

```go
type BWDriver interface {
	Check() error
	DriverInit() error
	Map(u interface{}, name string)
	// contains filtered or unexported methods
}
```


#### type BWMongo

```go
type BWMongo struct {
}
```


#### func (*BWMongo) Check

```go
func (this *BWMongo) Check() (err error)
```

#### func (*BWMongo) DriverInit

```go
func (this *BWMongo) DriverInit() (err error)
```

#### func (*BWMongo) Map

```go
func (this *BWMongo) Map(u interface{}, name string)
```

#### type BWMongoConf

```go
type BWMongoConf struct {
	Endpoint string
	DB       string
	User     string
	Password string
}
```

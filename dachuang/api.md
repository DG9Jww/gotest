

#注：

所有响应状态码都为200，不管失败还是成功

几乎所有响应格式都为

```json
{
    "status":true,
    "msg":"xxxx"
}
```

所有的失败响应status都为false



# 所有用户都能看的

## 所有文章信息

- [ ] GET  `http://localhosthttp://47.115.158.103:5777`

- [ ] ```json
  直接GET请求
  响应：,total是文章总数
  {
      "msg": {
          "articles": [
              {
                  "id": 5,
                  "title": "第一个标题",
                  "content": "油气储运工程系教工党支部书记文江波表示：作为一名基层高校教工党员，深深地为我党100年来带领中国走向",
                  "time": "2021-12-16 21:49:42",
                  "author": "李白",
                  "type": 0
              },
              {
                  "id": 6,
                  "title": "自传",
                  "content": "xxxxxxxxxxxxxxxxxxxxxxx",
                  "time": "2021-12-16 22:55:02",
                  "author": "张三",
                  "type": 0
              },
              {
                  "id": 7,
                  "title": "自传",
                  "content": "xxxxxxxxxxxxxxxxxxxxxxx",
                  "time": "2021-12-16 23:01:47",
                  "author": "李四",
                  "type": 0
              },
              {
                  "id": 11,
                  "title": "更改标题",
                  "content": "更改内容",
                  "time": "2021-12-27 20:40:38",
                  "author": "更改作者",
                  "type": 2
              },
              {
                  "id": 12,
                  "title": "这是一个标题",
                  "content": "这是内容",
                  "time": "2021-12-27 20:37:35",
                  "author": "这是作者",
                  "type": 2
              }
          ],
          "total": 5
      },
      "status": true
  }
  ```

#注： 文章type=1为新闻，2为通知，3为简介

## 所有通知信息

- [ ] GET `http://localhost/v1/oth/notices`

  ```json
  {
      "msg": {
          "articles": [
              {
                  "id": 11,
                  "title": "更改标题",
                  "content": "更改内容",
                  "time": "2021-12-27 20:40:38",
                  "author": "更改作者",
                  "type": 2
              },
              {
                  "id": 12,
                  "title": "这是一个标题",
                  "content": "这是内容",
                  "time": "2021-12-27 20:37:35",
                  "author": "这是作者",
                  "type": 2
              }
          ],
          "total": 2
      },
      "status": true
  }
  ```

## 所有简介新闻

简介 http://localhost/v1/oth/introduction

新闻 http://localhost/v1/oth/news

响应和通知一模一样，type不一样罢了





## 单篇文章信息

- [ ] GET  `http://localhost	?id=5` 

```json
请求，id为文章id
响应
{
    "msg": {
        "id": 5,
        "title": "第一个标题",
        "content": "油气储运工程系教工党支部书记文江波表示：作为一名基层高校教工党员，深深地为我党100年来带领中国走向",
        "time": "2021-12-16 21:49:42",
        "author": "李白",
        "type": 0
    },
    "status": true
}
```



## 搜索文章

- [ ] GET `http://localhost/v1/oth/search?query=xxx`

  ```json
  query后面跟查询内容
  
  响应
  {
      "msg": {
          "articles": [
              {
                  "id": 6,
                  "title": "自传",
                  "content": "xxxxxxxxxxxxxxxxxxxxxxx",
                  "time": "2021-12-16 22:55:02",
                  "author": "张三",
                  "type": 3
              },
              {
                  "id": 7,
                  "title": "自传",
                  "content": "xxxxxxxxxxxxxxxxxxxxxxx",
                  "time": "2021-12-16 23:01:47",
                  "author": "李四",
                  "type": 1
              }
          ],
          "total": 2
      },
      "status": true
  }
  ```

  

# 论坛用户相关

## 论坛用户登录



- [ ] POST`http://localhost/v1/oth/comm_login` 

  ```json
  请求:
  {
      "username":"aaa",
      "password":"aaa123"
  }
  
  响应成功
  {
      "msg": {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjYsIlVzZXJSaWdodCI6MywiZXhwIjoxNjQwNjAwNDgyLCJpc3MiOiJkYWNodWFuZyJ9.CXdyQBTQY56nU40IaHyUiqyUemYS5xtnyT5_BFxTLmU"
      },
      "status": true
  }
  
  失败
  {
      "msg": "账号或密码错误",
      "status": false
  }
  ```

## 用户查看个人信息

- [ ] GET`http://localhost/v1/auth/comm_user/user` 

  ```json
  
  成功响应
  {
      "msg": {
          "id": 6,
          "username": "aaa",
          "nickname": "aaa",
          "photo": ""
      },
      "status": true
  }
  ```

  

## 用户更改资料

- [ ] PUT`http://localhost/v1/auth/comm_user/user` 

```json
请求
用户只能修改昵称和头像，这里为表单请求
参数为nickname和img,头像后缀只能问jpg,jpeg,png

成功响应
{
    "msg": "更新成功",
    "status": true
}
```

## 用户修改密码

- [ ] PUT `http://localhost/v1/auth/comm_user/passwd` 

  ```json
  请求
  {   
      "old":"aaa123",
      "new":"aaa12345"
  }
  
  成功响应
  {
      "msg": "成功修改密码",
      "status": true
  }
  失败status则为false
  ```



## 用户添加话题

- [ ] POST `http://localhost/v1/auth/comm_user/topic` 

  ```json
  请求
  {   
      "question":"请问怎么安装vscode?"
  }
  
  成功响应
  {
      "msg": "添加成功",
      "status": true
  }
  ```





## 新增回答

- [ ] POST `http://localhost/v1/auth/comm_user/answer` 

  ```json
  请求
  需要当前话题的QID,和回答内容即可
  {   
      "qid":1,
      "answer":"直接百度"
  }
  
  成功响应
  {
      "msg": "回答成功",
      "status": true
  }
  ```



##  **查看所有话题**

- [ ] GET `http://localhost/v1/auth/comm_user/topics`

  ```json
  直接请求即可
  成功响应
  
  {
      "msg": {
          "questions": [
              {
                  "qid": 1,
                  "question": "请问怎么安装vscode?",
                  "uid": 6,
                  "nickname": "",
                  "photo": "/static/photo/1640598292.jpg",
                  "time": "2021-12-27 17:52:58"
              }
          ],
          "total": 1
      },
      "status": true
  }
  ```

## 查看自己发布的话题

- [ ] GET `http://localhost/v1/auth/comm_user/my_topics` 

  ```json
  直接请求
  响应
  {
      "msg": {
          "questions": [
              {
                  "qid": 1,
                  "question": "请问怎么安装vscode?",
                  "uid": 6,
                  "time": "2021-12-27 17:52:58"
              }
          ]
      },
      "status": true
  }
  ```

## 查看回答

- [ ] GET `http://localhost/v1/auth/comm_user/topic?qid=1` **点进去查看该话题下的回答**

  ```json
  请求通过qid控制
  响应内容包含 回答信息和回答者的头像和昵称
  {
      "msg": [
          {
              "aid": 1,
              "qid": 1,
              "uid": 6,
              "answer": "直接百度",
              "time": "2021-12-27 18:03:56",
              "nickname": "",
              "photo": "/static/photo/1640598292.jpg"
          }
      ],
      "status": true
  }
  ```

## 删除话题

- [ ] DELETE `http://localhost/v1/auth/comm_user/topic` **用户删除自己发布的话题，一次只能删一个**

  ```json
  请求
  {   
      "qid":1
  }
  响应
  {
      "msg": "删除成功",
      "status": true
  }
  ```

## 删除答案

- [ ] DELETE `http://localhost/v1/auth/comm_user/answer` **用户删除自己发布话题下的回答**

  ```json
  请求
  qid为话题id,aid为回答id
  {   
      "qid":2,
      "aid":3
  }
  
  响应
  {
      "msg": "删除成功",
      "status": true
  }
  ```


## 搜索

- [ ] GET `http://localhost/v1/auth/comm_user/search?query=xxx`

```json
请求如上

响应
{
    "msg": {
        "questions": [
            {
                "qid": 3,
                "question": "hello",
                "uid": 6,
                "time": "2021-12-27 21:13:13"
            }
        ],
        "total": 1
    },
    "status": true
}

```



# 管理员相关

## 登录

- [ ] POST `http://localhost/v1/oth/admin_login` 管理员登录

  ```json
  请求
  {   
      "username":"tom",
      "password":"tom12"
  }
  登录成功
  {
      "msg": {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjUsIlVzZXJSaWdodCI6MCwiZXhwIjoxNjQwNjExMDcyLCJpc3MiOiJkYWNodWFuZyJ9.4ijYbencV3gc7BCRr6d-FXNsNZQ7hZJtl3hZ0i_M-vU"
      },
      "status": true
  }
  
  失败
  
  {
      "msg": "账号或者密码错误",
      "status": false
  }
  
  ```

## 单个文章信息

- [ ] GET `http://localhost/v1/auth/user/article?id=5` 查看单个文章所信息

  ```json
  直接GET请求，id控制文章id
  响应
  {
      "msg": {
          "id": 5,
          "title": "第一个标题",
          "content": "油气储运工程系教工党支部书记文江波表示：作为一名基层高校教工党员，深深地为我党100年来带领中国走向",
          "time": "2021-12-16 21:49:42",
          "author": "李白",
          "img": "/static/img/1234.jpg",
          "type": "0"
      },
      "status": true
  }
  ```

## 所有文章信息

- [ ] GET `http://localhost/v1/auth/user/articles` 

```json
直接请求，无参数
响应
{
    "msg": {
        "articles": [
            {
                "id": 5,
                "title": "第一个标题",
                "content": "油气储运工程系教工党支部书记文江波表示：作为一名基层高校教工党员，深深地为我党100年来带领中国走向",
                "time": "2021-12-16 21:49:42",
                "author": "李白",
                "img": "/static/img/1234.jpg",
                "type": "0"
            },
            {
                "id": 6,
                "title": "自传",
                "content": "xxxxxxxxxxxxxxxxxxxxxxx",
                "time": "2021-12-16 22:55:02",
                "author": "张三",
                "img": "/haha/2.jpg",
                "type": "0"
            },
            {
                "id": 7,
                "title": "自传",
                "content": "xxxxxxxxxxxxxxxxxxxxxxx",
                "time": "2021-12-16 23:01:47",
                "author": "李四",
                "img": "/haha/2.jpg",
                "type": "0"
            },
            {
                "id": 4,
                "title": "哈哈",
                "content": "xxxxxxxx哈哈哈哈哈哈哈哈哈哈哈哈哈xxxxxxxxxxxxxxx",
                "time": "2021-12-17 15:23:07",
                "author": "王五",
                "img": "/haha/2.jpg",
                "type": "0"
            },
            {
                "id": 8,
                "title": "aa1",
                "content": "aaaa",
                "time": "2021-12-18 22:46:26",
                "author": "wang",
                "img": "",
                "type": "0"
            }
        ],
        "total": 5
    },
    "status": true
}
```

## 添加文章

- [ ] POST `http://localhost/v1/auth/user/article` 

- [ ] ```json
  
  请求
  {   
      "title":"这是一个标题",
      "content":"这是内容",
      "author":"这是作者",
      "type":2
  }
  响应
  {
      "msg": "添加成功",
      "status": true
  }
  
  ```

##  更新文章

- [ ] PUT `http://localhost/v1/auth/user/article`

  ```json
  请求
  文章id和文章内容
  {   
      "id":11,
      "title":"更改标题",
      "content":"更改内容",
      "author":"更改作者"
  }
  响应
  {   
      "id":11,
      "title":"更改标题",
      "content":"更改内容",
      "author":"更改作者"
  }
  ```

## 删除文章

- [ ] DELETE `http://localhost/v1/auth/user/article`

```json
请求，看清楚不是字符串
{   
    "id":[1,2,3]
}
响应
{
    "msg": "删除成功!",
    "status": true
}
```



## 个人信息

- [ ] GET `http://localhost/v1/auth/user/user`

```json
直接请求即可
响应
{
    "msg": {
        "id": 5,
        "username": "tom",
        "nickname": "",
        "email": "",
        "power": 0
    },
    "status": true
}
```

## 修改信息

- [ ] PUT `http://localhost/v1/auth/user/user` 	

```json
请求，这里只能修改昵称和email
{   
    "nickname":"hello",
    "email":"hahahhaha"
}

响应
{
    "msg": "更新成功",
    "status": true
}

```



## 修改密码

- [ ] PUT `http://localhost/v1/auth/user/passwd`

```json
请求
{   
    "old":"tom124",
    "new":"tom123"
}

响应
{
    "msg": "成功修改密码",
    "status": true
}
```



## 查看所有社区用户

- [ ] GET `http://localhost/v1/auth/user/comm_users` 

```json
直接请求即可

响应
{
    "msg": {
        "total": 3,
        "users": [
            {
                "id": 6,
                "username": "aaa",
                "nickname": "",
                "photo": "/static/photo/1640598292.jpg"
            },
            {
                "id": 7,
                "username": "bbb",
                "nickname": "iambbb",
                "photo": ""
            },
            {
                "id": 8,
                "username": "ccc",
                "nickname": "",
                "photo": "/static/photo/1640536030.jpg"
            }
        ]
    },
    "status": true
}
```

## 添加用户

- [ ] POST `http://localhost/v1/auth/user/comm_user`

```json
请求
{   
    "username":"nihao",
    "password":"nihao123",
    "nickname":"nnn"
}
响应
{
    "msg": "添加成功",
    "status": true
}
```

## 删除用户

- [ ] DELETE `http://localhost/v1/auth/user/comm_user`

```json
请求,注意是数组
{   
    "id":[7,8]
}

响应
{
    "msg": "删除成功!",
    "status": true
}
```



## 查看所有话题

- [ ] GET `http://localhost/v1/auth/user/topics`

```json
直接请求即可
响应，这里和论坛用户得到的数据是一样的，total是总数
{
    "msg": {
        "questions": [
            {
                "qid": 2,
                "question": "请问怎么安装vscode?",
                "uid": 6,
                "nickname": "",
                "photo": "/static/photo/1640598292.jpg",
                "time": "2021-12-27 20:02:26"
            }
        ],
        "total": 1
    },
    "status": true
}
```

## 删除话题

- [ ] DELETE `http://localhost/v1/auth/user/topic`

```json
请求，注意是列表，因为可以同时多个删除
{   
    "id":[3,4]
}
响应
{
    "msg": "删除成功",
    "status": true
}
```

## 查看回答

点击一个话题，然后显示它的回答

- [ ] GET `http://localhost/v1/auth/user/topic?qid=2`

```json
这里和普通用户一样的
请求，id控制参数

{
    "msg": [
        {
            "aid": 4,
            "qid": 2,
            "uid": 6,
            "answer": "直接百度2",
            "time": "2021-12-27 20:03:47",
            "nickname": "",
            "photo": "/static/photo/1640598292.jpg"
        }
    ],
    "status": true
}
```

##   删除回答

- [ ] DELETE `http://localhost/v1/auth/user/answer`

```json
请求，数组（管理员才可以批量删）
    {   
        "id":[5,6]
    }

响应
{
    "msg": "删除成功",
    "status": true
}
```



# 超级管理员

## 查看普通管理员

- [ ] GET `http://localhost/v1/auth/admin/admins`

```json
直接请求即可
响应
{
    "msg": {
        "total": 5,
        "users": [
            {
                "id": 1,
                "username": "jack",
                "nickname": "jack",
                "email": "a@qq.com",
                "power": 0
            },
            {
                "id": 6,
                "username": "marry",
                "nickname": "aimarry",
                "email": "asd@qq.com",
                "power": 0
            },
            {
                "id": 5,
                "username": "tom",
                "nickname": "hello",
                "email": "hahahhaha",
                "power": 0
            },
            {
                "id": 4,
                "username": "aaa",
                "nickname": "hhasdh",
                "email": "a@asdsd.com",
                "power": 0
            },
            {
                "id": 7,
                "username": "admin",
                "nickname": "admin",
                "email": "admin@qq.com",
                "power": 1
            }
        ]
    },
    "status": true
}
```

## 添加管理员

- [ ] POST `http://localhost/v1/auth/admin/admin`

```json
请求
{   
    "username":"nihao",
    "password":"nihao123"
}
响应
{
    "msg": "添加成功",
    "status": true
}
```

## 删除管理员

- [ ] DELETE `http://localhost/v1/auth/admin/admin`

```json
请求，数组
{   
    "id":[7,11]
}

响应
{
    "msg": "删除成功!",
    "status": true
}
```


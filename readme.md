# 필독 브랜치 규칙
main - prod, stage - stage , dev - 개발서버

pr 순서 feat => dev => (hotfix/bug)stage => main


1. 최신 dev브랜치에서 feature 만들기
2. dev에 push전 dev pull 받기
3. bug/hotfix 를 제외한 브랜치(ex:feat)로 main/stage에 직접pr금지

### makefile

```shell
# local postgres run (docker-compose)
make local-db
# local postgres migrate init
make local-init
# local postgres apply migrate
make local-migrate
```

# swagger 설정 [출처](https://www.soberkoder.com/swagger-go-api-swaggo/)

## dev 설정

```shell
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

## main에

```code
   import (_ "[project명]/docs")
```

```shell
# swagger json 생성   swag init -g [project main path].go
swag init -g cmd/app/main.go
```

## [스웨거 링크](http://localhost:8082/swagger/index.html)


## Cafe입니다.

1. 게시글을 작성할수있는 카페를 생성할수있습니다.
2. 카페 생성시 게시글(boardType) 종류와 권한(role)에 따른 CRUD(액션)를 지정할수있습니다.
3. 게시글 종류(boardType)생성전 Role을 생성해야합니다(ex:manger,user,geust?등등)
4. bardType(ex:notice(공지) , free(자유) 등등)별로 action과 action에 따른 role을 부여해야합니다.


## entity 구조

```text
    Cafe{ // cafe 가 만들어 질때 멤버 등록을 해주어야함
        id 
        name 
        owner_id(카페 생성자_id)
        description
        created_at
    }
        Member{ //  1인 같은카페 1번만 가입 가능 , 카페당 유니크한 닉네임을 가짐 
            (uniq key:  cafe_id + nickname, user_id + cafe_id)
            id
            cafe_id
            user_id 
            nickname
        }   
        CafeRole{ // cafe owner 만  수정,삭제 ,추가 가능 
                  // 카페당 같은 이름의 룰은 1개만 가능
            (uniq key: cafe_id+role_name)
            id
            cafe_id
            name
            description
        }
            MemberRole{ // 한 유저 , 한카페 다중 권한을 가질수있습니다. 실제 멤버 cafeRole 이 저장되있는 entity  
                (uniq key: cafe_role_id, member_id)
                id 
                member_id 
                cafe_id // 해당 카페 권한 확인용
                cafe_role_ids // int arr (string 타입 ) 
            }
        BoardType{ 
            (uniq key: cafe_id+name)
            id
            cafe_id
            name 
            description
            []roles
        }
            BoardAction{
                id
                cafe_id
                board_type_id
                create_roles: []roles varchar // 생성권한
                read_roles: []roles varchar  // 매니저는 읽기 권한이 기본적으로 추가됩니다.
                update_roles: []roles varchar  // 업데이트 권한 없으면 수정불가능한 게시글이됩니다.
                update_able:  bool // 업데이트가 가능한지 불가능시 roles 에 관계없이 수정이 불가능합니다 
                delete_roles: []roles varchar  // 매니저는 무조건 삭제권한을 얻습니다(관리 차원)
            }
        ban{    // 한 유저 + 카페 당 벤은 1번만 가능 
            (uniq key: user_id + cafe_id
            member_id
            
        }
```


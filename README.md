## Todo list app 

### Brief description

​		This is a todo list app that have some features as follow:

- Allow user to register
- Allow user to login
- Allow user to logout
- Allow user get their profile
- Allow user to make their own workflow of todo list
- Allow user to get their todo list
- Allow user to get their detail todo
- Allow user to add new todo
- Allow user to remove todo from the list
- Allow user to update their todo
- Allow user to move their todo to the next workflow

### Database schema

​	This app is formed from three major entiries and the data was done using redis. Below are the three major entities along with its brief description.

1. User entity  

   ![](/Users/moris/Desktop/IceHouse/todo_list_exercise/user entity.png)

   ​	As shown above, the user entitiy consists of username, password, name, email, picture, login_status as well as the todo_list.

   Table below will explain each of the field

   | Field name   | Field type     | Description                                                  | Value example                                                |
   | ------------ | -------------- | ------------------------------------------------------------ | :----------------------------------------------------------- |
   | Username     | String         | This field is the primary key for the entity. It holds unique value which can be used to distinguish users | Ex. Moris123, helloWorld.                                    |
   | Password     | String         | This field holds the user password which will be used when user login to the application | Ex. secRetPassworD, iBetYouDontKnowMyPassword                |
   | Name         | String         | This field holds the user name and this field is not unique. | Ex. Moris, raja                                              |
   | Email        | String         | This field holds the user email and this field must be unique so there won't be two users with the same email | Ex. moris@gmail.com, raja123@hotmail.com                     |
   | Picture      | String url     | This field holds the user picture.                           | Ex. /Users/moris/Desktop/IceHouse/todo_list_exercise/user entity.png, /Users/moris/Desktop/IceHouse/todo_list_exercise/todo entity.png |
   | Login_status | Boolean        | This field holds the user login status to help the app to determine if the user has logged in or not. The initial value for this field is false | Ex. True (when the user has logged in ), false(when the user has logged out). |
   | Todo_list    | Radis hash map | This field holds the todo list of user. This todo list can have more than one todo inside the list. This field only stores the todo_id from the todo entity as the key and the todo name as the value | Ex. {101, learn golang, 102, swim, 103, learn redis}         |

   

    

2. Todo entity

   ![](/Users/moris/Desktop/IceHouse/todo_list_exercise/todo entity.png)



3. Workflow entity

   ![](/Users/moris/Desktop/IceHouse/todo_list_exercise/workflow entity.png)





|      |      |      |
| ---- | ---- | ---- |
|      |      |      |
|      |      |      |
|      |      |      |


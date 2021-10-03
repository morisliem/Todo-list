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

   ![](img/user%20entity.png)

   ​	As shown above, the entity was done in redis hash, username will be the key and the subvalue consists of username, password, name, email, picture, login_status as well as the todo_list.

   Table below will explain each of the field

   | Field name   | Field type | Description                                                  | Value example                                                |
   | ------------ | ---------- | ------------------------------------------------------------ | :----------------------------------------------------------- |
   | Username     | String     | This field is the primary key for the entity. It holds unique value which can be used to distinguish users | Ex. Moris123, helloWorld.                                    |
   | Password     | String     | This field holds the user password which will be used when user login to the application | Ex. secRetPassworD, iBetYouDontKnowMyPassword                |
   | Name         | String     | This field holds the user name and this field is not unique. | Ex. Moris, raja                                              |
   | Email        | String     | This field holds the user email and this field must be unique so there won't be two users with the same email | Ex. moris@gmail.com, raja123@hotmail.com                     |
   | Picture      | String url | This field holds the user picture.                           | Ex. /Users/moris/Desktop/IceHouse/todo_list_exercise/user entity.png, /Users/moris/Desktop/IceHouse/todo_list_exercise/todo entity.png |
   | Login_status | Boolean    | This field holds the user login status to help the app to determine if the user has logged in or not. The initial value for this field is false | Ex. True (when the user has logged in ), false(when the user has logged out). |
   | Todo_list    | Radis hash | This field holds the todo list of user. This todo list can have more than one todo inside the list. This field only stores the todo_id from the todo entity as the key and the todo title as the value | Ex. {101, learn golang, 102, swim, 103, learn redis}         |

2. Todo entity

   ![](img/todo%20entity.png)
   
   ​	As shown above, the entity was done in redis hash, todo_id will be the key and the subvalue consists of todo_id, title, description, label, deadline, severity, priority as well as the workflow_state.
   
   Table below will explain each of the field
   
   | Field name     | Field type | Description                                                  | Value example                                      |
   | -------------- | ---------- | ------------------------------------------------------------ | -------------------------------------------------- |
   | Todo_id        | String     | This field is the primary key for the entity. It holds unique value which can be used to distinguish the todos | Ex. 101, 102, 103, 104                             |
   | Title          | String     | This field holds todo title.                                 | Ex. Running, pick up laundry, finish todo list app |
   | Description    | String     | This field holds todo description. This description is used to clarify the todo | Ex. Pick up pants and shirts laundry at xxx place  |
   | Label          | String     | This fields is used to categorize the todo                   | Ex. Study, Exercise,                               |
   | Deadline       | Date       | This field shows when is the latest time the user can finish their todo | Ex. 3-10-2021, 10-10-2021                          |
   | Severity       |            |                                                              |                                                    |
   | Priority       | Enum       | This field shows how important the todo is. It helps the user to prioritize which todo they should do first | Ex. Low, medium, high                              |
   | Workflow_state | String     | This fields is used to specifiy in which state the todo is in and it  can be moved to another state | Ex. Initiate, Ready, Finish                        |
   

3. Workflow entity

   ![](img/workflow%20entity.png)
   
   ​	As shown above, the entity was done in redis set, workflow be the key and the value could be added as much as the user wants. In my image above, three workflow values were specified.
   
   Table below will explain each of the field
   
   
   
   | Field name | Field type | Description                                                  | Value example |
   | ---------- | ---------- | ------------------------------------------------------------ | ------------- |
   | Workflow   | String     | This field will be the key for the redis set.                | Ex. Workflow  |
   | Workflow1  | String     | This field will be the value of the redis set key. It holds the state of the workflow. | Ex. Initiate  |
   | Workflow2  | String     | This field will be the value of the redis set key. It holds the state of the workflow. | Ex. Running   |
   | Workflow3  | String     | This field will be the value of the redis set key. It holds the state of the workflow. | Ex. Finish    |
   | . . .      | String     | This field is used to show that it's possible to add more workflow to the redis set. | Ex. Review    |
   
   



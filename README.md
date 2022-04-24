# Forum

* **Launch the project with:**
```bash 
go run .\main\main.go
```
* **Open your browser and go to:**
```url
http://localhost:8080/
```

## Availables features
### <ins>Basic Forum:</ins>
- ### Authentication:
You can register / login on the forum by inputting credentials.
This allowed you to create your own post and leave comments, like, dislike on posts and use more specific filters.
Even if you are not connected, you can see posts, like, dislike, comments but not created them.
- ### Creation of Posts and Comments:
If you are logged, you can create your post with the following options:
*Title:* Choose a good title (Required field).
*Content:* Write the content of the post (Required field).
*Select Category:* Choose the category of your post.
*Picture:* You can add a Picture (See [Upload Image Section](#Image-Upload)).

For comment a post, just go in view comments section button at the bottom of the post and click on the button *create comment*
- ### Like & Dislike:
If you are logged, you can like and dislike posts and comments with :+1: & :-1:.
- ### Filters:
Available at the top of the main webpage, you can choose between 3 different filters:
- By category
- By created posts (must be logged in)
- By liked posts (must be logged in)
### <ins>Bonuses:</ins>
- ### Image Upload
Options available at post creation, it allows you to add an image at your post. Only JPEG, PNG and GIF are admitted with a size below 20 MB.
- ### Advanced features
Options available at the top of the main webpage if you are logged, you can go to your activity page for looks:
- Your created posts
- Your liked and disliked pots
- The post where your left comments

You can navigate easily with shortcut at the top of the page.
For your created posts and comments, you can remove them and edit them with buttons after them.
Editing of posts and comments is the same as their creation.

### <ins>Others:</ins>

For security, the password of users are hashed and we use UUID for users.

* **Equip:**
This Forum is made by:
[PORTE Brandon](https://git.ytrack.learn.ynov.com/BPORTE1)
[BRAVO Valentin](https://git.ytrack.learn.ynov.com/VBRAVO)
[MARTINEZ BIENABE Alienor](https://git.ytrack.learn.ynov.com/AMARTINEZBIENABE)
[BOURMAUD Hugo](https://git.ytrack.learn.ynov.com/HBOURMAUD)
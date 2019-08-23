package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/lpose/twitterGo/src/domain"
	"github.com/lpose/twitterGo/src/service"
)

func main() {
	tweetManager := service.NewTweetManager()
	userManager := service.GetInstance()
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Registrar usuario",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Ingresa tu id: ")
			idString := c.ReadLine()
			id, _ := strconv.Atoi(idString)
			c.Print("Usuario: ")
			name := c.ReadLine()
			c.Print("Contraseña: ")
			pass := c.ReadLine()
			c.Print("Email: ")
			mail := c.ReadLine()
			c.Print("Nick: ")
			nick := c.ReadLine()

			var registerUser *domain.User
			registerUser = domain.NewUser(id, name, mail, pass, nick)
			userManager.Register(registerUser)

			c.Print("Usuario registrado\n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Ingresa tu usuario: ")
			userName := c.ReadLine()
			_, user := userManager.GetUserLogiado(userName)
			var id int
			if user == nil {
				c.Print("El usuario no esta logiado - Utilice 'login' para iniciar sesion\n")
				return
			}
			c.Print("Write your tweet: ")
			text := c.ReadLine()

			var tweet *domain.TextTweet
			tweet = domain.NewTweet(id, user, text)
			id, err := tweetManager.PublishTweet(tweet)
			if err == nil {
				c.Print("Tweet sent\n")
			} else {
				c.Print(err)
			}

			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetByUser",
		Help: "Show all user's tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Ingrese el usuario: ")
			userName := c.ReadLine()
			user := userManager.GetUserByNick(userName)
			if user == nil {
				c.Print("No existe el usuario: ", userName, "\n")
				return
			}
			tweets := tweetManager.GetTweetsByUser(user)
			c.Println("Usuario: ", user.Name)
			for i, tweet := range tweets {
				c.Print("--------------------")
				c.Print("Tweet ", i)
				c.Print("--------------------\n")
				c.Println("Tweet: ", tweet.Text)
				c.Print("Hora: ", tweet.Date, "\n")

			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Login user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Ingresa tu nick: ")
			userNick := c.ReadLine()
			c.Print("Ingresa tu contraseña: ")
			pass := c.ReadPassword()

			err := userManager.Login(userNick, pass)
			if err != nil {
				c.Print(err)
				return
			}
			c.Print("Usuario logiado correctamente\n")
			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "logout user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Ingresa tu nick: ")
			userNick := c.ReadLine()
			c.Print("Ingresa tu contraseña: ")
			pass := c.ReadPassword()

			err := userManager.Logout(userNick, pass)
			if err != nil {
				c.Print(err)
				return
			}
			c.Print("Usuario desloogiado correctamente\n")
			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()
			if tweets == nil {
				c.Println("No hay tweets")
				return
			}

			for i, tweet := range tweets {
				c.Print("--------------------")
				c.Print("Tweet ", i)
				c.Print("--------------------\n")
				c.Println("El usuario: ", tweet.User.Name)
				c.Println("Tweet: ", tweet.Text)
				c.Print("Hora: ", tweet.Date, "\n")

			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUserLogin",
		Help: "Shows a users's login",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			users := userManager.GetUsersLogin()
			if len(users) == 0 {
				c.Println("No hay usuarios logiados")
				return
			}

			for _, user := range users {
				c.Print(user.Name, "\n")
			}

			return
		},
	})

	shell.Run()

}

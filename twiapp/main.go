package main

import (
	"fmt"
	"sync"

	"github.com/beatrichartz/martini-sockets"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/oauth2"
	goauth2 "golang.org/x/oauth2"

	"github.com/saromanov/twiapp/structs"
)


func Twitter(conf *goauth2.Config) martini.Handler {
	conf.Endpoint = goauth2.Endpoint{
		AuthURL:  "https://api.twitter.com/oauth/authorize",
		TokenURL: "https://api.twitter.com/oauth2/token",
	}
	return oauth2.NewOAuth2Provider(conf)
}

type Client struct {
	Name string
	in   <-chan *structs.Tweet
	out  chan<- *structs.Tweet

	done       <-chan bool
	err        <-chan error
	disconnect chan<- int
}

// Room level
type Room struct {
	sync.Mutex
	name               string
	clients            []*Client
}

// Add a client to a room
func (r *Room) appendClient(client *Client) {
	r.Lock()
	r.clients = append(r.clients, client)
	for _, c := range r.clients {
		//if c != client {
		c.out <- &structs.Tweet{"new", client.Name}
		//}
	}
	r.Unlock()
}

func main() {
	room := &Room{sync.Mutex{}, "test1", make([]*Client, 0)}
	m := martini.Classic()
	// Use Renderer
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123"))))
	m.Use(Twitter(
		&goauth2.Config{
			ClientID:     "962693048",
			ClientSecret: "04w26V2zLclDld96od93lACF8UxiZXxxP8BMM1EI8FAVmamrEc",
			Scopes:       []string{"https://api.twitter.com/oauth/authorize"},
			RedirectURL:  "redirect_url",
		},
	))

	// Autorization page
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})

	// Autorization page
	m.Get("/authorize", func(r render.Render, params martini.Params) {
		fmt.Println(r, params)
	})

	m.Get("/dashboard", func(r render.Render){
		r.HTML(200, "dashboard", "")
	})

	m.Get("/sockets/dashboard/:id", sockets.JSON(structs.Tweet{}), func(r render.Render, params martini.Params, receiver <-chan *structs.Tweet, sender chan<- *structs.Tweet, done <-chan bool, disconnect chan<- int, err <-chan error) (int, string) {
		client := &Client{params["id"], receiver, sender, done, err, disconnect}
		// A single select can be used to do all the messaging
		room.appendClient(client)
		for {
			select {
			case <-client.err:
				// Don't try to do this:
				// client.out <- &Message{"system", "system", "There has been an error with your connection"}
				// The socket connection is already long gone.
				// Use the error for statistics etc
			case msg := <-client.in:
				fmt.Println("New message: ", msg)
			case <-client.done:
				//room.removeClient(client)
				return 200, "OK"
			}
		}
	})
	m.Run()

}

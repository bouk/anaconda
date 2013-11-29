package anaconda_test

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"time"
)

// Initialize an client library for a given user.
// This only needs to be done *once* per user
func ExampleTwitterApi_InitializeClient() {
	anaconda.SetConsumerKey("your-consumer-key")
	anaconda.SetConsumerSecret("your-consumer-secret")
	api := anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
	fmt.Println(*api.Credentials)
}

func ExampleTwitterApi_GetSearch() {

	anaconda.SetConsumerKey("your-consumer-key")
	anaconda.SetConsumerSecret("your-consumer-secret")
	api := anaconda.NewTwitterApi("your-access-token", "your-access-token-secret")
	search_result, err := api.GetSearch("golang", nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range search_result {
		fmt.Print(tweet.Text)
	}
}

// Rate-limiting can easily be handled in the background, automatically
func ExampleTwitterApi_RateLimiting() {
	api := anaconda.NewTwitterApi("your-access-token", "your-access-token-secret")
	api.EnableRateLimiting(10*time.Second, 5)

	// These queries will execute in order
	// with appropriate delays inserted only if necessary
	golangTweets, err := api.GetSearch("golang", nil)
	anacondaTweets, err2 := api.GetSearch("anaconda", nil)

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err)
	}

	fmt.Println(golangTweets)
	fmt.Println(anacondaTweets)
}

// Fetch a list of all followers without any need for managing cursors
// (Each page is automatically fetched when the previous one is read)
func ExampleTwitterApi_GetFollowersListAll() {
	pages := api.GetFollowersListAll(nil)
	for page := range pages {
		//Print the current page of followers
		fmt.Println(page.Followers)
	}
}

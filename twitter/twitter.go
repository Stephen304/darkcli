package twitter

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Stephen304/cmdfolder"
)

/*
Folder wraps another folder and adds an api field
*/
type Folder struct {
	cmdfolder.Folder

	api *anaconda.TwitterApi
}

/*
New makes a new twitter folder
*/
func New() cmdfolder.Folder {
	// Make API
	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_API_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	// Make folder wrappers and store apis
	folder := &Folder{cmdfolder.New(), api}
	dmFolder := newDMFolder(api)

	// Add folders
	folder.AddFolder("dm", dmFolder)

	return folder
}

// func (folder *Folder) getTimeline(_ string) {
// 	fmt.Println("Here goes your timeline")
// }

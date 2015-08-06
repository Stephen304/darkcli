package twitter

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Stephen304/cmdfolder"
)

/*
ThreadFolder wraps generic folder to add a messages field
*/
type ThreadFolder struct {
	cmdfolder.Folder

	api    *anaconda.TwitterApi
	thread *dmThread
}

func newThreadFolder(api *anaconda.TwitterApi, thread *dmThread) cmdfolder.Folder {
	// Create the folder
	folder := &ThreadFolder{cmdfolder.New(), api, thread}

	// Add commands
	folder.AddCommand("ls", folder.lsDM)
	folder.AddCommand("say", folder.say)

	return folder
}

func (folder *ThreadFolder) lsDM(_ string) {
	for _, message := range folder.thread.messages {
		fmt.Println("@" + message.from)
		fmt.Println(message.text)
		fmt.Println(message.time.Format("Monday Jan 2, 3:04pm") + "\n")
	}
}

func (folder *ThreadFolder) say(message string) {
	_, err := folder.api.PostDMToScreenName(message[4:], folder.thread.to)
	if err != nil {
		fmt.Println(err)
	}
}

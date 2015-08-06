package twitter

import (
	"net/url"
	"sort"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Stephen304/cmdfolder"
)

type dm struct {
	from string
	text string
	time time.Time
}

type dmThread struct {
	to       string
	messages []dm
}

type dmThreads map[string]*dmThread

type dmByDate []dm

func (v dmByDate) Len() int           { return len(v) }
func (v dmByDate) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v dmByDate) Less(i, j int) bool { return v[i].time.Before(v[j].time) }

func newDMFolder(api *anaconda.TwitterApi) cmdfolder.Folder {
	// Create the folder
	folder := &Folder{cmdfolder.New(), api}

	// Add commands
	folder.AddCommand("ls", folder.lsDM)

	return folder
}

func (folder *Folder) lsDM(command string) {
	folder.ClearFolders()
	folder.injectDM(getDM(folder.api))
	folder.Ls(command)
}

func (folder *Folder) injectDM(threads dmThreads) {
	for to, thread := range threads {
		folder.AddFolder(to, newThreadFolder(folder.api, thread))
	}
}

func getDM(api *anaconda.TwitterApi) dmThreads {
	// Get dms
	v := url.Values{}
	v.Set("count", "200")
	messages, _ := api.GetDirectMessages(v)

	// Sort dms into threads
	threads := make(dmThreads)
	timezone, _ := time.LoadLocation("Local")
	for _, message := range messages {
		t, _ := time.Parse("Mon Jan 2 15:04:05 -0700 2006", message.CreatedAt)
		if threads[message.SenderScreenName] == nil {
			threads[message.SenderScreenName] = &dmThread{to: message.SenderScreenName}
		}
		threads[message.SenderScreenName].messages = append(threads[message.SenderScreenName].messages, dm{from: message.SenderScreenName, text: message.Text, time: t.In(timezone)})
	}

	// Sort messages in threads
	for _, thread := range threads {
		sort.Sort(dmByDate(thread.messages))
	}

	return threads
}

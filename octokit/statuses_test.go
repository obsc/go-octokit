package octokit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatuses(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/jingweno/gh/statuses/740211b9c6cd8e526a7124fe2b33115602fbc637", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("statuses.json"))
	})

	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	url, err := StatusesURL.Expand(M{"owner": "jingweno", "repo": "gh", "ref": sha})
	assert.NoError(t, err)

	statuses, err := client.Statuses(url).All()

	assert.Len(t, statuses, 2)
	firstStatus := statuses[0]
	assert.Equal(t, "pending", firstStatus.State)
	assert.Equal(t, "The Travis CI build is in progress", firstStatus.Description)
	assert.Equal(t, "https://travis-ci.org/jingweno/gh/builds/11911500", firstStatus.TargetURL)
}

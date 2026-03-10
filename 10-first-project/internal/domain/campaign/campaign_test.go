package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}

	campaign, _ = NewCampaign(name, content, contacts)
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidadeName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name or content is required", err.Error())
}

func Test_NewCampaign_MustValidadeContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("name or content is required", err.Error())
}

func Test_NewCampaign_MustValidadeContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}

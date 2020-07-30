package protocol

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	gethbridge "github.com/status-im/status-go/eth-node/bridge/geth"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/protocol/encryption/multidevice"
	"github.com/status-im/status-go/protocol/tt"
	"github.com/status-im/status-go/waku"
)

const statusChatID = "status"

func TestMessengerInstallationSuite(t *testing.T) {
	suite.Run(t, new(MessengerInstallationSuite))
}

type MessengerInstallationSuite struct {
	suite.Suite
	m          *Messenger        // main instance of Messenger
	privateKey *ecdsa.PrivateKey // private key for the main instance of Messenger

	// If one wants to send messages between different instances of Messenger,
	// a single Waku service should be shared.
	shh types.Waku

	tmpFiles []*os.File // files to clean up
	logger   *zap.Logger
}

func (s *MessengerInstallationSuite) SetupTest() {
	s.logger = tt.MustCreateTestLogger()

	config := waku.DefaultConfig
	config.MinimumAcceptedPoW = 0
	shh := waku.New(&config, s.logger)
	s.shh = gethbridge.NewGethWakuWrapper(shh)
	s.Require().NoError(shh.Start(nil))

	s.m = s.newMessenger(s.shh)
	s.privateKey = s.m.identity
}

func (s *MessengerInstallationSuite) newMessengerWithKey(shh types.Waku, privateKey *ecdsa.PrivateKey) *Messenger {
	tmpFile, err := ioutil.TempFile("", "")
	s.Require().NoError(err)

	options := []Option{
		WithCustomLogger(s.logger),
		WithMessagesPersistenceEnabled(),
		WithDatabaseConfig(tmpFile.Name(), "some-key"),
		WithDatasync(),
	}
	installationID := uuid.New().String()
	m, err := NewMessenger(
		privateKey,
		&testNode{shh: shh},
		installationID,
		options...,
	)
	s.Require().NoError(err)

	err = m.Init()
	s.Require().NoError(err)

	s.tmpFiles = append(s.tmpFiles, tmpFile)

	return m
}

func (s *MessengerInstallationSuite) newMessenger(shh types.Waku) *Messenger {
	privateKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	return s.newMessengerWithKey(s.shh, privateKey)
}

func (s *MessengerInstallationSuite) TestReceiveInstallation() {
	theirMessenger := s.newMessengerWithKey(s.shh, s.privateKey)

	err := theirMessenger.SetInstallationMetadata(theirMessenger.installationID, &multidevice.InstallationMetadata{
		Name:       "their-name",
		DeviceType: "their-device-type",
	})
	s.Require().NoError(err)
	response, err := theirMessenger.SendPairInstallation(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(response)
	s.Require().Len(response.Chats, 1)
	s.Require().False(response.Chats[0].Active)

	// Wait for the message to reach its destination
	response, err = WaitOnMessengerResponse(
		s.m,
		func(r *MessengerResponse) bool { return len(r.Installations) > 0 },
		"installation not received",
	)

	s.Require().NoError(err)
	actualInstallation := response.Installations[0]
	s.Require().Equal(theirMessenger.installationID, actualInstallation.ID)
	s.Require().NotNil(actualInstallation.InstallationMetadata)
	s.Require().Equal("their-name", actualInstallation.InstallationMetadata.Name)
	s.Require().Equal("their-device-type", actualInstallation.InstallationMetadata.DeviceType)

	err = s.m.EnableInstallation(theirMessenger.installationID)
	s.Require().NoError(err)

	contactKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	contact, err := buildContact(&contactKey.PublicKey)
	s.Require().NoError(err)
	contact.SystemTags = append(contact.SystemTags, contactAdded)
	err = s.m.SaveContact(contact)
	s.Require().NoError(err)

	// Wait for the message to reach its destination
	response, err = WaitOnMessengerResponse(
		theirMessenger,
		func(r *MessengerResponse) bool { return len(r.Contacts) > 0 && r.Contacts[0].ID == contact.ID },
		"contact not received",
	)
	s.Require().NoError(err)

	actualContact := response.Contacts[0]
	s.Require().Equal(contact.ID, actualContact.ID)
	s.Require().True(actualContact.IsAdded())

	chat := CreatePublicChat(statusChatID, s.m.transport)
	err = s.m.SaveChat(&chat)
	s.Require().NoError(err)

	response, err = WaitOnMessengerResponse(
		theirMessenger,
		func(r *MessengerResponse) bool { return len(r.Chats) > 0 },
		"sync chat not received",
	)

	s.Require().NoError(err)

	actualChat := response.Chats[0]
	s.Require().Equal(statusChatID, actualChat.ID)
	s.Require().True(actualChat.Active)
}

func (s *MessengerInstallationSuite) TestSyncInstallation() {

	// add contact
	contactKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	contact, err := buildContact(&contactKey.PublicKey)
	s.Require().NoError(err)
	contact.SystemTags = append(contact.SystemTags, contactAdded)
	err = s.m.SaveContact(contact)
	s.Require().NoError(err)

	// add chat
	chat := CreatePublicChat(statusChatID, s.m.transport)
	err = s.m.SaveChat(&chat)
	s.Require().NoError(err)

	// pair
	theirMessenger := s.newMessengerWithKey(s.shh, s.privateKey)

	err = theirMessenger.SetInstallationMetadata(theirMessenger.installationID, &multidevice.InstallationMetadata{
		Name:       "their-name",
		DeviceType: "their-device-type",
	})
	s.Require().NoError(err)
	response, err := theirMessenger.SendPairInstallation(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(response)
	s.Require().Len(response.Chats, 1)
	s.Require().False(response.Chats[0].Active)

	// Wait for the message to reach its destination
	response, err = WaitOnMessengerResponse(
		s.m,
		func(r *MessengerResponse) bool { return len(r.Installations) > 0 },
		"installation not received",
	)

	s.Require().NoError(err)
	actualInstallation := response.Installations[0]
	s.Require().Equal(theirMessenger.installationID, actualInstallation.ID)
	s.Require().NotNil(actualInstallation.InstallationMetadata)
	s.Require().Equal("their-name", actualInstallation.InstallationMetadata.Name)
	s.Require().Equal("their-device-type", actualInstallation.InstallationMetadata.DeviceType)

	err = s.m.EnableInstallation(theirMessenger.installationID)
	s.Require().NoError(err)

	// sync
	err = s.m.SyncDevices(context.Background(), "ens-name", "profile-image")
	s.Require().NoError(err)

	var allChats []*Chat
	var actualContact *Contact
	// Wait for the message to reach its destination
	err = tt.RetryWithBackOff(func() error {
		var err error
		response, err = theirMessenger.RetrieveAll()
		if err != nil {
			return err
		}

		allChats = append(allChats, response.Chats...)

		if len(allChats) >= 2 && len(response.Contacts) == 1 {
			actualContact = response.Contacts[0]
			return nil
		}

		return errors.New("Not received all chats & contacts")

	})

	s.Require().NoError(err)

	var statusChat *Chat
	for _, c := range allChats {
		if c.ID == statusChatID {
			statusChat = c
		}
	}

	s.Require().NotNil(statusChat)

	s.Require().True(actualContact.IsAdded())
}

func (s *MessengerInstallationSuite) TestSyncInstallationNewMessages() {

	bob1 := s.m
	// pair
	bob2 := s.newMessengerWithKey(s.shh, s.privateKey)
	alice := s.newMessenger(s.shh)

	err := bob2.SetInstallationMetadata(bob2.installationID, &multidevice.InstallationMetadata{
		Name:       "their-name",
		DeviceType: "their-device-type",
	})
	s.Require().NoError(err)
	response, err := bob2.SendPairInstallation(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(response)
	s.Require().Len(response.Chats, 1)
	s.Require().False(response.Chats[0].Active)

	// Wait for the message to reach its destination
	response, err = WaitOnMessengerResponse(
		bob1,
		func(r *MessengerResponse) bool { return len(r.Installations) > 0 },
		"installation not received",
	)

	s.Require().NoError(err)
	actualInstallation := response.Installations[0]
	s.Require().Equal(bob2.installationID, actualInstallation.ID)
	err = bob1.EnableInstallation(bob2.installationID)
	s.Require().NoError(err)

	// send a message from bob1 to alice, it should be received on both bob1 and bob2

	alicePkString := types.EncodeHex(crypto.FromECDSAPub(&alice.identity.PublicKey))
	chat := CreateOneToOneChat(alicePkString, &alice.identity.PublicKey, bob1.transport)
	s.Require().NoError(bob1.SaveChat(&chat))

	inputMessage := buildTestMessage(chat)
	_, err = s.m.SendChatMessage(context.Background(), inputMessage)
	s.Require().NoError(err)

	// Wait for the message to reach its destination
	_, err = WaitOnMessengerResponse(
		bob2,
		func(r *MessengerResponse) bool { return len(r.Messages) > 0 },
		"message not received",
	)
	s.Require().NoError(err)
}

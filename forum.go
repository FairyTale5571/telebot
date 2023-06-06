package telebot

// ForumTopic https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicCreated https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct{}

// ForumTopicEdited https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicReopened https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct{}

// GeneralForumTopicHidden https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct{}

// UserShared https://core.telegram.org/bots/api#usershared
type UserShared struct {
	RequestID int    `json:"request_id"`
	UserID    ChatID `json:"user_id"`
}

// ChatShared https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestID int    `json:"request_id"`
	ChatID    ChatID `json:"chat_id"`
}

// WriteAccessAllowed https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	WebAppName string `json:"web_app_name,omitempty"`
}

// Params for methods

type CreateForumTopicParams struct {
	ChatID            ChatID `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type EditForumTopicParams struct {
	ChatID            ChatID `json:"chat_id"`
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type CloseForumTopicParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int    `json:"message_thread_id"`
}

type ReopenForumTopicParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int    `json:"message_thread_id"`
}

type DeleteForumTopicParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int    `json:"message_thread_id"`
}

type UnpinAllForumTopicMessagesParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int    `json:"message_thread_id"`
}

type EditGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
	Name   string `json:"name"`
}

type CloseGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

type ReopenGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

type HideGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

type UnhideGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

// CreateForumTopic creates a new forum topic.
func (b *Bot) CreateForumTopic(topic *CreateForumTopicParams) (*ForumTopic, error) {
	var res *ForumTopic

	return nil, nil
}

// EditForumTopic edits a forum topic.
func (b *Bot) EditForumTopic(topic *EditForumTopicParams) error {
	return nil
}

// CloseForumTopic closes a forum topic.
func (b *Bot) CloseForumTopic(topic *CloseForumTopicParams) error {
	return nil
}

// ReopenForumTopic reopens a forum topic.
func (b *Bot) ReopenForumTopic(topic *ReopenForumTopicParams) error {
	return nil
}

// DeleteForumTopic deletes a forum topic.
func (b *Bot) DeleteForumTopic(topic *DeleteForumTopicParams) error {
	return nil
}

// UnpinAllForumTopicMessages unpins all messages in a forum topic.
func (b *Bot) UnpinAllForumTopicMessages(topic *UnpinAllForumTopicMessagesParams) error {
	return nil
}

// EditGeneralForumTopic edits a general forum topic.
func (b *Bot) EditGeneralForumTopic(topic *EditGeneralForumTopicParams) (*ForumTopic, error) {
	return nil, nil
}

// CloseGeneralForumTopic closes a general forum topic.
func (b *Bot) CloseGeneralForumTopic(topic *CloseGeneralForumTopicParams) (*ForumTopic, error) {
	return nil, nil
}

// ReopenGeneralForumTopic reopens a general forum topic.
func (b *Bot) ReopenGeneralForumTopic(topic *ReopenGeneralForumTopicParams) (*ForumTopic, error) {
	return nil, nil
}

// HideGeneralForumTopic hides a general forum topic.
func (b *Bot) HideGeneralForumTopic(topic *HideGeneralForumTopicParams) (*ForumTopic, error) {
	return nil, nil
}

// UnhideGeneralForumTopic unhides a general forum topic.
func (b *Bot) UnhideGeneralForumTopic(topic *UnhideGeneralForumTopicParams) (*ForumTopic, error) {
	return nil, nil
}

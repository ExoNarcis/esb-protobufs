// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/notifications/slack.proto

package slack

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SlackSendParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId   string        `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ThreadTs    string        `protobuf:"bytes,2,opt,name=thread_ts,json=threadTs,proto3" json:"thread_ts,omitempty"`
	Text        string        `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Attachments []*Attachment `protobuf:"bytes,4,rep,name=attachments,proto3" json:"attachments,omitempty"`
	// https://api.slack.com/methods/chat.postMessage#arg_blocks
	// https://api.slack.com/reference/block-kit/blocks
	// https://api.slack.com/messaging/composing/layouts
	Blocks []*any.Any `protobuf:"bytes,5,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *SlackSendParams) Reset() {
	*x = SlackSendParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlackSendParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackSendParams) ProtoMessage() {}

func (x *SlackSendParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackSendParams.ProtoReflect.Descriptor instead.
func (*SlackSendParams) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{0}
}

func (x *SlackSendParams) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SlackSendParams) GetThreadTs() string {
	if x != nil {
		return x.ThreadTs
	}
	return ""
}

func (x *SlackSendParams) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SlackSendParams) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *SlackSendParams) GetBlocks() []*any.Any {
	if x != nil {
		return x.Blocks
	}
	return nil
}

// attachments (https://api.slack.com/methods/chat.postMessage#arg_attachments)
type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color         string              `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Fallback      string              `protobuf:"bytes,2,opt,name=fallback,proto3" json:"fallback,omitempty"`
	CallbackId    string              `protobuf:"bytes,3,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	Id            int32               `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId      string              `protobuf:"bytes,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	AuthorName    string              `protobuf:"bytes,6,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	AuthorSubName string              `protobuf:"bytes,7,opt,name=author_sub_name,json=authorSubName,proto3" json:"author_sub_name,omitempty"`
	AuthorLink    string              `protobuf:"bytes,8,opt,name=author_link,json=authorLink,proto3" json:"author_link,omitempty"`
	AuthorIcon    string              `protobuf:"bytes,9,opt,name=author_icon,json=authorIcon,proto3" json:"author_icon,omitempty"`
	Title         string              `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	TitleLink     string              `protobuf:"bytes,11,opt,name=title_link,json=titleLink,proto3" json:"title_link,omitempty"`
	PreText       string              `protobuf:"bytes,12,opt,name=pre_text,json=preText,proto3" json:"pre_text,omitempty"`
	Text          string              `protobuf:"bytes,13,opt,name=text,proto3" json:"text,omitempty"`
	ImageUrl      string              `protobuf:"bytes,14,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	ThumbUrl      string              `protobuf:"bytes,15,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	Fields        []*AttachmentField  `protobuf:"bytes,16,rep,name=fields,proto3" json:"fields,omitempty"`
	Actions       []*AttachmentAction `protobuf:"bytes,17,rep,name=actions,proto3" json:"actions,omitempty"`
	Footer        string              `protobuf:"bytes,18,opt,name=footer,proto3" json:"footer,omitempty"`
	FooterIcon    string              `protobuf:"bytes,19,opt,name=footer_icon,json=footerIcon,proto3" json:"footer_icon,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{1}
}

func (x *Attachment) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Attachment) GetFallback() string {
	if x != nil {
		return x.Fallback
	}
	return ""
}

func (x *Attachment) GetCallbackId() string {
	if x != nil {
		return x.CallbackId
	}
	return ""
}

func (x *Attachment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Attachment) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Attachment) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *Attachment) GetAuthorSubName() string {
	if x != nil {
		return x.AuthorSubName
	}
	return ""
}

func (x *Attachment) GetAuthorLink() string {
	if x != nil {
		return x.AuthorLink
	}
	return ""
}

func (x *Attachment) GetAuthorIcon() string {
	if x != nil {
		return x.AuthorIcon
	}
	return ""
}

func (x *Attachment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Attachment) GetTitleLink() string {
	if x != nil {
		return x.TitleLink
	}
	return ""
}

func (x *Attachment) GetPreText() string {
	if x != nil {
		return x.PreText
	}
	return ""
}

func (x *Attachment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Attachment) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Attachment) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *Attachment) GetFields() []*AttachmentField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Attachment) GetActions() []*AttachmentAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Attachment) GetFooter() string {
	if x != nil {
		return x.Footer
	}
	return ""
}

func (x *Attachment) GetFooterIcon() string {
	if x != nil {
		return x.FooterIcon
	}
	return ""
}

type AttachmentField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Short bool   `protobuf:"varint,3,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *AttachmentField) Reset() {
	*x = AttachmentField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentField) ProtoMessage() {}

func (x *AttachmentField) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentField.ProtoReflect.Descriptor instead.
func (*AttachmentField) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{2}
}

func (x *AttachmentField) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AttachmentField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AttachmentField) GetShort() bool {
	if x != nil {
		return x.Short
	}
	return false
}

type AttachmentAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text            string                         `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Style           string                         `protobuf:"bytes,3,opt,name=style,proto3" json:"style,omitempty"`
	Type            string                         `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Value           string                         `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	DataSource      string                         `protobuf:"bytes,6,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	Url             string                         `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Options         []*AttachmentActionOption      `protobuf:"bytes,8,rep,name=options,proto3" json:"options,omitempty"`
	SelectedOptions []*AttachmentActionOption      `protobuf:"bytes,9,rep,name=selected_options,json=selectedOptions,proto3" json:"selected_options,omitempty"`
	OptionGroups    []*AttachmentActionOptionGroup `protobuf:"bytes,10,rep,name=option_groups,json=optionGroups,proto3" json:"option_groups,omitempty"`
	Confirm         *AttachmentConfirm             `protobuf:"bytes,11,opt,name=confirm,proto3" json:"confirm,omitempty"`
}

func (x *AttachmentAction) Reset() {
	*x = AttachmentAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentAction) ProtoMessage() {}

func (x *AttachmentAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentAction.ProtoReflect.Descriptor instead.
func (*AttachmentAction) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{3}
}

func (x *AttachmentAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttachmentAction) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AttachmentAction) GetStyle() string {
	if x != nil {
		return x.Style
	}
	return ""
}

func (x *AttachmentAction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttachmentAction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AttachmentAction) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

func (x *AttachmentAction) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AttachmentAction) GetOptions() []*AttachmentActionOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *AttachmentAction) GetSelectedOptions() []*AttachmentActionOption {
	if x != nil {
		return x.SelectedOptions
	}
	return nil
}

func (x *AttachmentAction) GetOptionGroups() []*AttachmentActionOptionGroup {
	if x != nil {
		return x.OptionGroups
	}
	return nil
}

func (x *AttachmentAction) GetConfirm() *AttachmentConfirm {
	if x != nil {
		return x.Confirm
	}
	return nil
}

type AttachmentActionOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text        string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Value       string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AttachmentActionOption) Reset() {
	*x = AttachmentActionOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentActionOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentActionOption) ProtoMessage() {}

func (x *AttachmentActionOption) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentActionOption.ProtoReflect.Descriptor instead.
func (*AttachmentActionOption) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{4}
}

func (x *AttachmentActionOption) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AttachmentActionOption) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AttachmentActionOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AttachmentActionOptionGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string                    `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Options []*AttachmentActionOption `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *AttachmentActionOptionGroup) Reset() {
	*x = AttachmentActionOptionGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentActionOptionGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentActionOptionGroup) ProtoMessage() {}

func (x *AttachmentActionOptionGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentActionOptionGroup.ProtoReflect.Descriptor instead.
func (*AttachmentActionOptionGroup) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{5}
}

func (x *AttachmentActionOptionGroup) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AttachmentActionOptionGroup) GetOptions() []*AttachmentActionOption {
	if x != nil {
		return x.Options
	}
	return nil
}

type AttachmentConfirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text         string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	OkText       string `protobuf:"bytes,3,opt,name=ok_text,json=okText,proto3" json:"ok_text,omitempty"`
	DissmissText string `protobuf:"bytes,4,opt,name=dissmiss_text,json=dissmissText,proto3" json:"dissmiss_text,omitempty"`
}

func (x *AttachmentConfirm) Reset() {
	*x = AttachmentConfirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentConfirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentConfirm) ProtoMessage() {}

func (x *AttachmentConfirm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentConfirm.ProtoReflect.Descriptor instead.
func (*AttachmentConfirm) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{6}
}

func (x *AttachmentConfirm) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AttachmentConfirm) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AttachmentConfirm) GetOkText() string {
	if x != nil {
		return x.OkText
	}
	return ""
}

func (x *AttachmentConfirm) GetDissmissText() string {
	if x != nil {
		return x.DissmissText
	}
	return ""
}

type SlackSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Ts      string `protobuf:"bytes,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *SlackSendResponse) Reset() {
	*x = SlackSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notifications_slack_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlackSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackSendResponse) ProtoMessage() {}

func (x *SlackSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_slack_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackSendResponse.ProtoReflect.Descriptor instead.
func (*SlackSendResponse) Descriptor() ([]byte, []int) {
	return file_proto_notifications_slack_proto_rawDescGZIP(), []int{7}
}

func (x *SlackSendResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *SlackSendResponse) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *SlackSendResponse) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

var File_proto_notifications_slack_proto protoreflect.FileDescriptor

var file_proto_notifications_slack_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0f, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x54, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3b, 0x0a,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0xe1, 0x04, 0x0a, 0x0a, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x63,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c,
	0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x6f, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x0f,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x22, 0xcd, 0x03, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x50, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4f, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x22, 0x64, 0x0a, 0x16, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x1b, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7b, 0x0a, 0x11, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6b,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6b, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x73,
	0x6d, 0x69, 0x73, 0x73, 0x54, 0x65, 0x78, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x53, 0x6c, 0x61, 0x63,
	0x6b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x32, 0x6c, 0x0a, 0x05, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x12, 0x63, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x65,
	0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notifications_slack_proto_rawDescOnce sync.Once
	file_proto_notifications_slack_proto_rawDescData = file_proto_notifications_slack_proto_rawDesc
)

func file_proto_notifications_slack_proto_rawDescGZIP() []byte {
	file_proto_notifications_slack_proto_rawDescOnce.Do(func() {
		file_proto_notifications_slack_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notifications_slack_proto_rawDescData)
	})
	return file_proto_notifications_slack_proto_rawDescData
}

var file_proto_notifications_slack_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_notifications_slack_proto_goTypes = []interface{}{
	(*SlackSendParams)(nil),             // 0: notifications.SlackSendParams
	(*Attachment)(nil),                  // 1: notifications.Attachment
	(*AttachmentField)(nil),             // 2: notifications.AttachmentField
	(*AttachmentAction)(nil),            // 3: notifications.AttachmentAction
	(*AttachmentActionOption)(nil),      // 4: notifications.AttachmentActionOption
	(*AttachmentActionOptionGroup)(nil), // 5: notifications.AttachmentActionOptionGroup
	(*AttachmentConfirm)(nil),           // 6: notifications.AttachmentConfirm
	(*SlackSendResponse)(nil),           // 7: notifications.SlackSendResponse
	(*any.Any)(nil),                     // 8: google.protobuf.Any
}
var file_proto_notifications_slack_proto_depIdxs = []int32{
	1,  // 0: notifications.SlackSendParams.attachments:type_name -> notifications.Attachment
	8,  // 1: notifications.SlackSendParams.blocks:type_name -> google.protobuf.Any
	2,  // 2: notifications.Attachment.fields:type_name -> notifications.AttachmentField
	3,  // 3: notifications.Attachment.actions:type_name -> notifications.AttachmentAction
	4,  // 4: notifications.AttachmentAction.options:type_name -> notifications.AttachmentActionOption
	4,  // 5: notifications.AttachmentAction.selected_options:type_name -> notifications.AttachmentActionOption
	5,  // 6: notifications.AttachmentAction.option_groups:type_name -> notifications.AttachmentActionOptionGroup
	6,  // 7: notifications.AttachmentAction.confirm:type_name -> notifications.AttachmentConfirm
	4,  // 8: notifications.AttachmentActionOptionGroup.options:type_name -> notifications.AttachmentActionOption
	0,  // 9: notifications.Slack.Send:input_type -> notifications.SlackSendParams
	7,  // 10: notifications.Slack.Send:output_type -> notifications.SlackSendResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_notifications_slack_proto_init() }
func file_proto_notifications_slack_proto_init() {
	if File_proto_notifications_slack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_notifications_slack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlackSendParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentActionOption); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentActionOptionGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentConfirm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_notifications_slack_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlackSendResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_notifications_slack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notifications_slack_proto_goTypes,
		DependencyIndexes: file_proto_notifications_slack_proto_depIdxs,
		MessageInfos:      file_proto_notifications_slack_proto_msgTypes,
	}.Build()
	File_proto_notifications_slack_proto = out.File
	file_proto_notifications_slack_proto_rawDesc = nil
	file_proto_notifications_slack_proto_goTypes = nil
	file_proto_notifications_slack_proto_depIdxs = nil
}
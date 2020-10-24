package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents an List request.
type ListRequest struct {
	Limit                    uint   `json:"limit"`
	AttachedToUserOnly       bool   `json:"attachedToUserOnly"`
	AttachedToLocalUserOnly  bool   `json:"attachedToLocalUserOnly"`
	AttachedToRemoteUserOnly bool   `json:"attachedToRemoteUserOnly"`
	Sort                     string `json:"sort"`
}

// ListOptions are all the options you can play with.
type ListOptions struct {
	Limit                    uint
	AttachedToUserOnly       bool
	AttachedToLocalUserOnly  bool
	AttachedToRemoteUserOnly bool
	Sort                     string
}

// Validate options.
func (options *ListOptions) Validate() error {
	if options.Sort == "" {
		return core.MissingOptionsError{
			Endpoint:      "Hashtags/List",
			Struct:        "ListOptions",
			MissingFields: []string{"Sort"},
		}
	}

	return nil
}

// List endpoint.
func (s *Service) List(options *ListOptions) ([]models.Hashtag, error) {
	var response []models.Hashtag

	if options == nil {
		return response, core.MissingOptionsError{
			Endpoint: "Hashtags/List",
			Struct:   "ListOptions",
		}
	}

	err := options.Validate()
	if err != nil {
		return response, err
	}

	request := ListRequest{
		Limit:                    options.Limit,
		AttachedToUserOnly:       options.AttachedToUserOnly,
		AttachedToLocalUserOnly:  options.AttachedToLocalUserOnly,
		AttachedToRemoteUserOnly: options.AttachedToRemoteUserOnly,
		Sort:                     options.Sort,
	}

	err = s.Call(
		&core.BaseRequest{Request: &request, Path: "/hashtags/list"},
		&response,
	)

	return response, err
}

package emby

import "time"

// SystemInfo is returned from a call to GetSystemInformation
type SystemInfo struct {
	SystemUpdateLevel                    string        `json:"SystemUpdateLevel"`
	OperatingSystemDisplayName           string        `json:"OperatingSystemDisplayName"`
	HasPendingRestart                    bool          `json:"HasPendingRestart"`
	IsShuttingDown                       bool          `json:"IsShuttingDown"`
	SupportsLibraryMonitor               bool          `json:"SupportsLibraryMonitor"`
	WebSocketPortNumber                  int           `json:"WebSocketPortNumber"`
	CompletedInstallations               []interface{} `json:"CompletedInstallations"`
	CanSelfRestart                       bool          `json:"CanSelfRestart"`
	CanSelfUpdate                        bool          `json:"CanSelfUpdate"`
	CanLaunchWebBrowser                  bool          `json:"CanLaunchWebBrowser"`
	ProgramDataPath                      string        `json:"ProgramDataPath"`
	ItemsByNamePath                      string        `json:"ItemsByNamePath"`
	CachePath                            string        `json:"CachePath"`
	LogPath                              string        `json:"LogPath"`
	InternalMetadataPath                 string        `json:"InternalMetadataPath"`
	TranscodingTempPath                  string        `json:"TranscodingTempPath"`
	HTTPServerPortNumber                 int           `json:"HttpServerPortNumber"`
	SupportsHTTPS                        bool          `json:"SupportsHttps"`
	HTTPSPortNumber                      int           `json:"HttpsPortNumber"`
	HasUpdateAvailable                   bool          `json:"HasUpdateAvailable"`
	SupportsAutoRunAtStartup             bool          `json:"SupportsAutoRunAtStartup"`
	HardwareAccelerationRequiresPremiere bool          `json:"HardwareAccelerationRequiresPremiere"`
	LocalAddress                         string        `json:"LocalAddress"`
	WanAddress                           string        `json:"WanAddress"`
	ServerName                           string        `json:"ServerName"`
	Version                              string        `json:"Version"`
	OperatingSystem                      string        `json:"OperatingSystem"`
	ID                                   string        `json:"Id"`
}

// UserViewsResponse wraps a list of views visible to a given user
type UserViewsResponse struct {
	Items            []UserView `json:"Items"`
	TotalRecordCount int        `json:"TotalRecordCount"`
}

// UserView represents a single user-visible view
type UserView struct {
	Name           string        `json:"Name"`
	ServerID       string        `json:"ServerId"`
	ID             string        `json:"Id"`
	Etag           string        `json:"Etag"`
	DateCreated    string        `json:"DateCreated"`
	CanDelete      bool          `json:"CanDelete"`
	CanDownload    bool          `json:"CanDownload"`
	SortName       string        `json:"SortName"`
	ExternalUrls   []string      `json:"ExternalUrls"`
	Path           string        `json:"Path"`
	Taglines       []string      `json:"Taglines"`
	Genres         []string      `json:"Genres"`
	PlayAccess     string        `json:"PlayAccess"`
	RemoteTrailers []interface{} `json:"RemoteTrailers"`
	ProviderIds    interface{}   `json:"ProviderIds"`
	IsFolder       bool          `json:"IsFolder"`
	ParentID       string        `json:"ParentId"`
	Type           string        `json:"Type"`
	Studios        []interface{} `json:"Studios"`
	GenreItems     []interface{} `json:"GenreItems"`
	UserData       struct {
		PlaybackPositionTicks int    `json:"PlaybackPositionTicks"`
		PlayCount             int    `json:"PlayCount"`
		IsFavorite            bool   `json:"IsFavorite"`
		Played                bool   `json:"Played"`
		Key                   string `json:"Key"`
	} `json:"UserData"`
	ChildCount              int      `json:"ChildCount"`
	DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
	Tags                    []string `json:"Tags"`
	PrimaryImageAspectRatio float32  `json:"PrimaryImageAspectRatio,omitempty"`
	CollectionType          string   `json:"CollectionType"`
	ImageTags               struct {
		Primary string `json:"Primary"`
	} `json:"ImageTags"`
	BackdropImageTags []interface{} `json:"BackdropImageTags"`
	LockedFields      []interface{} `json:"LockedFields"`
	LockData          bool          `json:"LockData"`
}

// MediaItemList is the list of items belonging to a user view
type MediaItemList struct {
	Items            []MediaItem `json:"Items"`
	TotalRecordCount int         `json:"TotalRecordCount"`
}

// MediaItem represents summary-level information about a given media item in a view
type MediaItem struct {
	Name              string    `json:"Name"`
	ServerID          string    `json:"ServerId"`
	ID                string    `json:"Id"`
	PremiereDate      time.Time `json:"PremiereDate,omitempty"`
	OfficialRating    string    `json:"OfficialRating,omitempty"`
	CommunityRating   float32   `json:"CommunityRating,omitempty"`
	RunTimeTicks      int64     `json:"RunTimeTicks"`
	AspectRatio       string    `json:"AspectRatio,omitempty"`
	ProductionYear    int       `json:"ProductionYear"`
	IsFolder          bool      `json:"IsFolder"`
	Type              string    `json:"Type"`
	UserData          UserData  `json:"UserData"`
	ImageTags         ImageTags `json:"ImageTags"`
	BackdropImageTags []string  `json:"BackdropImageTags"`
	MediaType         string    `json:"MediaType"`
	CriticRating      int       `json:"CriticRating,omitempty"`
	HasSubtitles      bool      `json:"HasSubtitles,omitempty"`
	PartCount         int       `json:"PartCount,omitempty"`
}

// UserData is user-specific data for that media item
type UserData struct {
	PlaybackPositionTicks int    `json:"PlaybackPositionTicks"`
	PlayCount             int    `json:"PlayCount"`
	IsFavorite            bool   `json:"IsFavorite"`
	Played                bool   `json:"Played"`
	Key                   string `json:"Key"`
}

// ImageTags are image tagging details for a media item
type ImageTags struct {
	Primary string `json:"Primary"`
	Logo    string `json:"Logo"`
	Thumb   string `json:"Thumb"`
}

// ExternalURL represents an external URL
type ExternalURL struct {
	Name string `json:"Name"`
	URL  string `json:"Url"`
}

// MediaStream encapulates metadata about a single media stream
type MediaStream struct {
	Codec                  string  `json:"Codec"`
	CodecTag               string  `json:"CodecTag"`
	Language               string  `json:"Language"`
	ColorTransfer          string  `json:"ColorTransfer"`
	ColorSpace             string  `json:"ColorSpace"`
	TimeBase               string  `json:"TimeBase"`
	CodecTimeBase          string  `json:"CodecTimeBase"`
	VideoRange             string  `json:"VideoRange"`
	DisplayTitle           string  `json:"DisplayTitle"`
	NalLengthSize          string  `json:"NalLengthSize"`
	IsInterlaced           bool    `json:"IsInterlaced"`
	IsAVC                  bool    `json:"IsAVC"`
	Bitrate                int64   `json:"Bitrate"`
	BitDepth               int64   `json:"BitDepth"`
	RefFrames              int64   `json:"RefFrames"`
	IsDefault              bool    `json:"IsDefault"`
	IsForced               bool    `json:"IsForced"`
	Height                 int64   `json:"Height"`
	Width                  int64   `json:"Width"`
	AverageFrameRate       float32 `json:"AverageFrameRate"`
	RealFrameRate          float32 `json:"RealFrameRate"`
	Profile                string  `json:"Profile"`
	Type                   string  `json:"Type"`
	AspectRatio            string  `json:"AspectRatio"`
	Index                  int64   `json:"Index"`
	IsExternal             bool    `json:"IsExternal"`
	IsTextSubtitleStream   bool    `json:"IsTextSubtitleStream"`
	SupportsExternalStream bool    `json:"SupportsExternalStream"`
	PixelFormat            string  `json:"PixelFormat"`
	Level                  int64   `json:"Level"`
	IsAnamorphic           bool    `json:"IsAnamorphic"`
}

// MediaSource encapsulates metadata about a source of media
type MediaSource struct {
	Protocol              string            `json:"Protocol"`
	ID                    string            `json:"Id"`
	Path                  string            `json:"Path"`
	Type                  string            `json:"Type"`
	Container             string            `json:"Container"`
	Size                  int64             `json:"Size"`
	Name                  string            `json:"Name"`
	IsRemote              bool              `json:"IsRemote"`
	RunTimeTicks          int64             `json:"RunTimeTicks"`
	ReadAtNativeFramerate bool              `json:"ReadAtNativeFramerate"`
	DiscardCorruptPts     bool              `json:"DiscardCorruptPts"`
	FillWallClockDts      bool              `json:"FillWallClockDts"`
	IgnoreDts             bool              `json:"IgnoreDts"`
	IgnoreIndex           bool              `json:"IgnoreIndex"`
	SupportsTranscoding   bool              `json:"SupportsTranscoding"`
	SupportsDirectStream  bool              `json:"SupportsDirectStream"`
	SupportsDirectPlay    bool              `json:"SupportsDirectPlay"`
	IsInfiniteStream      bool              `json:"IsInfiniteStream"`
	RequiresOpening       bool              `json:"RequiresOpening"`
	RequiresClosing       bool              `json:"RequiresClosing"`
	RequiresLooping       bool              `json:"RequiresLooping"`
	SupportsProbing       bool              `json:"SupportsProbing"`
	Bitrate               int64             `json:"Bitrate"`
	RequiredHTTPHeaders   map[string]string `json:"RequiredHttpHeaders,omitempty"`
	MediaStreams          []MediaStream     `json:"MediaStreams,omitempty"`
}

// ProviderIDs is a wrapper structure around the possible external metadata provider IDs for a media item
type ProviderIDs struct {
	IMDB           string `json:"Imdb"`
	TMDB           string `json:"Tmdb"`
	TMDBCollection string `json:"TmdbCollection"`
}

// Person encapsulates metadata about people who perform specific roles in media
type Person struct {
	Name            string `json:"Name"`
	ID              string `json:"Id"`
	Role            string `json:"Role"`
	Type            string `json:"Type"`
	PrimaryImageTag string `json:"PrimaryImageTag"`
}

// UniqueEntity is a grab-bag struct used for a number of parent elements that need Name/ID pairs
type UniqueEntity struct {
	Name string `json:"Name"`
	ID   int64  `json:"Id"`
}

// Chapter represents a chapter definition in a media element.
type Chapter struct {
	StartPositionTicks int64  `json:"StartPositionTicks"`
	Name               string `json:"Name"`
}

// MediaItemDetail represents the detail on a single media item
type MediaItemDetail struct {
	Name                         string         `json:"Name"`
	OriginalTitle                string         `json:"OriginalTitle"`
	ServerID                     string         `json:"ServerId"`
	ID                           string         `json:"Id"`
	Etag                         string         `json:"Etag"`
	DateCreated                  time.Time      `json:"DateCreated"`
	CanDelete                    bool           `json:"CanDelete"`
	CanDownload                  bool           `json:"CanDownload"`
	PreferredMetadataLanguage    string         `json:"PreferredMetadataLanguage"`
	PreferredMetadataCountryCode string         `json:"PreferredMetadataCountryCode"`
	SupportsSync                 bool           `json:"SupportsSync"`
	Container                    string         `json:"Container"`
	SortName                     string         `json:"SortName"`
	ForcedSortName               string         `json:"ForcedSortName"`
	PremiereDate                 time.Time      `json:"PremiereDate"`
	ExternalURLs                 []ExternalURL  `json:"ExternalUrls"`
	MediaSources                 []MediaSource  `json:"MediaSources"`
	ProductionLocations          []string       `json:"ProductionLocations"`
	Path                         string         `json:"Path"`
	OfficialRating               string         `json:"OfficialRating"`
	CustomRating                 string         `json:"CustomRating"`
	Overview                     string         `json:"Overview"`
	Taglines                     []string       `json:"Taglines"`
	Genres                       []string       `json:"Genres"`
	CommunityRating              float32        `json:"CommunityRating"`
	RunTimeTicks                 float64        `json:"RunTimeTicks"`
	PlayAccess                   string         `json:"PlayAccess"`
	AspectRatio                  string         `json:"AspectRatio"`
	ProductionYear               int64          `json:"ProductionYear"`
	RemoteTrailers               []ExternalURL  `json:"RemoteTrailers"`
	ProviderIDs                  ProviderIDs    `json:"ProviderIDs"`
	IsFolder                     bool           `json:"IsFolder"`
	ParentID                     string         `json:"ParentId"`
	Type                         string         `json:"Movie"`
	People                       []Person       `json:"People"`
	Studios                      []UniqueEntity `json:"Studios"`
	GenreItems                   []UniqueEntity `json:"GenreItems"`
	UserData                     UserData       `json:"UserData"`
	DisplayPreferencesID         string         `json:"DisplayPreferencesId"`
	Tags                         []string       `json:"Tags"`
	PrimaryAspectRatio           float64        `json:"PrimaryAspectRatio"`
	MediaStreams                 []MediaStream  `json:"MediaStreams"`
	ImageTags                    ImageTags      `json:"ImageTags"`
	BackdropImageTags            []string       `json:"BackdropImageTags"`
	Chapters                     []Chapter      `json:"Chapters"`
	MediaType                    string         `json:"MediaType"`
	LockedFields                 []interface{}  `json:"LockedFields"`
	LockData                     bool           `json:"LockData"`
	Width                        int64          `json:"Width"`
	Height                       int64          `json:"Height"`
}

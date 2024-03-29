package transmission

type BasicRequest struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

type BasicResponse struct {
	Result string `json:"result"`
}

// Torrents -----------------------------------------------------------------------

// Get
type TorrentGetRequestArgs struct {
	Ids    *[]int   `json:"ids"`    // torrent ids
	Fields []string `json:"fields"` // array of keys
	Format *string  `json:"format"` // string specifying how to format the torrents response field. Allowed values are objects (default) and table
}

type TorrentGetResponse struct {
	Arguments TorrentGetResponseArguments `json:"arguments"`
	Result    string                      `json:"result"`
}

type TorrentGetResponseArguments struct {
	Torrents []Torrent `json:"torrents"`
}

type Torrent struct {
	ID        int64         `json:"id"`
	Labels    []interface{} `json:"labels"`
	Name      string        `json:"name"`
	Status    int64         `json:"status"`
	TotalSize int64         `json:"totalSize"`
}

// Add
type TorrentAddRequestArgs struct {
	Cookies           *string   `json:"cookies"`           // pointer to a string of one or more cookies.
	DownloadDir       *string   `json:"download-dir"`      // path to download the torrent to
	Filename          *string   `json:"filename"`          // filename or URL of the .torrent file
	Labels            *[]string `json:"labels"`            // array of string labels
	Metainfo          *string   `json:"metainfo"`          // base64-encoded .torrent content
	Paused            *bool     `json:"paused"`            // if true, don't start the torrent
	PeerLimit         *int      `json:"peer-limit"`        // maximum number of peers
	BandwidthPriority *int      `json:"bandwidthPriority"` // torrent's bandwidth tr_priority_t
	FilesWanted       *[]string `json:"files-wanted"`      // indices of file(s) to download
	FilesUnwanted     *[]string `json:"files-unwanted"`    // indices of file(s) to not download
	PriorityHigh      *[]string `json:"priority-high"`     // indices of high-priority file(s)
	PriorityLow       *[]string `json:"priority-low"`      // indices of low-priority file(s)
	PriorityNormal    *[]string `json:"priority-normal"`   // indices of normal-priority file(s)
}

type TorrentAddResponse struct {
	Arguments TorrentAddResponseArguments `json:"arguments"`
	Result    string                      `json:"result"`
}

type TorrentAddResponseArguments struct {
	TorrentAdded TorrentAdded `json:"torrent-added"`
}

type TorrentAdded struct {
	HashString string `json:"hashString"`
	ID         int64  `json:"id"`
	Name       string `json:"name"`
}

// Remove
type TorrentRemoveRequestArgs struct {
	Ids             *[]int `json:"ids"`               // torrent ids
	DeleteLocalData *bool  `json:"delete-local-data"` // delete local data. (default: false)
}

// Session ------------------------------------------------------------------------

type SessionResponse struct {
	Arguments SessionResponseArguments `json:"arguments"`
	Result    string                   `json:"result"`
}

type SessionResponseArguments struct {
	AltSpeedDown                     int64  `json:"alt-speed-down"`
	AltSpeedEnabled                  bool   `json:"alt-speed-enabled"`
	AltSpeedTimeBegin                int64  `json:"alt-speed-time-begin"`
	AltSpeedTimeDay                  int64  `json:"alt-speed-time-day"`
	AltSpeedTimeEnabled              bool   `json:"alt-speed-time-enabled"`
	AltSpeedTimeEnd                  int64  `json:"alt-speed-time-end"`
	AltSpeedUp                       int64  `json:"alt-speed-up"`
	AntiBruteForceEnabled            bool   `json:"anti-brute-force-enabled"`
	AntiBruteForceThreshold          int64  `json:"anti-brute-force-threshold"`
	BlocklistEnabled                 bool   `json:"blocklist-enabled"`
	BlocklistSize                    int64  `json:"blocklist-size"`
	BlocklistURL                     string `json:"blocklist-url"`
	CacheSizeMB                      int64  `json:"cache-size-mb"`
	ConfigDir                        string `json:"config-dir"`
	DefaultTrackers                  string `json:"default-trackers"`
	DhtEnabled                       bool   `json:"dht-enabled"`
	DownloadDir                      string `json:"download-dir"`
	DownloadDirFreeSpace             int64  `json:"download-dir-free-space"`
	DownloadQueueEnabled             bool   `json:"download-queue-enabled"`
	DownloadQueueSize                int64  `json:"download-queue-size"`
	Encryption                       string `json:"encryption"`
	IdleSeedingLimit                 int64  `json:"idle-seeding-limit"`
	IdleSeedingLimitEnabled          bool   `json:"idle-seeding-limit-enabled"`
	IncompleteDir                    string `json:"incomplete-dir"`
	IncompleteDirEnabled             bool   `json:"incomplete-dir-enabled"`
	LpdEnabled                       bool   `json:"lpd-enabled"`
	PeerLimitGlobal                  int64  `json:"peer-limit-global"`
	PeerLimitPerTorrent              int64  `json:"peer-limit-per-torrent"`
	PeerPort                         int64  `json:"peer-port"`
	PeerPortRandomOnStart            bool   `json:"peer-port-random-on-start"`
	PexEnabled                       bool   `json:"pex-enabled"`
	PortForwardingEnabled            bool   `json:"port-forwarding-enabled"`
	QueueStalledEnabled              bool   `json:"queue-stalled-enabled"`
	QueueStalledMinutes              int64  `json:"queue-stalled-minutes"`
	RenamePartialFiles               bool   `json:"rename-partial-files"`
	RPCVersion                       int64  `json:"rpc-version"`
	RPCVersionMinimum                int64  `json:"rpc-version-minimum"`
	RPCVersionSemver                 string `json:"rpc-version-semver"`
	ScriptTorrentAddedEnabled        bool   `json:"script-torrent-added-enabled"`
	ScriptTorrentAddedFilename       string `json:"script-torrent-added-filename"`
	ScriptTorrentDoneEnabled         bool   `json:"script-torrent-done-enabled"`
	ScriptTorrentDoneFilename        string `json:"script-torrent-done-filename"`
	ScriptTorrentDoneSeedingEnabled  bool   `json:"script-torrent-done-seeding-enabled"`
	ScriptTorrentDoneSeedingFilename string `json:"script-torrent-done-seeding-filename"`
	SeedQueueEnabled                 bool   `json:"seed-queue-enabled"`
	SeedQueueSize                    int64  `json:"seed-queue-size"`
	SeedRatioLimit                   int64  `json:"seedRatioLimit"`
	SeedRatioLimited                 bool   `json:"seedRatioLimited"`
	SessionID                        string `json:"session-id"`
	SpeedLimitDown                   int64  `json:"speed-limit-down"`
	SpeedLimitDownEnabled            bool   `json:"speed-limit-down-enabled"`
	SpeedLimitUp                     int64  `json:"speed-limit-up"`
	SpeedLimitUpEnabled              bool   `json:"speed-limit-up-enabled"`
	StartAddedTorrents               bool   `json:"start-added-torrents"`
	TCPEnabled                       bool   `json:"tcp-enabled"`
	TrashOriginalTorrentFiles        bool   `json:"trash-original-torrent-files"`
	Units                            Units  `json:"units"`
	UTPEnabled                       bool   `json:"utp-enabled"`
	Version                          string `json:"version"`
}

type Units struct {
	MemoryBytes int64    `json:"memory-bytes"`
	MemoryUnits []string `json:"memory-units"`
	SizeBytes   int64    `json:"size-bytes"`
	SizeUnits   []string `json:"size-units"`
	SpeedBytes  int64    `json:"speed-bytes"`
	SpeedUnits  []string `json:"speed-units"`
}

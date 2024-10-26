package vndb

type (
	Spoiler   int
	Sexual    int
	Violence  int
	DevStatus int
)

const (
	SpoilerNone Spoiler = iota
	SpoilerLight
	SpoilerFull
)

const (
	SexualSafe Sexual = iota
	SexualSuggestive
	SexualErotic
)

const (
	ViolenceNone Violence = iota
	ViolenceLight
	ViolenceExtreme
)

const (
	Finished DevStatus = iota
	InDevelopment
	Cancelled
)

type Image struct {
	URL           string   `json:"url"`
	ID            string   `json:"id"`
	Violence      Violence `json:"violence"`
	Sexual        Sexual   `json:"sexual"`
	Dims          string   `json:"dims"`
	Thumbnail     string   `json:"thumbnail"`
	ThumbnailDims int      `json:"thumbnail_dims"`
	Votecount     int      `json:"votecount"`
}

type PatchUList struct {
	Vote        int      `json:"vote,omitempty"`
	Notes       string   `json:"notes,omitempty"`
	Started     string   `json:"started,omitempty"`
	Finished    string   `json:"finished"`
	Labels      []int    `json:"labels,omitempty"`
	LabelsSet   []string `json:"labels_set,omitempty"`
	LabelsUnset []string `json:"labels_unset,omitempty"`
}

type Request struct {
	// See https://api.vndb.org/kana#filters
	Filters interface{} `json:"filters,omitempty"`
	// Comma-separated list of fields to fetch for each database item. Dot notation can be used to select nested JSON objects, e.g. "image.url" will select the url field inside the image object. Multiple nested fields can be selected with brackets, e.g. "image{id,url,dims}" is equivalent to "image.id, image.url,image.dims".
	// Every field of interest must be explicitely mentioned, there is no support for wildcard matching. The same applies to nested objects, it is an error to list image without sub-fields in the example above.
	// The top-level id field is always selected by default and does not have to be mentioned in this list.
	Fields string `json:"fields,omitempty"`
	// Field to sort on. Supported values depend on the type of data being queried and are documented separately.
	Sort string `json:"sort,omitempty"`
	// Set to true to sort in descending order.
	Reverse bool `json:"reverse,omitempty"`
	// Number of results per page, max 100. Can also be set to 0 if you’re not interested in the results at all, but just want to verify your query or get the count, compact_filters or normalized_filters.
	Results int `json:"results,omitempty"`
	// Page number to request, starting from 1.
	Page int `json:"page,omitempty"`
	// User ID. This field is mainly used for POST /ulist, but it also sets the default user ID to use for the visual novel “label” filter. Defaults to the currently authenticated user.
	User string `json:"user,omitempty"`
	// Whether the response should include the count field (see below). This option should be avoided when the count is not needed since it has a considerable performance impact.
	Count bool `json:"count,omitempty"`
	// Whether the response should include the compact_filters field (see below).
	CompactFilters bool `json:"compact_filters,omitempty"`
	// Whether the response should include the normalized_filters field (see below).
	NormalizedFilters bool `json:"normalized_filters,omitempty"`
}

type Stats struct {
	Chars     int `json:"chars"`
	Producers int `json:"producers"`
	Releases  int `json:"releases"`
	Staff     int `json:"staff"`
	Tags      int `json:"tags"`
	Traits    int `json:"traits"`
	Vn        int `json:"vn"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	// number of play time votes this user has submitted.
	Lengthvotes int `json:"lengthvotes"`
	// sum of the user’s play time votes, in minutes.
	LengthvotesSum int `json:"lengthvotes_sum"`
}

type AuthInfo struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
}

type UListLabels struct {
	Labels []struct {
		Label   string `json:"label"`
		Private bool   `json:"private"`
		ID      int    `json:"id"`
	} `json:"labels"`
}

type Response struct {
	// Array of objects representing the query results.
	Results interface{} `json:"results"`
	// When true, repeating the query with an incremented page number will yield more results. This is a cheaper form of pagination than using the count field.
	More bool `json:"more"`
	// Only present if the query contained "count":true. Indicates the total number of entries that matched the given filters.
	Count int `json:"count"`
	// Only present if the query contained "compact_filters":true. This is a compact string representation of the filters given in the query.
	CompactFilters string `json:"compact_filters"`
	// Only present if the query contained "normalized_filters":true. This is a normalized JSON representation of the filters given in the query.
	NormalizedFilters interface{} `json:"normalized_filters"`
}

type Tag struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Applicable  bool     `json:"applicable"`
	ID          string   `json:"id"`
	VnCount     int      `json:"vn_count"`
	Category    string   `json:"category"`
	Searchable  string   `json:"searchable"`
	Aliases     []string `json:"aliases"`
}
type Character struct {
	Birthday    [2]int   `json:"birthday,omitempty"`
	Aliases     []string `json:"aliases"`
	Original    string   `json:"original,omitempty"`
	Height      int      `json:"height,omitempty"`
	Age         int      `json:"age,omitempty"`
	Description string   `json:"description"`
	Traits      []Trait  `json:"traits"`
	Waist       int      `json:"waist,omitempty"`
	Name        int      `json:"name,omitempty"`
	Cup         string   `json:"cup,omitempty"`
	Vns         []struct {
		Spoiler Spoiler `json:"spoiler"`
		Role    string  `json:"role"`
		Vn      `json:"vn"`
	} `json:"vns"`
	Image     Image     `json:"image"`
	Weight    int       `json:"weight,omitempty"`
	Hips      int       `json:"hips,omitempty"`
	Sex       [2]string `json:"sex,omitempty"`
	Bust      int       `json:"bust,omitempty"`
	BloodType int       `json:"blood_type,omitempty"`
	ID        string    `json:"id,omitempty"`
}
type Release struct {
	ID     string `json:"id"`
	Engine string `json:"engine,omitempty"`
	Media  []struct {
		Qty    int    `json:"qty"`
		Medium string `json:"medium"`
	} `json:"media"`
	Minage int `json:"minage,omitempty"`
	// yyyy-mm-dd
	Released   string     `json:"released"`
	Gtin       string     `json:"gtin,omitempty"`
	Notes      string     `json:"notes"`
	Extlinks   []Extlinks `json:"extlinks"`
	Languages  Languages  `json:"languages"`
	Official   bool       `json:"official"`
	Voiced     int        `json:"voiced"`
	Title      string     `json:"title"`
	Alttitle   string     `json:"alttitle,omitempty"`
	Platforms  []string   `json:"platforms"`
	Producer   Producer   `json:"producer"`
	Freeware   bool       `json:"freeware"`
	HasEro     bool       `json:"has_ero"`
	Vns        []Vn       `json:"vns"`
	Catalog    string     `json:"catalog,omitempty"`
	Resolution string     `json:"resolution,omitempty"`
	Patch      bool       `json:"patch"`
	Uncensored bool       `json:"uncensored"`
}

// POST /ulist.
type Ulist struct {
	Releases struct {
		ListStatus interface{} `json:"list_status"`
		Release    `json:"release"`
	} `json:"releases"`
	Vote int `json:"vote,omitempty"`
	// unix timpstamp.
	Added int `json:"added"`
	// yyyy-mm-dd
	Finished string `json:"finished,omitempty"`
	Vn       Vn     `json:"vn"`
	// unix timestamp
	Lastmod int `json:"lastmod"`
	Labels  []struct {
		Label string `json:"label"`
		ID    int    `json:"id"`
	} `json:"labels"`
	Started string `json:"started"`
	Notes   string `json:"notes"`
	Voted   int    `json:"voted,omitempty"`
	ID      string `json:"id"`
}
type Trait struct {
	CharCount   int      `json:"char_count"`
	Aliases     []string `json:"aliases"`
	Searchable  bool     `json:"searchable"`
	GroupID     string   `json:"group_id"`
	Applicable  bool     `json:"applicable"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	GroupName   string   `json:"group_name"`
	ID          string   `json:"id"`
}
type Vn struct {
	Aliases []string `json:"aliases"`
	Tags    struct {
		Spoiler Spoiler `json:"spoiler"`
		Rating  int     `json:"rating"`
		Lie     bool    `json:"lie"`
		Tag     `json:"tag"`
	} `json:"tags"`
	Image       Image     `json:"image"`
	Rating      int       `json:"rating,omitempty"`
	Description string    `json:"description"`
	Olang       string    `json:"olang"`
	Devstatus   DevStatus `json:"devstatus"`
	Relations   []struct {
		RelationOfficial bool   `json:"relation_official"`
		Relation         string `json:"relation"`
		Vn               `json:"vn"`
	} `json:"relations"`
	LengthVotes int      `json:"length_votes"`
	Platforms   []string `json:"platforms"`
	Titles      []struct {
		Title    string `json:"title"`
		Official bool   `json:"official"`
		Main     bool   `json:"main"`
		Lang     string `json:"lang"`
		Latin    string `json:"latin,omitempty"`
	} `json:"titles"`
	Alttitle      string     `json:"alttitle,omitempty"`
	Title         string     `json:"title"`
	Languages     []string   `json:"languages"`
	Votecount     int        `json:"votecount"`
	Developers    []Producer `json:"developers"`
	Released      string     `json:"released,omitempty"`
	Length        int        `json:"length,omitempty"`
	LengthMinutes int        `json:"length_minutes,omitempty"`
	ID            string     `json:"id"`
	Screenshots   []struct {
		ID            string   `json:"ID"`
		Sexual        Sexual   `json:"sexual"`
		Dims          [2]int   `json:"dims"`
		URL           string   `json:"url"`
		ThumbnailDims []int    `json:"thumbnail_dims"`
		Thumbnail     string   `json:"thumbnail"`
		Violence      Violence `json:"violence"`
		Votecount     int      `json:"votecount"`
		Release       Release  `json:"release"`
	} `json:"screenshots,omitempty"`
}
type Producer struct {
	Aliases     []string `json:"aliases"`
	Lang        string   `json:"lang"`
	Type        string   `json:"type"`
	Original    string   `json:"original"`
	ID          string   `json:"id"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name"`
}
type Extlinks struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	ID    string `json:"id"`
	URL   string `json:"url"`
}

type Medium []struct {
	Plural string `json:"plural"`
	Label  string `json:"label"`
	ID     string `json:"id"`
}
type Platform []struct {
	Label string `json:"label"`
	ID    string `json:"id"`
}
type Languages struct {
	Main  bool   `json:"main"`
	Latin string `json:"latin,omitempty"`
	Lang  string `json:"lang"`
	Title string `json:"title"`
	Mtl   bool   `json:"mtl"`
}

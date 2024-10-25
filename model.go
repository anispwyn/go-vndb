package vndb

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
	// Comma-separated list of fields to fetch for each database item. Dot notation can be used to select nested JSON objects, e.g. "image.url" will select the url field inside the image object. Multiple nested fields can be selected with brackets, e.g. "image{id,url,dims}" is equivalent to "image.id, image.url, image.dims".
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
	Name        interface{} `json:"name"`
	Description interface{} `json:"description"`
	Applicable  interface{} `json:"applicable"`
	ID          interface{} `json:"id"`
	VnCount     interface{} `json:"vn_count"`
	Category    interface{} `json:"category"`
	Searchable  interface{} `json:"searchable"`
	Aliases     interface{} `json:"aliases"`
}
type Character struct {
	Birthday    interface{} `json:"birthday"`
	Aliases     interface{} `json:"aliases"`
	Original    interface{} `json:"original"`
	Height      interface{} `json:"height"`
	Age         interface{} `json:"age"`
	Description interface{} `json:"description"`
	Traits      struct {
		Spoiler interface{} `json:"spoiler"`
		Lie     interface{} `json:"lie"`
	} `json:"traits"`
	Waist interface{} `json:"waist"`
	Name  interface{} `json:"name"`
	Cup   interface{} `json:"cup"`
	Vns   struct {
		ID          string
		Name        string
		Aliases     []string
		Description string
		Searchable  bool
		Applicable  bool
		GroupID     string
		GroupName   string
		CharCount   int
		Spoiler     interface{} `json:"spoiler"`
		Release     struct {
			ID        string
			Title     string
			Alttitle  string
			Languages []struct {
				Lang string
				Lain string
				Mtl  bool
				Main bool
			}
		} `json:"release,omitempty"`
		Role interface{} `json:"role"`
	} `json:"vns"`
	Image struct {
		URL       interface{} `json:"url"`
		ID        interface{} `json:"id"`
		Violence  interface{} `json:"violence"`
		Sexual    interface{} `json:"sexual"`
		Dims      interface{} `json:"dims"`
		Votecount interface{} `json:"votecount"`
	} `json:"image"`
	Weight    interface{} `json:"weight"`
	Hips      interface{} `json:"hips"`
	Sex       interface{} `json:"sex"`
	Bust      interface{} `json:"bust"`
	BloodType interface{} `json:"blood_type"`
	ID        interface{} `json:"id"`
}
type Release struct {
	ID     interface{} `json:"id"`
	Engine interface{} `json:"engine"`
	Media  struct {
		Qty    interface{} `json:"qty"`
		Medium interface{} `json:"medium"`
	} `json:"media"`
	Minage    interface{} `json:"minage"`
	Released  interface{} `json:"released"`
	Gtin      interface{} `json:"gtin"`
	Notes     interface{} `json:"notes"`
	Languages struct {
		Main  interface{} `json:"main"`
		Latin interface{} `json:"latin"`
		Lang  interface{} `json:"lang"`
		Title interface{} `json:"title"`
		Mtl   interface{} `json:"mtl"`
	} `json:"languages"`
	Extlinks struct {
		Name  interface{} `json:"name"`
		Label interface{} `json:"label"`
		ID    interface{} `json:"id"`
		URL   interface{} `json:"url"`
	} `json:"extlinks"`
	Official  interface{} `json:"official"`
	Voiced    interface{} `json:"voiced"`
	Title     interface{} `json:"title"`
	Alttitle  interface{} `json:"alttitle"`
	Platforms interface{} `json:"platforms"`
	Producers struct {
		Publisher   interface{} `json:"publisher"`
		Developer   interface{} `json:"developer"`
		ID          string
		Name        string
		Aliases     []string
		Lang        string
		Type        string
		Description string
	} `json:"producers"`
	Freeware interface{} `json:"freeware"`
	HasEro   interface{} `json:"has_ero"`
	Vns      struct {
		Aliases interface{} `json:"aliases"`
		Tags    struct {
			ID          string
			Name        string
			Aliases     []string
			Description string
			Catagory    string
			Searchable  bool
			Applicable  bool
			VnCount     int
			Spoiler     interface{} `json:"spoiler"`
			Rating      interface{} `json:"rating"`
			Lie         interface{} `json:"lie"`
		} `json:"tags"`
		Image struct {
			Sexual        interface{} `json:"sexual"`
			Dims          interface{} `json:"dims"`
			URL           interface{} `json:"url"`
			ThumbnailDims interface{} `json:"thumbnail_dims"`
			Thumbnail     interface{} `json:"thumbnail"`
			ID            interface{} `json:"id"`
			Violence      interface{} `json:"violence"`
			Votecount     interface{} `json:"votecount"`
		} `json:"image"`
		Rating      interface{} `json:"rating"`
		Description interface{} `json:"description"`
		Olang       interface{} `json:"olang"`
		Devstatus   interface{} `json:"devstatus"`
		Relations   struct {
			RelationOfficial interface{} `json:"relation_official"`
			Relation         interface{} `json:"relation"`
		} `json:"relations"`
		LengthVotes interface{} `json:"length_votes"`
		Platforms   interface{} `json:"platforms"`
		Popularity  interface{} `json:"popularity"`
		Titles      struct {
			Title    interface{} `json:"title"`
			Official interface{} `json:"official"`
			Main     interface{} `json:"main"`
			Lang     interface{} `json:"lang"`
			Latin    interface{} `json:"latin"`
		} `json:"titles"`
		Alttitle   interface{} `json:"alttitle"`
		Title      interface{} `json:"title"`
		Languages  interface{} `json:"languages"`
		Votecount  interface{} `json:"votecount"`
		Developers struct {
			Aliases     interface{} `json:"aliases"`
			Lang        interface{} `json:"lang"`
			Type        interface{} `json:"type"`
			Original    interface{} `json:"original"`
			ID          interface{} `json:"id"`
			Description interface{} `json:"description"`
			Name        interface{} `json:"name"`
		} `json:"developers"`
		Released      interface{} `json:"released"`
		Length        interface{} `json:"length"`
		LengthMinutes interface{} `json:"length_minutes"`
		ID            interface{} `json:"id"`
		Screenshots   struct {
			Votecount interface{} `json:"votecount"`
			ID        interface{} `json:"id"`
			Violence  interface{} `json:"violence"`
			URL       interface{} `json:"url"`
			Release   struct {
				ID     interface{} `json:"id"`
				Engine interface{} `json:"engine"`
				Media  struct {
					Qty    interface{} `json:"qty"`
					Medium interface{} `json:"medium"`
				} `json:"media"`
				Minage    interface{} `json:"minage"`
				Released  interface{} `json:"released"`
				Gtin      interface{} `json:"gtin"`
				Notes     interface{} `json:"notes"`
				Languages struct {
					Main  interface{} `json:"main"`
					Latin interface{} `json:"latin"`
					Lang  interface{} `json:"lang"`
					Title interface{} `json:"title"`
					Mtl   interface{} `json:"mtl"`
				} `json:"languages"`
				Extlinks struct {
					Name  interface{} `json:"name"`
					Label interface{} `json:"label"`
					ID    interface{} `json:"id"`
					URL   interface{} `json:"url"`
				} `json:"extlinks"`
				Official  interface{} `json:"official"`
				Voiced    interface{} `json:"voiced"`
				Title     interface{} `json:"title"`
				Alttitle  interface{} `json:"alttitle"`
				Platforms interface{} `json:"platforms"`
				Producers struct {
					Publisher   interface{} `json:"publisher"`
					Developer   interface{} `json:"developer"`
					ID          string
					Name        string
					Aliases     []string
					Lang        string
					Type        string
					Description string
				} `json:"producers"`
				Freeware interface{} `json:"freeware"`
				HasEro   interface{} `json:"has_ero"`
				Vns      struct {
					// this is dumb
					Rtype interface{} `json:"rtype"`
				} `json:"vns"`
				Catalog    interface{} `json:"catalog"`
				Resolution interface{} `json:"resolution"`
				Patch      interface{} `json:"patch"`
				Uncensored interface{} `json:"uncensored"`
			} `json:"release"`
			Dims          interface{} `json:"dims"`
			ThumbnailDims interface{} `json:"thumbnail_dims"`
			Thumbnail     interface{} `json:"thumbnail"`
			Sexual        interface{} `json:"sexual"`
		} `json:"screenshots"`
		Rtype interface{} `json:"rtype"`
	} `json:"vns"`
	Catalog    interface{} `json:"catalog"`
	Resolution interface{} `json:"resolution"`
	Patch      interface{} `json:"patch"`
	Uncensored interface{} `json:"uncensored"`
}
type Ulist struct {
	Releases struct {
		ListStatus interface{} `json:"list_status"`
		Inherit    string      `json:"_inherit"`
	} `json:"releases"`
	Vote     interface{} `json:"vote"`
	Added    interface{} `json:"added"`
	Finished interface{} `json:"finished"`
	Vn       struct {
		Inherit string `json:"_inherit"`
	} `json:"vn"`
	Lastmod interface{} `json:"lastmod"`
	Labels  struct {
		Label interface{} `json:"label"`
		ID    interface{} `json:"id"`
	} `json:"labels"`
	Started interface{} `json:"started"`
	Notes   interface{} `json:"notes"`
	Voted   interface{} `json:"voted"`
	ID      interface{} `json:"id"`
}
type Trait struct {
	CharCount   interface{} `json:"char_count"`
	Aliases     interface{} `json:"aliases"`
	Searchable  interface{} `json:"searchable"`
	GroupID     interface{} `json:"group_id"`
	Applicable  interface{} `json:"applicable"`
	Description interface{} `json:"description"`
	Name        interface{} `json:"name"`
	GroupName   interface{} `json:"group_name"`
	ID          interface{} `json:"id"`
}
type Vn struct {
	Aliases interface{} `json:"aliases"`
	Tags    struct {
		ID          string
		Name        string
		Aliases     []string
		Description string
		Catagory    string
		Searchable  bool
		Applicable  bool
		VnCount     int
		Spoiler     interface{} `json:"spoiler"`
		Rating      interface{} `json:"rating"`
		Lie         interface{} `json:"lie"`
	} `json:"tags"`
	Image struct {
		Sexual        interface{} `json:"sexual"`
		Dims          interface{} `json:"dims"`
		URL           interface{} `json:"url"`
		ThumbnailDims interface{} `json:"thumbnail_dims"`
		Thumbnail     interface{} `json:"thumbnail"`
		ID            interface{} `json:"id"`
		Violence      interface{} `json:"violence"`
		Votecount     interface{} `json:"votecount"`
	} `json:"image"`
	Rating      interface{} `json:"rating"`
	Description interface{} `json:"description"`
	Olang       interface{} `json:"olang"`
	Devstatus   interface{} `json:"devstatus"`
	Relations   struct {
		RelationOfficial interface{} `json:"relation_official"`
		Relation         interface{} `json:"relation"`
	} `json:"relations"`
	LengthVotes interface{} `json:"length_votes"`
	Platforms   interface{} `json:"platforms"`
	Popularity  interface{} `json:"popularity"`
	Titles      struct {
		Title    interface{} `json:"title"`
		Official interface{} `json:"official"`
		Main     interface{} `json:"main"`
		Lang     interface{} `json:"lang"`
		Latin    interface{} `json:"latin"`
	} `json:"titles"`
	Alttitle   interface{} `json:"alttitle"`
	Title      interface{} `json:"title"`
	Languages  interface{} `json:"languages"`
	Votecount  interface{} `json:"votecount"`
	Developers struct {
		Aliases     interface{} `json:"aliases"`
		Lang        interface{} `json:"lang"`
		Type        interface{} `json:"type"`
		Original    interface{} `json:"original"`
		ID          interface{} `json:"id"`
		Description interface{} `json:"description"`
		Name        interface{} `json:"name"`
	} `json:"developers"`
	Released      interface{} `json:"released"`
	Length        interface{} `json:"length"`
	LengthMinutes interface{} `json:"length_minutes"`
	ID            interface{} `json:"id"`
	Screenshots   struct {
		Votecount interface{} `json:"votecount"`
		ID        interface{} `json:"id"`
		Violence  interface{} `json:"violence"`
		URL       interface{} `json:"url"`
		Release   struct {
			ID     interface{} `json:"id"`
			Engine interface{} `json:"engine"`
			Media  struct {
				Qty    interface{} `json:"qty"`
				Medium interface{} `json:"medium"`
			} `json:"media"`
			Minage    interface{} `json:"minage"`
			Released  interface{} `json:"released"`
			Gtin      interface{} `json:"gtin"`
			Notes     interface{} `json:"notes"`
			Languages struct {
				Main  interface{} `json:"main"`
				Latin interface{} `json:"latin"`
				Lang  interface{} `json:"lang"`
				Title interface{} `json:"title"`
				Mtl   interface{} `json:"mtl"`
			} `json:"languages"`
			Extlinks struct {
				Name  interface{} `json:"name"`
				Label interface{} `json:"label"`
				ID    interface{} `json:"id"`
				URL   interface{} `json:"url"`
			} `json:"extlinks"`
			Official  interface{} `json:"official"`
			Voiced    interface{} `json:"voiced"`
			Title     interface{} `json:"title"`
			Alttitle  interface{} `json:"alttitle"`
			Platforms interface{} `json:"platforms"`
			Producers struct {
				Publisher   interface{} `json:"publisher"`
				Developer   interface{} `json:"developer"`
				ID          string
				Name        string
				Aliases     []string
				Lang        string
				Type        string
				Description string
			} `json:"producers"`
			Freeware interface{} `json:"freeware"`
			HasEro   interface{} `json:"has_ero"`
			Vns      struct {
				// get the data from the vn u fetched idiot
				Rtype interface{} `json:"rtype"`
			} `json:"vns"`
			Catalog    interface{} `json:"catalog"`
			Resolution interface{} `json:"resolution"`
			Patch      interface{} `json:"patch"`
			Uncensored interface{} `json:"uncensored"`
		} `json:"release"`
		Dims          interface{} `json:"dims"`
		ThumbnailDims interface{} `json:"thumbnail_dims"`
		Thumbnail     interface{} `json:"thumbnail"`
		Sexual        interface{} `json:"sexual"`
	} `json:"screenshots"`
}
type Producer struct {
	Aliases     interface{} `json:"aliases"`
	Lang        interface{} `json:"lang"`
	Type        interface{} `json:"type"`
	Original    interface{} `json:"original"`
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Name        interface{} `json:"name"`
}
type Extlinks struct {
	Release []struct {
		Name      string `json:"name"`
		URLFormat string `json:"url_format"`
		Label     string `json:"label"`
	} `json:"/release"`
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
type Language []struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

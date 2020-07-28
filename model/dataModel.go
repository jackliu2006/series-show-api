package model


type XMLBody struct {
	Series Series 
	Episode  []Episode 
}


type Series struct {
	ID             string `json:"id,omitempty" xml:"id"`
	Actors         string `json:"Actors,omitempty" xml:"Actors"`
	Airs_DayOfWeek string `json:"Airs_DayOfWeek,omitempty" xml:"Airs_DayOfWeek"`
	Airs_Time      string `json:"Airs_Time,omitempty" xml:"Airs_Time"`
	ContentRating  string `json:"ContentRating,omitempty" xml:"ContentRating"`
	FirstAired     string `json:"FirstAired,omitempty" xml:"FirstAired"`
	Genre          string `json:"Genre,omitempty" xml:"Genre"`
	IMDB_ID        string `json:"IMDB_ID,omitempty" xml:"IMDB_ID"`
	Language       string `json:"Language,omitempty" xml:"Language"`
	Network        string `json:"Network,omitempty" xml:"Network"`
	NetworkID      string `json:"NetworkID,omitempty" xml:"NetworkID"`
	Overview       string `json:"Overview,omitempty" xml:"Overview"`
	Rating         string `json:"Rating,omitempty" xml:"Rating"`
	RatingCount    string `json:"RatingCount,omitempty" xml:"RatingCount"`
	Runtime        string `json:"Runtime,omitempty" xml:"FirstName"`
	SeriesID       string `json:"SeriesID,omitempty" xml:"SeriesID"`
	SeriesName     string `json:"SeriesName,omitempty" xml:"SeriesName"`
	Status         string `json:"Status,omitempty" xml:"Status"`
	Added          string `json:"added,omitempty" xml:"added"`
	AddedBy        string `json:"addedBy,omitempty" xml:"addedBy"`
	Banner         string `json:"banner,omitempty" xml:"banner"`
	Fanart         string `json:"fanart,omitempty" xml:"fanart"`
	Finale_aired   string `json:"finale_aired,omitempty" xml:"finale_aired"`
	Lastupdated    string `json:"lastupdated,omitempty" xml:"lastupdated"`
	Poster         string `json:"poster,omitempty" xml:"poster"`
	Tms_wanted_old string `json:"tms_wanted_old,omitempty" xml:"tms_wanted_old"`
	Zap2it_id      string `json:"zap2it_id,omitempty" xml:"zap2it_id"`
}

type Episode struct {
	Id                     string `json:"id,omitempty" xml:"id"`
	Combined_episodenumber string `json:"Combined_episodenumber,omitempty" xml:"Combined_episodenumber"`
	Combined_season        string `json:"Combined_season,omitempty" xml:"Combined_season"`
	DVD_chapter            string `json:"DVD_chapter,omitempty" xml:"DVD_chapter"`
	DVD_discid             string `json:"DVD_discid,omitempty" xml:"DVD_discid"`
	DVD_episodenumber      string `json:"DVD_episodenumber,omitempty" xml:"DVD_episodenumber"`
	DVD_season             string `json:"DVD_season,omitempty" xml:"DVD_season"`
	Director               string `json:"Director,omitempty" xml:"Director"`
	EpImgFlag              string `json:"EpImgFlag,omitempty" xml:"EpImgFlag"`
	EpisodeName            string `json:"EpisodeName,omitempty" xml:"EpisodeName"`
	EpisodeNumber          string `json:"EpisodeNumber,omitempty" xml:"EpisodeNumber"`
	FirstAired             string `json:"FirstAired,omitempty" xml:"FirstAired"`
	GuestStars             string `json:"GuestStars,omitempty" xml:"GuestStars"`
	IMDB_ID                string `json:"IMDB_ID,omitempty" xml:"IMDB_ID"`
	Language               string `json:"Language,omitempty" xml:"Language"`
	Overview               string `json:"Overview,omitempty" xml:"Overview"`
	ProductionCode         string `json:"ProductionCode,omitempty" xml:"ProductionCode"`
	Rating                 string `json:"Rating,omitempty" xml:"Rating"`
	RatingCount            string `json:"RatingCount,omitempty" xml:"RatingCount"`
	SeasonNumber           string `json:"SeasonNumber,omitempty" xml:"SeasonNumber"`
	Writer                 string `json:"Writer,omitempty" xml:"Writer"`
	Absolute_number        string `json:"absolute_number,omitempty" xml:"absolute_number"`
	Airsafter_season       string `json:"airsafter_season,omitempty" xml:"airsafter_season"`
	Airsbefore_episode     string `json:"airsbefore_episode,omitempty" xml:"airsbefore_episode"`
	airsbefore_season      string `json:"airsbefore_season,omitempty" xml:"airsbefore_season"`
	Filename               string `json:"filename,omitempty" xml:"filename"`
	Is_movie               string `json:"is_movie,omitempty" xml:"is_movie"`
	Lastupdated            string `json:"lastupdated,omitempty" xml:"lastupdated"`
	Seasonid               string `json:"iseasonidd,omitempty" xml:"seasonid"`
	Seriesid               string `json:"seriesid,omitempty" xml:"seriesid"`
	Thumb_added            string `json:"thumb_added,omitempty" xml:"thumb_added"`
	Thumb_height           string `json:"thumb_height,omitempty" xml:"thumb_height"`
	Thumb_width            string `json:"thumb_width,omitempty" xml:"thumb_width"`
}
package marvel

type Character struct {
  Id int
  Name string
  Description string
  Modified string
  ResourceURI string
  Urls []string
  Thumbnail string
  Comics []
}

type CharacterDataContainer {
  Offset int
  Limit int
  Total int
  Count int
  Results []Character
}

type CharacterDataWrapper struct {
  Code int
  Status string
  Copyright string
  AttributionText string
  AttributionHTML string
  Data CharacterDataContainer
  Etag string
}

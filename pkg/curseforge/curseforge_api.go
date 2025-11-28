package curseforge

import (
"time"
)

// CurseForgeServer represents a CurseForge API server configuration
type CurseForgeServer struct {
ApiKey string `json:"apiKey"`
Name   string `json:"name"`
Url    string `json:"url"`
}

// ApiError represents an error response from the CurseForge API
type ApiError struct {
Error       string `json:"error"`
Description string `json:"description"`
}

// Game IDs for CurseForge
const (
GameIDMinecraft     = 432
GameIDMinecraftMods = 432
)

// Class IDs for different mod types
const (
ClassIDMods          = 6
ClassIDModpacks      = 4471
ClassIDResourcePacks = 12
ClassIDWorlds        = 17
ClassIDShaders       = 6552
)

// ModLoaderType represents the type of mod loader
type ModLoaderType int

const (
ModLoaderAny        ModLoaderType = 0
ModLoaderForge      ModLoaderType = 1
ModLoaderCauldron   ModLoaderType = 2
ModLoaderLiteLoader ModLoaderType = 3
ModLoaderFabric     ModLoaderType = 4
ModLoaderQuilt      ModLoaderType = 5
ModLoaderNeoForge   ModLoaderType = 6
)

// String returns the string representation of ModLoaderType
func (m ModLoaderType) String() string {
switch m {
case ModLoaderForge:
return "Forge"
case ModLoaderCauldron:
return "Cauldron"
case ModLoaderLiteLoader:
return "LiteLoader"
case ModLoaderFabric:
return "Fabric"
case ModLoaderQuilt:
return "Quilt"
case ModLoaderNeoForge:
return "NeoForge"
default:
return "Any"
}
}

// ModLoaderFromString converts a string to ModLoaderType
func ModLoaderFromString(loader string) ModLoaderType {
switch loader {
case "forge", "Forge":
return ModLoaderForge
case "fabric", "Fabric":
return ModLoaderFabric
case "quilt", "Quilt":
return ModLoaderQuilt
case "neoforge", "NeoForge":
return ModLoaderNeoForge
case "liteloader", "LiteLoader":
return ModLoaderLiteLoader
case "cauldron", "Cauldron":
return ModLoaderCauldron
default:
return ModLoaderAny
}
}

// FileReleaseType represents the release type of a file
type FileReleaseType int

const (
ReleaseTypeRelease FileReleaseType = 1
ReleaseTypeBeta    FileReleaseType = 2
ReleaseTypeAlpha   FileReleaseType = 3
)

// String returns the string representation of FileReleaseType
func (f FileReleaseType) String() string {
switch f {
case ReleaseTypeRelease:
return "release"
case ReleaseTypeBeta:
return "beta"
case ReleaseTypeAlpha:
return "alpha"
default:
return "unknown"
}
}

// FileStatus represents the status of a file
type FileStatus int

const (
FileStatusProcessing       FileStatus = 1
FileStatusChangesRequired  FileStatus = 2
FileStatusUnderReview      FileStatus = 3
FileStatusApproved         FileStatus = 4
FileStatusRejected         FileStatus = 5
FileStatusMalwareDetected  FileStatus = 6
FileStatusDeleted          FileStatus = 7
FileStatusArchived         FileStatus = 8
FileStatusTesting          FileStatus = 9
FileStatusReleased         FileStatus = 10
FileStatusReadyForReview   FileStatus = 11
FileStatusDeprecated       FileStatus = 12
FileStatusBaking           FileStatus = 13
FileStatusAwaitingPublish  FileStatus = 14
FileStatusFailedPublishing FileStatus = 15
)

// DependencyType represents the type of dependency
type DependencyType int

const (
DependencyTypeEmbeddedLibrary DependencyType = 1
DependencyTypeOptional        DependencyType = 2
DependencyTypeRequired        DependencyType = 3
DependencyTypeTool            DependencyType = 4
DependencyTypeIncompatible    DependencyType = 5
DependencyTypeInclude         DependencyType = 6
)

// String returns the string representation of DependencyType
func (d DependencyType) String() string {
switch d {
case DependencyTypeEmbeddedLibrary:
return "embeddedLibrary"
case DependencyTypeOptional:
return "optional"
case DependencyTypeRequired:
return "required"
case DependencyTypeTool:
return "tool"
case DependencyTypeIncompatible:
return "incompatible"
case DependencyTypeInclude:
return "include"
default:
return "unknown"
}
}

// HashAlgo represents the hash algorithm type
type HashAlgo int

const (
HashAlgoSha1 HashAlgo = 1
HashAlgoMd5  HashAlgo = 2
)

// Response wrapper types
type Response[T any] struct {
Data T `json:"data"`
}

type PaginatedResponse[T any] struct {
Data       T          `json:"data"`
Pagination Pagination `json:"pagination"`
}

type Pagination struct {
Index       int `json:"index"`
PageSize    int `json:"pageSize"`
ResultCount int `json:"resultCount"`
TotalCount  int `json:"totalCount"`
}

// Mod represents a CurseForge mod/project
type Mod struct {
ID                            int         `json:"id"`
GameID                        int         `json:"gameId"`
Name                          string      `json:"name"`
Slug                          string      `json:"slug"`
Links                         ModLinks    `json:"links"`
Summary                       string      `json:"summary"`
Status                        int         `json:"status"`
DownloadCount                 int64       `json:"downloadCount"`
IsFeatured                    bool        `json:"isFeatured"`
PrimaryCategoryID             int         `json:"primaryCategoryId"`
Categories                    []Category  `json:"categories"`
ClassID                       int         `json:"classId"`
Authors                       []Author    `json:"authors"`
Logo                          *ModAsset   `json:"logo"`
Screenshots                   []ModAsset  `json:"screenshots"`
MainFileID                    int         `json:"mainFileId"`
LatestFiles                   []File      `json:"latestFiles"`
LatestFilesIndexes            []FileIndex `json:"latestFilesIndexes"`
LatestEarlyAccessFilesIndexes []FileIndex `json:"latestEarlyAccessFilesIndexes"`
DateCreated                   time.Time   `json:"dateCreated"`
DateModified                  time.Time   `json:"dateModified"`
DateReleased                  time.Time   `json:"dateReleased"`
AllowModDistribution          *bool       `json:"allowModDistribution"`
GamePopularityRank            int         `json:"gamePopularityRank"`
IsAvailable                   bool        `json:"isAvailable"`
ThumbsUpCount                 int         `json:"thumbsUpCount"`
}

// ModLinks contains various links for a mod
type ModLinks struct {
WebsiteURL string `json:"websiteUrl"`
WikiURL    string `json:"wikiUrl"`
IssuesURL  string `json:"issuesUrl"`
SourceURL  string `json:"sourceUrl"`
}

// Category represents a mod category
type Category struct {
ID               int       `json:"id"`
GameID           int       `json:"gameId"`
Name             string    `json:"name"`
Slug             string    `json:"slug"`
URL              string    `json:"url"`
IconURL          string    `json:"iconUrl"`
DateModified     time.Time `json:"dateModified"`
IsClass          bool      `json:"isClass"`
ClassID          int       `json:"classId"`
ParentCategoryID int       `json:"parentCategoryId"`
DisplayIndex     int       `json:"displayIndex"`
}

// Author represents a mod author
type Author struct {
ID   int    `json:"id"`
Name string `json:"name"`
URL  string `json:"url"`
}

// ModAsset represents an asset (logo, screenshot) for a mod
type ModAsset struct {
ID           int    `json:"id"`
ModID        int    `json:"modId"`
Title        string `json:"title"`
Description  string `json:"description"`
ThumbnailURL string `json:"thumbnailUrl"`
URL          string `json:"url"`
}

// File represents a mod file/version
type File struct {
ID                   int                   `json:"id"`
GameID               int                   `json:"gameId"`
ModID                int                   `json:"modId"`
IsAvailable          bool                  `json:"isAvailable"`
DisplayName          string                `json:"displayName"`
FileName             string                `json:"fileName"`
ReleaseType          FileReleaseType       `json:"releaseType"`
FileStatus           FileStatus            `json:"fileStatus"`
Hashes               []FileHash            `json:"hashes"`
FileDate             time.Time             `json:"fileDate"`
FileLength           int64                 `json:"fileLength"`
DownloadCount        int64                 `json:"downloadCount"`
FileSizeOnDisk       *int64                `json:"fileSizeOnDisk"`
DownloadURL          string                `json:"downloadUrl"`
GameVersions         []string              `json:"gameVersions"`
SortableGameVersions []SortableGameVersion `json:"sortableGameVersions"`
Dependencies         []FileDependency      `json:"dependencies"`
ExposeAsAlternative  *bool                 `json:"exposeAsAlternative"`
ParentProjectFileID  *int                  `json:"parentProjectFileId"`
AlternateFileID      *int                  `json:"alternateFileId"`
IsServerPack         *bool                 `json:"isServerPack"`
ServerPackFileID     *int                  `json:"serverPackFileId"`
IsEarlyAccessContent *bool                 `json:"isEarlyAccessContent"`
EarlyAccessEndDate   *time.Time            `json:"earlyAccessEndDate"`
FileFingerprint      int64                 `json:"fileFingerprint"`
Modules              []FileModule          `json:"modules"`
}

// FileHash represents a file hash
type FileHash struct {
Value string   `json:"value"`
Algo  HashAlgo `json:"algo"`
}

// SortableGameVersion represents a sortable game version
type SortableGameVersion struct {
GameVersionName        string    `json:"gameVersionName"`
GameVersionPadded      string    `json:"gameVersionPadded"`
GameVersion            string    `json:"gameVersion"`
GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
GameVersionTypeID      int       `json:"gameVersionTypeId"`
}

// FileDependency represents a file dependency
type FileDependency struct {
ModID        int            `json:"modId"`
RelationType DependencyType `json:"relationType"`
}

// FileModule represents a module/file in a mod file
type FileModule struct {
Name        string `json:"name"`
Fingerprint int64  `json:"fingerprint"`
}

// FileIndex represents a file index entry
type FileIndex struct {
GameVersion       string          `json:"gameVersion"`
FileID            int             `json:"fileId"`
Filename          string          `json:"filename"`
ReleaseType       FileReleaseType `json:"releaseType"`
GameVersionTypeID *int            `json:"gameVersionTypeId"`
ModLoader         ModLoaderType   `json:"modLoader"`
}

// FingerprintMatch represents a fingerprint match result
type FingerprintMatch struct {
ID          int    `json:"id"`
File        File   `json:"file"`
LatestFiles []File `json:"latestFiles"`
}

// FingerprintMatchesResult represents the result of fingerprint matching
type FingerprintMatchesResult struct {
IsCacheBuilt             bool                            `json:"isCacheBuilt"`
ExactMatches             []FingerprintMatch              `json:"exactMatches"`
ExactFingerprints        []int64                         `json:"exactFingerprints"`
PartialMatches           []FingerprintMatch              `json:"partialMatches"`
PartialMatchFingerprints FingerprintMatchesResultPartial `json:"partialMatchFingerprints"`
InstalledFingerprints    []int64                         `json:"installedFingerprints"`
UnmatchedFingerprints    []int64                         `json:"unmatchedFingerprints"`
}

// FingerprintMatchesResultPartial represents partial fingerprint matches
type FingerprintMatchesResultPartial struct {
Fingerprints []int64 `json:"fingerprints"`
}

// FingerprintsMatchesRequest represents a request for fingerprint matching
type FingerprintsMatchesRequest struct {
Fingerprints []int64 `json:"fingerprints"`
}

// SearchModsRequest represents search parameters
type SearchModsRequest struct {
GameID            int             `json:"gameId"`
ClassID           int             `json:"classId,omitempty"`
CategoryID        int             `json:"categoryId,omitempty"`
GameVersion       string          `json:"gameVersion,omitempty"`
SearchFilter      string          `json:"searchFilter,omitempty"`
SortField         int             `json:"sortField,omitempty"`
SortOrder         string          `json:"sortOrder,omitempty"`
ModLoaderType     ModLoaderType   `json:"modLoaderType,omitempty"`
ModLoaderTypes    []ModLoaderType `json:"modLoaderTypes,omitempty"`
GameVersionTypeID int             `json:"gameVersionTypeId,omitempty"`
AuthorID          int             `json:"authorId,omitempty"`
Slug              string          `json:"slug,omitempty"`
Index             int             `json:"index,omitempty"`
PageSize          int             `json:"pageSize,omitempty"`
}

// SortField values for search
const (
SortFieldFeatured         = 1
SortFieldPopularity       = 2
SortFieldLastUpdated      = 3
SortFieldName             = 4
SortFieldAuthor           = 5
SortFieldTotalDownloads   = 6
SortFieldCategory         = 7
SortFieldGameVersion      = 8
SortFieldEarlyAccess      = 9
SortFieldFeaturedReleased = 10
SortFieldReleasedDate     = 11
SortFieldRating           = 12
)

// GetModFilesRequest represents request parameters for getting mod files
type GetModFilesRequest struct {
GameVersion       string        `json:"gameVersion,omitempty"`
ModLoaderType     ModLoaderType `json:"modLoaderType,omitempty"`
GameVersionTypeID int           `json:"gameVersionTypeId,omitempty"`
Index             int           `json:"index,omitempty"`
PageSize          int           `json:"pageSize,omitempty"`
}

// GetModsByIDsRequest represents a request to get multiple mods by ID
type GetModsByIDsRequest struct {
ModIDs       []int `json:"modIds"`
FilterPCOnly *bool `json:"filterPcOnly,omitempty"`
}

// GetFilesRequest represents a request to get multiple files by ID
type GetFilesRequest struct {
FileIDs []int `json:"fileIds"`
}

// StringArray is a helper type for encoding string arrays in requests
type StringArray []string

// Game represents a game on CurseForge
type Game struct {
ID           int        `json:"id"`
Name         string     `json:"name"`
Slug         string     `json:"slug"`
DateModified time.Time  `json:"dateModified"`
Assets       GameAssets `json:"assets"`
Status       int        `json:"status"`
APIStatus    int        `json:"apiStatus"`
}

// GameAssets represents game assets
type GameAssets struct {
IconURL  string `json:"iconUrl"`
TileURL  string `json:"tileUrl"`
CoverURL string `json:"coverUrl"`
}

// GameVersionType represents a game version type
type GameVersionType struct {
ID     int    `json:"id"`
GameID int    `json:"gameId"`
Name   string `json:"name"`
Slug   string `json:"slug"`
}

// GameVersion represents a game version
type GameVersion struct {
ID   int    `json:"id"`
Slug string `json:"slug"`
Name string `json:"name"`
}

// MinecraftVersionInfo represents Minecraft version info
type MinecraftVersionInfo struct {
ID                    int       `json:"id"`
GameVersionID         int       `json:"gameVersionId"`
VersionString         string    `json:"versionString"`
JarDownloadURL        string    `json:"jarDownloadUrl"`
JSONDownloadURL       string    `json:"jsonDownloadUrl"`
Approved              bool      `json:"approved"`
DateModified          time.Time `json:"dateModified"`
GameVersionTypeID     int       `json:"gameVersionTypeId"`
GameVersionStatus     int       `json:"gameVersionStatus"`
GameVersionTypeStatus int       `json:"gameVersionTypeStatus"`
}

// MinecraftModLoaderInfo represents Minecraft mod loader info
type MinecraftModLoaderInfo struct {
Name         string        `json:"name"`
GameVersion  string        `json:"gameVersion"`
Latest       bool          `json:"latest"`
Recommended  bool          `json:"recommended"`
DateModified time.Time     `json:"dateModified"`
Type         ModLoaderType `json:"type"`
}
